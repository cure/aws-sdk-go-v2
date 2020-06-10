// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateDashboardInput struct {
	_ struct{} `type:"structure"`

	// A unique case-sensitive identifier that you can provide to ensure the idempotency
	// of the request. Don't reuse this client token if a new idempotent request
	// is required.
	ClientToken *string `locationName:"clientToken" min:"36" type:"string" idempotencyToken:"true"`

	// The new dashboard definition, as specified in a JSON literal. For detailed
	// information, see Creating Dashboards (CLI) (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// DashboardDefinition is a required field
	DashboardDefinition *string `locationName:"dashboardDefinition" type:"string" required:"true"`

	// A new description for the dashboard.
	DashboardDescription *string `locationName:"dashboardDescription" min:"1" type:"string"`

	// The ID of the dashboard to update.
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"dashboardId" min:"36" type:"string" required:"true"`

	// A new friendly name for the dashboard.
	//
	// DashboardName is a required field
	DashboardName *string `locationName:"dashboardName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateDashboardInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDashboardInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDashboardInput"}
	if s.ClientToken != nil && len(*s.ClientToken) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 36))
	}

	if s.DashboardDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardDefinition"))
	}
	if s.DashboardDescription != nil && len(*s.DashboardDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardDescription", 1))
	}

	if s.DashboardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardId"))
	}
	if s.DashboardId != nil && len(*s.DashboardId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardId", 36))
	}

	if s.DashboardName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardName"))
	}
	if s.DashboardName != nil && len(*s.DashboardName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDashboardInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardDefinition != nil {
		v := *s.DashboardDefinition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dashboardDefinition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardDescription != nil {
		v := *s.DashboardDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dashboardDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardName != nil {
		v := *s.DashboardName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dashboardName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "dashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDashboardOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateDashboardOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDashboardOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateDashboard = "UpdateDashboard"

// UpdateDashboardRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Updates an AWS IoT SiteWise Monitor dashboard.
//
//    // Example sending a request using UpdateDashboardRequest.
//    req := client.UpdateDashboardRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/UpdateDashboard
func (c *Client) UpdateDashboardRequest(input *UpdateDashboardInput) UpdateDashboardRequest {
	op := &aws.Operation{
		Name:       opUpdateDashboard,
		HTTPMethod: "PUT",
		HTTPPath:   "/dashboards/{dashboardId}",
	}

	if input == nil {
		input = &UpdateDashboardInput{}
	}

	req := c.newRequest(op, input, &UpdateDashboardOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("monitor.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return UpdateDashboardRequest{Request: req, Input: input, Copy: c.UpdateDashboardRequest}
}

// UpdateDashboardRequest is the request type for the
// UpdateDashboard API operation.
type UpdateDashboardRequest struct {
	*aws.Request
	Input *UpdateDashboardInput
	Copy  func(*UpdateDashboardInput) UpdateDashboardRequest
}

// Send marshals and sends the UpdateDashboard API request.
func (r UpdateDashboardRequest) Send(ctx context.Context) (*UpdateDashboardResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDashboardResponse{
		UpdateDashboardOutput: r.Request.Data.(*UpdateDashboardOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDashboardResponse is the response type for the
// UpdateDashboard API operation.
type UpdateDashboardResponse struct {
	*UpdateDashboardOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDashboard request.
func (r *UpdateDashboardResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}