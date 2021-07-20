//
// DISCLAIMER
//
// Copyright 2020-2021 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Adam Janikowski
// Author Tomasz Mielech
//

package resources

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arangodb/kube-arangodb/pkg/util/errors"

	"github.com/arangodb/kube-arangodb/pkg/deployment/features"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/jwt"
	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	"github.com/arangodb/kube-arangodb/pkg/deployment/pod"
	"github.com/arangodb/kube-arangodb/pkg/util"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil/probes"
	core "k8s.io/api/core/v1"
)

type Probe interface {
	Create() *core.Probe

	SetSpec(spec *api.ServerGroupProbeSpec)
}

type probeCheckBuilder struct {
	liveness, readiness probeBuilder
}

type probeBuilder func(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error)

func nilProbeBuilder(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	return nil, nil
}

func (r *Resources) getReadinessProbe(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	if !r.isReadinessProbeEnabled(spec, group) {
		return nil, nil
	}

	builders := r.probeBuilders()

	builder, ok := builders[group]
	if !ok {
		return nil, nil
	}

	config, err := builder.readiness(spec, group, version)
	if err != nil {
		return nil, err
	}

	groupSpec := spec.GetServerGroupSpec(group)

	if !groupSpec.HasProbesSpec() {
		return config, nil
	}

	probeSpec := groupSpec.GetProbesSpec()

	config.SetSpec(probeSpec.ReadinessProbeSpec)

	return config, nil
}

func (r *Resources) getLivenessProbe(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	if !r.isLivenessProbeEnabled(spec, group) {
		return nil, nil
	}

	builders := r.probeBuilders()

	builder, ok := builders[group]
	if !ok {
		return nil, nil
	}

	config, err := builder.liveness(spec, group, version)
	if err != nil {
		return nil, err
	}

	groupSpec := spec.GetServerGroupSpec(group)

	if !groupSpec.HasProbesSpec() {
		return config, nil
	}

	probeSpec := groupSpec.GetProbesSpec()

	config.SetSpec(probeSpec.LivenessProbeSpec)

	return config, nil
}

func (r *Resources) isReadinessProbeEnabled(spec api.DeploymentSpec, group api.ServerGroup) bool {
	probe := pod.ReadinessSpec(group)

	groupSpec := spec.GetServerGroupSpec(group)

	if groupSpec.HasProbesSpec() {
		if p := groupSpec.GetProbesSpec().GetReadinessProbeDisabled(); p != nil {
			return !*p && probe.CanBeEnabled
		}
	}

	return probe.CanBeEnabled && probe.EnabledByDefault
}

func (r *Resources) isLivenessProbeEnabled(spec api.DeploymentSpec, group api.ServerGroup) bool {
	probe := pod.LivenessSpec(group)

	groupSpec := spec.GetServerGroupSpec(group)

	if groupSpec.HasProbesSpec() {
		if p := groupSpec.GetProbesSpec().LivenessProbeDisabled; p != nil {
			return !*p && probe.CanBeEnabled
		}
	}

	return probe.CanBeEnabled && probe.EnabledByDefault
}

func (r *Resources) probeBuilders() map[api.ServerGroup]probeCheckBuilder {
	return map[api.ServerGroup]probeCheckBuilder{
		api.ServerGroupSingle: {
			liveness:  r.probeBuilderLivenessCoreSelect(),
			readiness: r.probeBuilderReadinessCoreSelect(),
		},
		api.ServerGroupAgents: {
			liveness:  r.probeBuilderLivenessCoreSelect(),
			readiness: r.probeBuilderReadinessSimpleCoreSelect(),
		},
		api.ServerGroupDBServers: {
			liveness:  r.probeBuilderLivenessCoreSelect(),
			readiness: r.probeBuilderReadinessSimpleCoreSelect(),
		},
		api.ServerGroupCoordinators: {
			liveness:  r.probeBuilderLivenessCoreSelect(),
			readiness: r.probeBuilderReadinessCoreSelect(),
		},
		api.ServerGroupSyncMasters: {
			liveness:  r.probeBuilderLivenessSync,
			readiness: nilProbeBuilder,
		},
		api.ServerGroupSyncWorkers: {
			liveness:  r.probeBuilderLivenessSync,
			readiness: nilProbeBuilder,
		},
	}
}

func (r *Resources) probeCommand(spec api.DeploymentSpec, endpoint string) ([]string, error) {
	binaryPath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exePath := filepath.Join(k8sutil.LifecycleVolumeMountDir, filepath.Base(binaryPath))
	args := []string{
		exePath,
		"lifecycle",
		"probe",
		fmt.Sprintf("--endpoint=%s", endpoint),
	}

	if spec.IsSecure() {
		args = append(args, "--ssl")
	}

	if spec.IsAuthenticated() {
		args = append(args, "--auth")
	}

	return args, nil
}

