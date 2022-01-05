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
// Author Ewout Prangsma
// Author Tomasz Mielech
//

package deployment

import (
	"context"

	"github.com/arangodb/kube-arangodb/pkg/util/globals"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/arangodb/kube-arangodb/pkg/deployment/resources/inspector"
	"github.com/arangodb/kube-arangodb/pkg/util"
	"github.com/arangodb/kube-arangodb/pkg/util/k8sutil"
	inspectorInterface "github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector"
)

// removePodFinalizers removes all finalizers from all pods owned by us.
func (d *Deployment) removePodFinalizers(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	log := d.deps.Log

	if err := cachedStatus.IteratePods(func(pod *core.Pod) error {
		if err := k8sutil.RemovePodFinalizers(ctx, cachedStatus, log, d.PodsModInterface(), pod, pod.GetFinalizers(), true); err != nil {
			log.Warn().Err(err).Msg("Failed to remove pod finalizers")
			return err
		}

		ctxChild, cancel := globals.GetGlobalTimeouts().Kubernetes().WithTimeout(ctx)
		defer cancel()

		if err := d.PodsModInterface().Delete(ctxChild, pod.GetName(), meta.DeleteOptions{
			GracePeriodSeconds: util.NewInt64(1),
		}); err != nil {
			if !k8sutil.IsNotFound(err) {
				log.Warn().Err(err).Msg("Failed to remove pod")
				return err
			}
		}
		return nil
	}, inspector.FilterPodsByLabels(k8sutil.LabelsForDeployment(d.GetName(), ""))); err != nil {
		return err
	}

	return nil
}

// removePVCFinalizers removes all finalizers from all PVCs owned by us.
func (d *Deployment) removePVCFinalizers(ctx context.Context, cachedStatus inspectorInterface.Inspector) error {
	log := d.deps.Log

	if err := cachedStatus.IteratePersistentVolumeClaims(func(pvc *core.PersistentVolumeClaim) error {
		if err := k8sutil.RemovePVCFinalizers(ctx, cachedStatus, log, d.PersistentVolumeClaimsModInterface(), pvc, pvc.GetFinalizers(), true); err != nil {
			log.Warn().Err(err).Msg("Failed to remove PVC finalizers")
			return err
		}
		return nil
	}, inspector.FilterPersistentVolumeClaimsByLabels(k8sutil.LabelsForDeployment(d.GetName(), ""))); err != nil {
		return err
	}

	return nil
}
