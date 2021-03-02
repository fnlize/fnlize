/*
Copyright The Fission Authors.

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
	corev1 "github.com/fnlize/fnlize/pkg/apis/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCanaryConfigs implements CanaryConfigInterface
type FakeCanaryConfigs struct {
	Fake *FakeCoreV1
	ns   string
}

var canaryconfigsResource = schema.GroupVersionResource{Group: "fission.io", Version: "v1", Resource: "canaryconfigs"}

var canaryconfigsKind = schema.GroupVersionKind{Group: "fission.io", Version: "v1", Kind: "CanaryConfig"}

// Get takes name of the _canaryConfig, and returns the corresponding canaryConfig object, and an error if there is any.
func (c *FakeCanaryConfigs) Get(name string, options v1.GetOptions) (result *corev1.CanaryConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(canaryconfigsResource, c.ns, name), &corev1.CanaryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.CanaryConfig), err
}

// List takes label and field selectors, and returns the list of CanaryConfigs that match those selectors.
func (c *FakeCanaryConfigs) List(opts v1.ListOptions) (result *corev1.CanaryConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(canaryconfigsResource, canaryconfigsKind, c.ns, opts), &corev1.CanaryConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.CanaryConfigList{ListMeta: obj.(*corev1.CanaryConfigList).ListMeta}
	for _, item := range obj.(*corev1.CanaryConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested canaryConfigs.
func (c *FakeCanaryConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(canaryconfigsResource, c.ns, opts))

}

// Create takes the representation of a _canaryConfig and creates it.  Returns the server's representation of the canaryConfig, and an error, if there is any.
func (c *FakeCanaryConfigs) Create(_canaryConfig *corev1.CanaryConfig) (result *corev1.CanaryConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(canaryconfigsResource, c.ns, _canaryConfig), &corev1.CanaryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.CanaryConfig), err
}

// Update takes the representation of a _canaryConfig and updates it. Returns the server's representation of the canaryConfig, and an error, if there is any.
func (c *FakeCanaryConfigs) Update(_canaryConfig *corev1.CanaryConfig) (result *corev1.CanaryConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(canaryconfigsResource, c.ns, _canaryConfig), &corev1.CanaryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.CanaryConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCanaryConfigs) UpdateStatus(_canaryConfig *corev1.CanaryConfig) (*corev1.CanaryConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(canaryconfigsResource, "status", c.ns, _canaryConfig), &corev1.CanaryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.CanaryConfig), err
}

// Delete takes name of the _canaryConfig and deletes it. Returns an error if one occurs.
func (c *FakeCanaryConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(canaryconfigsResource, c.ns, name), &corev1.CanaryConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCanaryConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(canaryconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &corev1.CanaryConfigList{})
	return err
}

// Patch applies the patch and returns the patched canaryConfig.
func (c *FakeCanaryConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *corev1.CanaryConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(canaryconfigsResource, c.ns, name, pt, data, subresources...), &corev1.CanaryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.CanaryConfig), err
}
