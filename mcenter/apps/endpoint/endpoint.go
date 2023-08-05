package endpoint

func NewRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Items: []*CreateEndpointRequest{},
	}
}

func (e *RegistryRequest) Add(item *CreateEndpointRequest) {
	e.Items = append(e.Items, item)
}
