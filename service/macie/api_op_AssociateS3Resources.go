// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateS3ResourcesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Amazon Macie member account whose resources you want to associate
	// with Macie.
	MemberAccountId *string `locationName:"memberAccountId" type:"string"`

	// The S3 resources that you want to associate with Amazon Macie for monitoring
	// and data classification.
	//
	// S3Resources is a required field
	S3Resources []S3ResourceClassification `locationName:"s3Resources" type:"list" required:"true"`
}

// String returns the string representation
func (s AssociateS3ResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateS3ResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateS3ResourcesInput"}

	if s.S3Resources == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3Resources"))
	}
	if s.S3Resources != nil {
		for i, v := range s.S3Resources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "S3Resources", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateS3ResourcesOutput struct {
	_ struct{} `type:"structure"`

	// S3 resources that couldn't be associated with Amazon Macie. An error code
	// and an error message are provided for each failed item.
	FailedS3Resources []FailedS3Resource `locationName:"failedS3Resources" type:"list"`
}

// String returns the string representation
func (s AssociateS3ResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateS3Resources = "AssociateS3Resources"

// AssociateS3ResourcesRequest returns a request value for making API operation for
// Amazon Macie.
//
// Associates specified S3 resources with Amazon Macie for monitoring and data
// classification. If memberAccountId isn't specified, the action associates
// specified S3 resources with Macie for the current master account. If memberAccountId
// is specified, the action associates specified S3 resources with Macie for
// the specified member account.
//
//    // Example sending a request using AssociateS3ResourcesRequest.
//    req := client.AssociateS3ResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie-2017-12-19/AssociateS3Resources
func (c *Client) AssociateS3ResourcesRequest(input *AssociateS3ResourcesInput) AssociateS3ResourcesRequest {
	op := &aws.Operation{
		Name:       opAssociateS3Resources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateS3ResourcesInput{}
	}

	req := c.newRequest(op, input, &AssociateS3ResourcesOutput{})
	return AssociateS3ResourcesRequest{Request: req, Input: input, Copy: c.AssociateS3ResourcesRequest}
}

// AssociateS3ResourcesRequest is the request type for the
// AssociateS3Resources API operation.
type AssociateS3ResourcesRequest struct {
	*aws.Request
	Input *AssociateS3ResourcesInput
	Copy  func(*AssociateS3ResourcesInput) AssociateS3ResourcesRequest
}

// Send marshals and sends the AssociateS3Resources API request.
func (r AssociateS3ResourcesRequest) Send(ctx context.Context) (*AssociateS3ResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateS3ResourcesResponse{
		AssociateS3ResourcesOutput: r.Request.Data.(*AssociateS3ResourcesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateS3ResourcesResponse is the response type for the
// AssociateS3Resources API operation.
type AssociateS3ResourcesResponse struct {
	*AssociateS3ResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateS3Resources request.
func (r *AssociateS3ResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
