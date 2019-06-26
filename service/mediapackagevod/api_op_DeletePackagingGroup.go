// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeletePackagingGroupRequest
type DeletePackagingGroupInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePackagingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePackagingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePackagingGroupInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePackagingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeletePackagingGroupResponse
type DeletePackagingGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePackagingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeletePackagingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeletePackagingGroup = "DeletePackagingGroup"

// DeletePackagingGroupRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Deletes a MediaPackage VOD PackagingGroup resource.
//
//    // Example sending a request using DeletePackagingGroupRequest.
//    req := client.DeletePackagingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DeletePackagingGroup
func (c *Client) DeletePackagingGroupRequest(input *DeletePackagingGroupInput) DeletePackagingGroupRequest {
	op := &aws.Operation{
		Name:       opDeletePackagingGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/packaging_groups/{id}",
	}

	if input == nil {
		input = &DeletePackagingGroupInput{}
	}

	req := c.newRequest(op, input, &DeletePackagingGroupOutput{})
	return DeletePackagingGroupRequest{Request: req, Input: input, Copy: c.DeletePackagingGroupRequest}
}

// DeletePackagingGroupRequest is the request type for the
// DeletePackagingGroup API operation.
type DeletePackagingGroupRequest struct {
	*aws.Request
	Input *DeletePackagingGroupInput
	Copy  func(*DeletePackagingGroupInput) DeletePackagingGroupRequest
}

// Send marshals and sends the DeletePackagingGroup API request.
func (r DeletePackagingGroupRequest) Send(ctx context.Context) (*DeletePackagingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePackagingGroupResponse{
		DeletePackagingGroupOutput: r.Request.Data.(*DeletePackagingGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePackagingGroupResponse is the response type for the
// DeletePackagingGroup API operation.
type DeletePackagingGroupResponse struct {
	*DeletePackagingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePackagingGroup request.
func (r *DeletePackagingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}