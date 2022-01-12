//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	appsv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/apps/v1"
	backupv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/backup/v1"
	databasev1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/deployment/v1"
	databasev2alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/deployment/v2alpha1"
	replicationv1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/replication/v1"
	replicationv2alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/replication/v2alpha1"
	storagev1alpha "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/typed/storage/v1alpha"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AppsV1() appsv1.AppsV1Interface
	BackupV1() backupv1.BackupV1Interface
	DatabaseV1() databasev1.DatabaseV1Interface
	DatabaseV2alpha1() databasev2alpha1.DatabaseV2alpha1Interface
	ReplicationV1() replicationv1.ReplicationV1Interface
	ReplicationV2alpha1() replicationv2alpha1.ReplicationV2alpha1Interface
	StorageV1alpha() storagev1alpha.StorageV1alphaInterface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	appsV1              *appsv1.AppsV1Client
	backupV1            *backupv1.BackupV1Client
	databaseV1          *databasev1.DatabaseV1Client
	databaseV2alpha1    *databasev2alpha1.DatabaseV2alpha1Client
	replicationV1       *replicationv1.ReplicationV1Client
	replicationV2alpha1 *replicationv2alpha1.ReplicationV2alpha1Client
	storageV1alpha      *storagev1alpha.StorageV1alphaClient
}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return c.appsV1
}

// BackupV1 retrieves the BackupV1Client
func (c *Clientset) BackupV1() backupv1.BackupV1Interface {
	return c.backupV1
}

// DatabaseV1 retrieves the DatabaseV1Client
func (c *Clientset) DatabaseV1() databasev1.DatabaseV1Interface {
	return c.databaseV1
}

// DatabaseV2alpha1 retrieves the DatabaseV2alpha1Client
func (c *Clientset) DatabaseV2alpha1() databasev2alpha1.DatabaseV2alpha1Interface {
	return c.databaseV2alpha1
}

// ReplicationV1 retrieves the ReplicationV1Client
func (c *Clientset) ReplicationV1() replicationv1.ReplicationV1Interface {
	return c.replicationV1
}

// ReplicationV2alpha1 retrieves the ReplicationV2alpha1Client
func (c *Clientset) ReplicationV2alpha1() replicationv2alpha1.ReplicationV2alpha1Interface {
	return c.replicationV2alpha1
}

// StorageV1alpha retrieves the StorageV1alphaClient
func (c *Clientset) StorageV1alpha() storagev1alpha.StorageV1alphaInterface {
	return c.storageV1alpha
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.appsV1, err = appsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.backupV1, err = backupv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.databaseV1, err = databasev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.databaseV2alpha1, err = databasev2alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.replicationV1, err = replicationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.replicationV2alpha1, err = replicationv2alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1alpha, err = storagev1alpha.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.appsV1 = appsv1.NewForConfigOrDie(c)
	cs.backupV1 = backupv1.NewForConfigOrDie(c)
	cs.databaseV1 = databasev1.NewForConfigOrDie(c)
	cs.databaseV2alpha1 = databasev2alpha1.NewForConfigOrDie(c)
	cs.replicationV1 = replicationv1.NewForConfigOrDie(c)
	cs.replicationV2alpha1 = replicationv2alpha1.NewForConfigOrDie(c)
	cs.storageV1alpha = storagev1alpha.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.appsV1 = appsv1.New(c)
	cs.backupV1 = backupv1.New(c)
	cs.databaseV1 = databasev1.New(c)
	cs.databaseV2alpha1 = databasev2alpha1.New(c)
	cs.replicationV1 = replicationv1.New(c)
	cs.replicationV2alpha1 = replicationv2alpha1.New(c)
	cs.storageV1alpha = storagev1alpha.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
