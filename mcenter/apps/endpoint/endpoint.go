package endpoint

import (
	"fmt"
	"hash/fnv"

	"github.com/hezihua/devplat/mcenter/common/meta"
)

func NewEndpointSet () *EndpointSet{
	return &EndpointSet{
		Items: []*Endpoint{},
	}
}

func NewRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Items: []*CreateEndpointRequest{},
	}
}

func (r *CreateEndpointRequest) UnitKey() string {
	return fmt.Sprintf("%s.%s.%s", r.ServiceId, r.Method, r.Path)
}

func (e *RegistryRequest) Add(item *CreateEndpointRequest) {
	e.Items = append(e.Items, item)
}


func New(spec *CreateEndpointRequest) *Endpoint {
	ep := &Endpoint{
		Meta: meta.NewMeta(),
		Spec: spec,
	}

	h := fnv.New32a()
	_, err := h.Write([]byte(ep.Spec.UnitKey()))
	if err != nil {
		panic(err)
	}
	ep.Meta.Id = fmt.Sprintf("%d", h.Sum32())
	return ep
}