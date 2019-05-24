// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetGroupRequest
type GetGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the group.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// GroupName is a required field
	GroupName *string `min:"1" type:"string" required:"true"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s GetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGroupInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetGroup request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetGroupResponse
type GetGroupOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the group.
	//
	// Group is a required field
	Group *Group `type:"structure" required:"true"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `min:"1" type:"string"`

	// A list of users in the group.
	//
	// Users is a required field
	Users []User `type:"list" required:"true"`
}

// String returns the string representation
func (s GetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetGroup = "GetGroup"

// GetGroupRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Returns a list of IAM users that are in the specified IAM group. You can
// paginate the results using the MaxItems and Marker parameters.
//
//    // Example sending a request using GetGroupRequest.
//    req := client.GetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetGroup
func (c *Client) GetGroupRequest(input *GetGroupInput) GetGroupRequest {
	op := &aws.Operation{
		Name:       opGetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &GetGroupInput{}
	}

	req := c.newRequest(op, input, &GetGroupOutput{})
	return GetGroupRequest{Request: req, Input: input, Copy: c.GetGroupRequest}
}

// GetGroupRequest is the request type for the
// GetGroup API operation.
type GetGroupRequest struct {
	*aws.Request
	Input *GetGroupInput
	Copy  func(*GetGroupInput) GetGroupRequest
}

// Send marshals and sends the GetGroup API request.
func (r GetGroupRequest) Send(ctx context.Context) (*GetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGroupResponse{
		GetGroupOutput: r.Request.Data.(*GetGroupOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetGroupRequestPaginator returns a paginator for GetGroup.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetGroupRequest(input)
//   p := iam.NewGetGroupRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetGroupPaginator(req GetGroupRequest) GetGroupPaginator {
	return GetGroupPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetGroupInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetGroupPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetGroupPaginator struct {
	aws.Pager
}

func (p *GetGroupPaginator) CurrentPage() *GetGroupOutput {
	return p.Pager.CurrentPage().(*GetGroupOutput)
}

// GetGroupResponse is the response type for the
// GetGroup API operation.
type GetGroupResponse struct {
	*GetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGroup request.
func (r *GetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}