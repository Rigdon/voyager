/*
Copyright 2018 The Voyager Authors.

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

package fake

import (
	v1beta1 "github.com/appscode/voyager/apis/voyager/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificates implements CertificateInterface
type FakeCertificates struct {
	Fake *FakeVoyagerV1beta1
	ns   string
}

var certificatesResource = schema.GroupVersionResource{Group: "voyager.appscode.com", Version: "v1beta1", Resource: "certificates"}

var certificatesKind = schema.GroupVersionKind{Group: "voyager.appscode.com", Version: "v1beta1", Kind: "Certificate"}

// Get takes name of the certificate, and returns the corresponding certificate object, and an error if there is any.
func (c *FakeCertificates) Get(name string, options v1.GetOptions) (result *v1beta1.Certificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificatesResource, c.ns, name), &v1beta1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Certificate), err
}

// List takes label and field selectors, and returns the list of Certificates that match those selectors.
func (c *FakeCertificates) List(opts v1.ListOptions) (result *v1beta1.CertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificatesResource, certificatesKind, c.ns, opts), &v1beta1.CertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CertificateList{}
	for _, item := range obj.(*v1beta1.CertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificates.
func (c *FakeCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificatesResource, c.ns, opts))

}

// Create takes the representation of a certificate and creates it.  Returns the server's representation of the certificate, and an error, if there is any.
func (c *FakeCertificates) Create(certificate *v1beta1.Certificate) (result *v1beta1.Certificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificatesResource, c.ns, certificate), &v1beta1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Certificate), err
}

// Update takes the representation of a certificate and updates it. Returns the server's representation of the certificate, and an error, if there is any.
func (c *FakeCertificates) Update(certificate *v1beta1.Certificate) (result *v1beta1.Certificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificatesResource, c.ns, certificate), &v1beta1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Certificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificates) UpdateStatus(certificate *v1beta1.Certificate) (*v1beta1.Certificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificatesResource, "status", c.ns, certificate), &v1beta1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Certificate), err
}

// Delete takes name of the certificate and deletes it. Returns an error if one occurs.
func (c *FakeCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(certificatesResource, c.ns, name), &v1beta1.Certificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.CertificateList{})
	return err
}

// Patch applies the patch and returns the patched certificate.
func (c *FakeCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Certificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificatesResource, c.ns, name, data, subresources...), &v1beta1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Certificate), err
}