func (r *Resources) probeBuilderLivenessCoreSelect() probeBuilder {
	if features.JWTRotation().Enabled() {
		return r.probeBuilderLivenessCoreOperator
	}

	return r.probeBuilderLivenessCore
}

func (r *Resources) probeBuilderLivenessCoreOperator(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	args, err := r.probeCommand(spec, "/_api/version")
	if err != nil {
		return nil, err
	}

	return &probes.CMDProbeConfig{
		Command: args,
	}, nil
}

func (r *Resources) probeBuilderLivenessCore(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	authorization := ""
	if spec.IsAuthenticated() {
		secretData, err := r.getJWTSecret(spec)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		authorization, err = jwt.CreateArangodJwtAuthorizationHeaderAllowedPaths(secretData, "kube-arangodb", []string{"/_api/version"})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return &probes.HTTPProbeConfig{
		LocalPath:     "/_api/version",
		Secure:        spec.IsSecure(),
		Authorization: authorization,
	}, nil
}

func (r *Resources) probeBuilderReadinessSimpleCoreSelect() probeBuilder {
	if features.JWTRotation().Enabled() {
		return r.probeBuilderReadinessSimpleCoreOperator
	}

	return r.probeBuilderReadinessSimpleCore
}

func (r *Resources) probeBuilderReadinessSimpleCoreOperator(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	p, err := r.probeBuilderReadinessCoreOperator(spec, group, version)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, nil
	}

	p.SetSpec(&api.ServerGroupProbeSpec{
		InitialDelaySeconds: util.NewInt32(15),
		PeriodSeconds:       util.NewInt32(10),
	})

	return p, nil
}

func (r *Resources) probeBuilderReadinessSimpleCore(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	p, err := r.probeBuilderReadinessCore(spec, group, version)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, nil
	}

	p.SetSpec(&api.ServerGroupProbeSpec{
		InitialDelaySeconds: util.NewInt32(15),
		PeriodSeconds:       util.NewInt32(10),
	})

	return p, nil
}

func (r *Resources) probeBuilderReadinessCoreSelect() probeBuilder {
	if features.JWTRotation().Enabled() {
		return r.probeBuilderReadinessCoreOperator
	}

	return r.probeBuilderReadinessCore
}

func (r *Resources) probeBuilderReadinessCoreOperator(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	// /_admin/server/availability is the way to go, it is available since 3.3.9
	args, err := r.probeCommand(spec, "/_admin/server/availability")
	if err != nil {
		return nil, err
	}

	return &probes.CMDProbeConfig{
		Command:             args,
		InitialDelaySeconds: 2,
		PeriodSeconds:       2,
	}, nil
}

func (r *Resources) probeBuilderReadinessCore(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	// /_admin/server/availability is the way to go, it is available since 3.3.9
	localPath := "/_admin/server/availability"

	authorization := ""
	if spec.IsAuthenticated() {
		secretData, err := r.getJWTSecret(spec)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		authorization, err = jwt.CreateArangodJwtAuthorizationHeaderAllowedPaths(secretData, "kube-arangodb", []string{localPath})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	probeCfg := &probes.HTTPProbeConfig{
		LocalPath:           localPath,
		Secure:              spec.IsSecure(),
		Authorization:       authorization,
		InitialDelaySeconds: 2,
		PeriodSeconds:       2,
	}

	return probeCfg, nil
}

func (r *Resources) probeBuilderLivenessSync(spec api.DeploymentSpec, group api.ServerGroup, version driver.Version) (Probe, error) {
	authorization := ""
	port := k8sutil.ArangoSyncMasterPort
	if group == api.ServerGroupSyncWorkers {
		port = k8sutil.ArangoSyncWorkerPort
	}
	if spec.Sync.Monitoring.GetTokenSecretName() != "" {
		// Use monitoring token
		token, err := r.getSyncMonitoringToken(spec)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		authorization = "bearer " + token
	} else if group == api.ServerGroupSyncMasters {
		// Fall back to JWT secret
		secretData, err := r.getSyncJWTSecret(spec)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		authorization, err = jwt.CreateArangodJwtAuthorizationHeaderAllowedPaths(secretData, "kube-arangodb", []string{"/_api/version"})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	} else {
		// Don't have a probe
		return nil, nil
	}
	return &probes.HTTPProbeConfig{
		LocalPath:     "/_api/version",
		Secure:        spec.IsSecure(),
		Authorization: authorization,
		Port:          port,
	}, nil
}
