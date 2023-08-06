package impl

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/endpoint"
)

func (i *impl) RegistryEndpoint(
	ctx context.Context, in *endpoint.RegistryRequest) (
		*endpoint.EndpointSet, error) {
			set := endpoint.NewEndpointSet()
			for m := range in.Items {
				r := in.Items[m]
				ep := endpoint.New(r)
				_, err := i.col.InsertOne(ctx, ep)
				if err != nil {
					return nil, err
				}
				set.Items = append(set.Items, ep)
			}
			return set, nil

}