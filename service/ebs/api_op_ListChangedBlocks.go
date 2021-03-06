// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ebs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListChangedBlocksInput struct {
	_ struct{} `type:"structure"`

	// The ID of the first snapshot to use for the comparison.
	//
	// The FirstSnapshotID parameter must be specified with a SecondSnapshotId parameter;
	// otherwise, an error occurs.
	FirstSnapshotId *string `location:"querystring" locationName:"firstSnapshotId" min:"1" type:"string"`

	// The number of results to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"100" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `location:"querystring" locationName:"pageToken" type:"string"`

	// The ID of the second snapshot to use for the comparison.
	//
	// The SecondSnapshotId parameter must be specified with a FirstSnapshotID parameter;
	// otherwise, an error occurs.
	//
	// SecondSnapshotId is a required field
	SecondSnapshotId *string `location:"uri" locationName:"secondSnapshotId" min:"1" type:"string" required:"true"`

	// The block index from which the comparison should start.
	//
	// The list in the response will start from this block index or the next valid
	// block index in the snapshots.
	StartingBlockIndex *int64 `location:"querystring" locationName:"startingBlockIndex" type:"integer"`
}

// String returns the string representation
func (s ListChangedBlocksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListChangedBlocksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListChangedBlocksInput"}
	if s.FirstSnapshotId != nil && len(*s.FirstSnapshotId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FirstSnapshotId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 100 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 100))
	}

	if s.SecondSnapshotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecondSnapshotId"))
	}
	if s.SecondSnapshotId != nil && len(*s.SecondSnapshotId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecondSnapshotId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListChangedBlocksInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SecondSnapshotId != nil {
		v := *s.SecondSnapshotId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "secondSnapshotId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FirstSnapshotId != nil {
		v := *s.FirstSnapshotId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "firstSnapshotId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "pageToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartingBlockIndex != nil {
		v := *s.StartingBlockIndex

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "startingBlockIndex", protocol.Int64Value(v), metadata)
	}
	return nil
}

type ListChangedBlocksOutput struct {
	_ struct{} `type:"structure"`

	// The size of the block.
	BlockSize *int64 `type:"integer"`

	// An array of objects containing information about the changed blocks.
	ChangedBlocks []ChangedBlock `type:"list"`

	// The time when the BlockToken expires.
	ExpiryTime *time.Time `type:"timestamp"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `type:"string"`

	// The size of the volume in GB.
	VolumeSize *int64 `type:"long"`
}

// String returns the string representation
func (s ListChangedBlocksOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListChangedBlocksOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BlockSize != nil {
		v := *s.BlockSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BlockSize", protocol.Int64Value(v), metadata)
	}
	if s.ChangedBlocks != nil {
		v := s.ChangedBlocks

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ChangedBlocks", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ExpiryTime != nil {
		v := *s.ExpiryTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExpiryTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VolumeSize != nil {
		v := *s.VolumeSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VolumeSize", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opListChangedBlocks = "ListChangedBlocks"

// ListChangedBlocksRequest returns a request value for making API operation for
// Amazon Elastic Block Store.
//
// Returns the block indexes and block tokens for blocks that are different
// between two Amazon Elastic Block Store snapshots of the same volume/snapshot
// lineage.
//
//    // Example sending a request using ListChangedBlocksRequest.
//    req := client.ListChangedBlocksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ebs-2019-11-02/ListChangedBlocks
func (c *Client) ListChangedBlocksRequest(input *ListChangedBlocksInput) ListChangedBlocksRequest {
	op := &aws.Operation{
		Name:       opListChangedBlocks,
		HTTPMethod: "GET",
		HTTPPath:   "/snapshots/{secondSnapshotId}/changedblocks",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListChangedBlocksInput{}
	}

	req := c.newRequest(op, input, &ListChangedBlocksOutput{})

	return ListChangedBlocksRequest{Request: req, Input: input, Copy: c.ListChangedBlocksRequest}
}

// ListChangedBlocksRequest is the request type for the
// ListChangedBlocks API operation.
type ListChangedBlocksRequest struct {
	*aws.Request
	Input *ListChangedBlocksInput
	Copy  func(*ListChangedBlocksInput) ListChangedBlocksRequest
}

// Send marshals and sends the ListChangedBlocks API request.
func (r ListChangedBlocksRequest) Send(ctx context.Context) (*ListChangedBlocksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListChangedBlocksResponse{
		ListChangedBlocksOutput: r.Request.Data.(*ListChangedBlocksOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListChangedBlocksRequestPaginator returns a paginator for ListChangedBlocks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListChangedBlocksRequest(input)
//   p := ebs.NewListChangedBlocksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListChangedBlocksPaginator(req ListChangedBlocksRequest) ListChangedBlocksPaginator {
	return ListChangedBlocksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListChangedBlocksInput
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

// ListChangedBlocksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListChangedBlocksPaginator struct {
	aws.Pager
}

func (p *ListChangedBlocksPaginator) CurrentPage() *ListChangedBlocksOutput {
	return p.Pager.CurrentPage().(*ListChangedBlocksOutput)
}

// ListChangedBlocksResponse is the response type for the
// ListChangedBlocks API operation.
type ListChangedBlocksResponse struct {
	*ListChangedBlocksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListChangedBlocks request.
func (r *ListChangedBlocksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
