// Code generated by goa v3.19.1, DO NOT EDIT.
//
// Reviews endpoints
//
// Command:
// $ goa gen github.com/millyleadley/roastadvisor/api/design -o api

package reviews

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "Reviews" service endpoints.
type Endpoints struct {
	List goa.Endpoint
}

// NewEndpoints wraps the methods of the "Reviews" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List: NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "Reviews" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
}

// NewListEndpoint returns an endpoint function that calls the method "List" of
// service "Reviews".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.List(ctx)
	}
}
