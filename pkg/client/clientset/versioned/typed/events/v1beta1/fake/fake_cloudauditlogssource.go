/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/google/knative-gcp/pkg/apis/events/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudAuditLogsSources implements CloudAuditLogsSourceInterface
type FakeCloudAuditLogsSources struct {
	Fake *FakeEventsV1beta1
	ns   string
}

var cloudauditlogssourcesResource = schema.GroupVersionResource{Group: "events.cloud.google.com", Version: "v1beta1", Resource: "cloudauditlogssources"}

var cloudauditlogssourcesKind = schema.GroupVersionKind{Group: "events.cloud.google.com", Version: "v1beta1", Kind: "CloudAuditLogsSource"}

// Get takes name of the cloudAuditLogsSource, and returns the corresponding cloudAuditLogsSource object, and an error if there is any.
func (c *FakeCloudAuditLogsSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CloudAuditLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudauditlogssourcesResource, c.ns, name), &v1beta1.CloudAuditLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudAuditLogsSource), err
}

// List takes label and field selectors, and returns the list of CloudAuditLogsSources that match those selectors.
func (c *FakeCloudAuditLogsSources) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CloudAuditLogsSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudauditlogssourcesResource, cloudauditlogssourcesKind, c.ns, opts), &v1beta1.CloudAuditLogsSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CloudAuditLogsSourceList{ListMeta: obj.(*v1beta1.CloudAuditLogsSourceList).ListMeta}
	for _, item := range obj.(*v1beta1.CloudAuditLogsSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudAuditLogsSources.
func (c *FakeCloudAuditLogsSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudauditlogssourcesResource, c.ns, opts))

}

// Create takes the representation of a cloudAuditLogsSource and creates it.  Returns the server's representation of the cloudAuditLogsSource, and an error, if there is any.
func (c *FakeCloudAuditLogsSources) Create(ctx context.Context, cloudAuditLogsSource *v1beta1.CloudAuditLogsSource, opts v1.CreateOptions) (result *v1beta1.CloudAuditLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudauditlogssourcesResource, c.ns, cloudAuditLogsSource), &v1beta1.CloudAuditLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudAuditLogsSource), err
}

// Update takes the representation of a cloudAuditLogsSource and updates it. Returns the server's representation of the cloudAuditLogsSource, and an error, if there is any.
func (c *FakeCloudAuditLogsSources) Update(ctx context.Context, cloudAuditLogsSource *v1beta1.CloudAuditLogsSource, opts v1.UpdateOptions) (result *v1beta1.CloudAuditLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudauditlogssourcesResource, c.ns, cloudAuditLogsSource), &v1beta1.CloudAuditLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudAuditLogsSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudAuditLogsSources) UpdateStatus(ctx context.Context, cloudAuditLogsSource *v1beta1.CloudAuditLogsSource, opts v1.UpdateOptions) (*v1beta1.CloudAuditLogsSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudauditlogssourcesResource, "status", c.ns, cloudAuditLogsSource), &v1beta1.CloudAuditLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudAuditLogsSource), err
}

// Delete takes name of the cloudAuditLogsSource and deletes it. Returns an error if one occurs.
func (c *FakeCloudAuditLogsSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudauditlogssourcesResource, c.ns, name), &v1beta1.CloudAuditLogsSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudAuditLogsSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudauditlogssourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CloudAuditLogsSourceList{})
	return err
}

// Patch applies the patch and returns the patched cloudAuditLogsSource.
func (c *FakeCloudAuditLogsSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CloudAuditLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudauditlogssourcesResource, c.ns, name, pt, data, subresources...), &v1beta1.CloudAuditLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudAuditLogsSource), err
}
