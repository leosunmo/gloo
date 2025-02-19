// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewEndpoint(namespace, name string) *Endpoint {
	endpoint := &Endpoint{}
	endpoint.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return endpoint
}

func (r *Endpoint) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Endpoint) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Upstreams,
		r.Address,
		r.Port,
	)
}

func (r *Endpoint) GroupVersionKind() schema.GroupVersionKind {
	return EndpointGVK
}

type EndpointList []*Endpoint

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list EndpointList) Find(namespace, name string) (*Endpoint, error) {
	for _, endpoint := range list {
		if endpoint.GetMetadata().Name == name {
			if namespace == "" || endpoint.GetMetadata().Namespace == namespace {
				return endpoint, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find endpoint %v.%v", namespace, name)
}

func (list EndpointList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, endpoint := range list {
		ress = append(ress, endpoint)
	}
	return ress
}

func (list EndpointList) Names() []string {
	var names []string
	for _, endpoint := range list {
		names = append(names, endpoint.GetMetadata().Name)
	}
	return names
}

func (list EndpointList) NamespacesDotNames() []string {
	var names []string
	for _, endpoint := range list {
		names = append(names, endpoint.GetMetadata().Namespace+"."+endpoint.GetMetadata().Name)
	}
	return names
}

func (list EndpointList) Sort() EndpointList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list EndpointList) Clone() EndpointList {
	var endpointList EndpointList
	for _, endpoint := range list {
		endpointList = append(endpointList, resources.Clone(endpoint).(*Endpoint))
	}
	return endpointList
}

func (list EndpointList) Each(f func(element *Endpoint)) {
	for _, endpoint := range list {
		f(endpoint)
	}
}

func (list EndpointList) EachResource(f func(element resources.Resource)) {
	for _, endpoint := range list {
		f(endpoint)
	}
}

func (list EndpointList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Endpoint) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for Endpoint

func (o *Endpoint) GetObjectKind() schema.ObjectKind {
	t := EndpointCrd.TypeMeta()
	return &t
}

func (o *Endpoint) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Endpoint)
}

var (
	EndpointCrd = crd.NewCrd(
		"endpoints",
		EndpointGVK.Group,
		EndpointGVK.Version,
		EndpointGVK.Kind,
		"ep",
		false,
		&Endpoint{})
)

func init() {
	if err := crd.AddCrd(EndpointCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}

var (
	EndpointGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gloo.solo.io",
		Kind:    "Endpoint",
	}
)
