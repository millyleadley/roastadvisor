// Code generated by goa v3.19.1, DO NOT EDIT.
//
// Reviews HTTP client types
//
// Command:
// $ goa gen github.com/millyleadley/roastadvisor/api/design -o api

package client

import (
	reviews "github.com/millyleadley/roastadvisor/api/gen/reviews"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "Reviews" service "List" endpoint HTTP
// response body.
type ListResponseBody []*ReviewResponse

// ReviewResponse is used to define fields on response body types.
type ReviewResponse struct {
	// Unique internal ID of the incident debrief
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ID of the user who left the review
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// The comment associated with the review
	Comment *string `form:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
}

// NewListReviewOK builds a "Reviews" service "List" endpoint result from a
// HTTP "OK" response.
func NewListReviewOK(body []*ReviewResponse) []*reviews.Review {
	v := make([]*reviews.Review, len(body))
	for i, val := range body {
		v[i] = unmarshalReviewResponseToReviewsReview(val)
	}

	return v
}

// ValidateReviewResponse runs the validations defined on ReviewResponse
func ValidateReviewResponse(body *ReviewResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "body"))
	}
	if body.Comment == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("comment", "body"))
	}
	return
}
