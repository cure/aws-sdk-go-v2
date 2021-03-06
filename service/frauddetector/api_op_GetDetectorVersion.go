// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDetectorVersionInput struct {
	_ struct{} `type:"structure"`

	// The detector ID.
	//
	// DetectorId is a required field
	DetectorId *string `locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The detector version ID.
	//
	// DetectorVersionId is a required field
	DetectorVersionId *string `locationName:"detectorVersionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDetectorVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDetectorVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDetectorVersionInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.DetectorVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorVersionId"))
	}
	if s.DetectorVersionId != nil && len(*s.DetectorVersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorVersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDetectorVersionOutput struct {
	_ struct{} `type:"structure"`

	// The timestamp when the detector version was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The detector version description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The detector ID.
	DetectorId *string `locationName:"detectorId" min:"1" type:"string"`

	// The detector version ID.
	DetectorVersionId *string `locationName:"detectorVersionId" min:"1" type:"string"`

	// The Amazon SageMaker model endpoints included in the detector version.
	ExternalModelEndpoints []string `locationName:"externalModelEndpoints" type:"list"`

	// The timestamp when the detector version was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The model versions included in the detector version.
	ModelVersions []ModelVersion `locationName:"modelVersions" type:"list"`

	// The execution mode of the rule in the dectector
	//
	// FIRST_MATCHED indicates that Amazon Fraud Detector evaluates rules sequentially,
	// first to last, stopping at the first matched rule. Amazon Fraud dectector
	// then provides the outcomes for that single rule.
	//
	// ALL_MATCHED indicates that Amazon Fraud Detector evaluates all rules and
	// returns the outcomes for all matched rules. You can define and edit the rule
	// mode at the detector version level, when it is in draft status.
	RuleExecutionMode RuleExecutionMode `locationName:"ruleExecutionMode" type:"string" enum:"true"`

	// The rules included in the detector version.
	Rules []Rule `locationName:"rules" type:"list"`

	// The status of the detector version.
	Status DetectorVersionStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetDetectorVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDetectorVersion = "GetDetectorVersion"

// GetDetectorVersionRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Gets a particular detector version.
//
//    // Example sending a request using GetDetectorVersionRequest.
//    req := client.GetDetectorVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/GetDetectorVersion
func (c *Client) GetDetectorVersionRequest(input *GetDetectorVersionInput) GetDetectorVersionRequest {
	op := &aws.Operation{
		Name:       opGetDetectorVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDetectorVersionInput{}
	}

	req := c.newRequest(op, input, &GetDetectorVersionOutput{})

	return GetDetectorVersionRequest{Request: req, Input: input, Copy: c.GetDetectorVersionRequest}
}

// GetDetectorVersionRequest is the request type for the
// GetDetectorVersion API operation.
type GetDetectorVersionRequest struct {
	*aws.Request
	Input *GetDetectorVersionInput
	Copy  func(*GetDetectorVersionInput) GetDetectorVersionRequest
}

// Send marshals and sends the GetDetectorVersion API request.
func (r GetDetectorVersionRequest) Send(ctx context.Context) (*GetDetectorVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDetectorVersionResponse{
		GetDetectorVersionOutput: r.Request.Data.(*GetDetectorVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDetectorVersionResponse is the response type for the
// GetDetectorVersion API operation.
type GetDetectorVersionResponse struct {
	*GetDetectorVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDetectorVersion request.
func (r *GetDetectorVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
