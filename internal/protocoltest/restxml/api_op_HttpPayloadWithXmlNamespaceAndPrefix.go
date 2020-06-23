// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The following example serializes a payload that uses an XML namespace.
func (c *Client) HttpPayloadWithXmlNamespaceAndPrefix(ctx context.Context, params *HttpPayloadWithXmlNamespaceAndPrefixInput, optFns ...func(*Options)) (*HttpPayloadWithXmlNamespaceAndPrefixOutput, error) {
	stack := middleware.NewStack("HttpPayloadWithXmlNamespaceAndPrefix", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "HttpPayloadWithXmlNamespaceAndPrefix",
	}, middleware.Before)
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "HttpPayloadWithXmlNamespaceAndPrefix",
			Err:           err,
		}
	}
	out := result.(*HttpPayloadWithXmlNamespaceAndPrefixOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithXmlNamespaceAndPrefixInput struct {
	Nested *types.PayloadWithXmlNamespaceAndPrefix
}

type HttpPayloadWithXmlNamespaceAndPrefixOutput struct {
	Nested *types.PayloadWithXmlNamespaceAndPrefix

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
