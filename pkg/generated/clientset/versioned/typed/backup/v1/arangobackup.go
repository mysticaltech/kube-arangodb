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

package v1

import (
	"context"
	"time"

	v1 "github.com/arangodb/kube-arangodb/pkg/apis/backup/v1"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoBackupsGetter has a method to return a ArangoBackupInterface.
// A group's client should implement this interface.
type ArangoBackupsGetter interface {
	ArangoBackups(namespace string) ArangoBackupInterface
}

// ArangoBackupInterface has methods to work with ArangoBackup resources.
type ArangoBackupInterface interface {
	Create(ctx context.Context, arangoBackup *v1.ArangoBackup, opts metav1.CreateOptions) (*v1.ArangoBackup, error)
	Update(ctx context.Context, arangoBackup *v1.ArangoBackup, opts metav1.UpdateOptions) (*v1.ArangoBackup, error)
	UpdateStatus(ctx context.Context, arangoBackup *v1.ArangoBackup, opts metav1.UpdateOptions) (*v1.ArangoBackup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ArangoBackup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ArangoBackupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ArangoBackup, err error)
	ArangoBackupExpansion
}

// arangoBackups implements ArangoBackupInterface
type arangoBackups struct {
	client rest.Interface
	ns     string
}

// newArangoBackups returns a ArangoBackups
func newArangoBackups(c *BackupV1Client, namespace string) *arangoBackups {
	return &arangoBackups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the arangoBackup, and returns the corresponding arangoBackup object, and an error if there is any.
func (c *arangoBackups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ArangoBackup, err error) {
	result = &v1.ArangoBackup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangobackups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoBackups that match those selectors.
func (c *arangoBackups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ArangoBackupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ArangoBackupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("arangobackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoBackups.
func (c *arangoBackups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("arangobackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a arangoBackup and creates it.  Returns the server's representation of the arangoBackup, and an error, if there is any.
func (c *arangoBackups) Create(ctx context.Context, arangoBackup *v1.ArangoBackup, opts metav1.CreateOptions) (result *v1.ArangoBackup, err error) {
	result = &v1.ArangoBackup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("arangobackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoBackup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a arangoBackup and updates it. Returns the server's representation of the arangoBackup, and an error, if there is any.
func (c *arangoBackups) Update(ctx context.Context, arangoBackup *v1.ArangoBackup, opts metav1.UpdateOptions) (result *v1.ArangoBackup, err error) {
	result = &v1.ArangoBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangobackups").
		Name(arangoBackup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoBackup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *arangoBackups) UpdateStatus(ctx context.Context, arangoBackup *v1.ArangoBackup, opts metav1.UpdateOptions) (result *v1.ArangoBackup, err error) {
	result = &v1.ArangoBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("arangobackups").
		Name(arangoBackup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(arangoBackup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the arangoBackup and deletes it. Returns an error if one occurs.
func (c *arangoBackups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangobackups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoBackups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("arangobackups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched arangoBackup.
func (c *arangoBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ArangoBackup, err error) {
	result = &v1.ArangoBackup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("arangobackups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
