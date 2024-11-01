// Code generated by goa v3.19.1, DO NOT EDIT.
//
// Reviews service
//
// Command:
// $ goa gen github.com/millyleadley/roastadvisor/api/design -o api

package reviews

import (
	"context"
)

// The reviews service provides operations to manage reviews.
type Service interface {
	// List all reviews.
	List(context.Context) (res []*Review, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "roastadvisor"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Reviews"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"List"}

type Review struct {
	// Unique internal ID of the incident debrief
	ID string
	// ID of the user who left the review
	UserID string
	// The comment associated with the review
	Comment string
}
