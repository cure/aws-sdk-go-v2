// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/ModifyWorkspacePropertiesRequest
type ModifyWorkspacePropertiesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	//
	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" required:"true"`

	// The properties of the WorkSpace.
	//
	// WorkspaceProperties is a required field
	WorkspaceProperties *WorkspaceProperties `type:"structure" required:"true"`
}

// String returns the string representation
func (s ModifyWorkspacePropertiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyWorkspacePropertiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyWorkspacePropertiesInput"}

	if s.WorkspaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceId"))
	}

	if s.WorkspaceProperties == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceProperties"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/ModifyWorkspacePropertiesResult
type ModifyWorkspacePropertiesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyWorkspacePropertiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyWorkspaceProperties = "ModifyWorkspaceProperties"

// ModifyWorkspacePropertiesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Modifies the specified WorkSpace properties.
//
//    // Example sending a request using ModifyWorkspacePropertiesRequest.
//    req := client.ModifyWorkspacePropertiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/ModifyWorkspaceProperties
func (c *Client) ModifyWorkspacePropertiesRequest(input *ModifyWorkspacePropertiesInput) ModifyWorkspacePropertiesRequest {
	op := &aws.Operation{
		Name:       opModifyWorkspaceProperties,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyWorkspacePropertiesInput{}
	}

	req := c.newRequest(op, input, &ModifyWorkspacePropertiesOutput{})
	return ModifyWorkspacePropertiesRequest{Request: req, Input: input, Copy: c.ModifyWorkspacePropertiesRequest}
}

// ModifyWorkspacePropertiesRequest is the request type for the
// ModifyWorkspaceProperties API operation.
type ModifyWorkspacePropertiesRequest struct {
	*aws.Request
	Input *ModifyWorkspacePropertiesInput
	Copy  func(*ModifyWorkspacePropertiesInput) ModifyWorkspacePropertiesRequest
}

// Send marshals and sends the ModifyWorkspaceProperties API request.
func (r ModifyWorkspacePropertiesRequest) Send(ctx context.Context) (*ModifyWorkspacePropertiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyWorkspacePropertiesResponse{
		ModifyWorkspacePropertiesOutput: r.Request.Data.(*ModifyWorkspacePropertiesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyWorkspacePropertiesResponse is the response type for the
// ModifyWorkspaceProperties API operation.
type ModifyWorkspacePropertiesResponse struct {
	*ModifyWorkspacePropertiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyWorkspaceProperties request.
func (r *ModifyWorkspacePropertiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}