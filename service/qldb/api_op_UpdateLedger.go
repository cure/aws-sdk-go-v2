// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateLedgerInput struct {
	_ struct{} `type:"structure"`

	// The flag that prevents a ledger from being deleted by any user. If not provided
	// on ledger creation, this feature is enabled (true) by default.
	//
	// If deletion protection is enabled, you must first disable it before you can
	// delete the ledger using the QLDB API or the AWS Command Line Interface (AWS
	// CLI). You can disable it by calling the UpdateLedger operation to set the
	// flag to false. The QLDB console disables deletion protection for you when
	// you use it to delete a ledger.
	DeletionProtection *bool `type:"boolean"`

	// The name of the ledger.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateLedgerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLedgerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLedgerInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateLedgerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DeletionProtection != nil {
		v := *s.DeletionProtection

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeletionProtection", protocol.BoolValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateLedgerOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the ledger.
	Arn *string `min:"20" type:"string"`

	// The date and time, in epoch time format, when the ledger was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1,
	// 1970 UTC.)
	CreationDateTime *time.Time `type:"timestamp"`

	// The flag that prevents a ledger from being deleted by any user. If not provided
	// on ledger creation, this feature is enabled (true) by default.
	//
	// If deletion protection is enabled, you must first disable it before you can
	// delete the ledger using the QLDB API or the AWS Command Line Interface (AWS
	// CLI). You can disable it by calling the UpdateLedger operation to set the
	// flag to false. The QLDB console disables deletion protection for you when
	// you use it to delete a ledger.
	DeletionProtection *bool `type:"boolean"`

	// The name of the ledger.
	Name *string `min:"1" type:"string"`

	// The current status of the ledger.
	State LedgerState `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateLedgerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateLedgerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationDateTime != nil {
		v := *s.CreationDateTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDateTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.DeletionProtection != nil {
		v := *s.DeletionProtection

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeletionProtection", protocol.BoolValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opUpdateLedger = "UpdateLedger"

// UpdateLedgerRequest returns a request value for making API operation for
// Amazon QLDB.
//
// Updates properties on a ledger.
//
//    // Example sending a request using UpdateLedgerRequest.
//    req := client.UpdateLedgerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/qldb-2019-01-02/UpdateLedger
func (c *Client) UpdateLedgerRequest(input *UpdateLedgerInput) UpdateLedgerRequest {
	op := &aws.Operation{
		Name:       opUpdateLedger,
		HTTPMethod: "PATCH",
		HTTPPath:   "/ledgers/{name}",
	}

	if input == nil {
		input = &UpdateLedgerInput{}
	}

	req := c.newRequest(op, input, &UpdateLedgerOutput{})

	return UpdateLedgerRequest{Request: req, Input: input, Copy: c.UpdateLedgerRequest}
}

// UpdateLedgerRequest is the request type for the
// UpdateLedger API operation.
type UpdateLedgerRequest struct {
	*aws.Request
	Input *UpdateLedgerInput
	Copy  func(*UpdateLedgerInput) UpdateLedgerRequest
}

// Send marshals and sends the UpdateLedger API request.
func (r UpdateLedgerRequest) Send(ctx context.Context) (*UpdateLedgerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLedgerResponse{
		UpdateLedgerOutput: r.Request.Data.(*UpdateLedgerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLedgerResponse is the response type for the
// UpdateLedger API operation.
type UpdateLedgerResponse struct {
	*UpdateLedgerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLedger request.
func (r *UpdateLedgerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
