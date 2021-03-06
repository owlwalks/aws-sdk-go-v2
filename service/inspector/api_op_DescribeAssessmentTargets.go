// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAssessmentTargetsInput struct {
	_ struct{} `type:"structure"`

	// The ARNs that specifies the assessment targets that you want to describe.
	//
	// AssessmentTargetArns is a required field
	AssessmentTargetArns []string `locationName:"assessmentTargetArns" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeAssessmentTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAssessmentTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAssessmentTargetsInput"}

	if s.AssessmentTargetArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentTargetArns"))
	}
	if s.AssessmentTargetArns != nil && len(s.AssessmentTargetArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentTargetArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAssessmentTargetsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the assessment targets.
	//
	// AssessmentTargets is a required field
	AssessmentTargets []AssessmentTarget `locationName:"assessmentTargets" type:"list" required:"true"`

	// Assessment target details that cannot be described. An error code is provided
	// for each failed item.
	//
	// FailedItems is a required field
	FailedItems map[string]FailedItemDetails `locationName:"failedItems" type:"map" required:"true"`
}

// String returns the string representation
func (s DescribeAssessmentTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAssessmentTargets = "DescribeAssessmentTargets"

// DescribeAssessmentTargetsRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Describes the assessment targets that are specified by the ARNs of the assessment
// targets.
//
//    // Example sending a request using DescribeAssessmentTargetsRequest.
//    req := client.DescribeAssessmentTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DescribeAssessmentTargets
func (c *Client) DescribeAssessmentTargetsRequest(input *DescribeAssessmentTargetsInput) DescribeAssessmentTargetsRequest {
	op := &aws.Operation{
		Name:       opDescribeAssessmentTargets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAssessmentTargetsInput{}
	}

	req := c.newRequest(op, input, &DescribeAssessmentTargetsOutput{})
	return DescribeAssessmentTargetsRequest{Request: req, Input: input, Copy: c.DescribeAssessmentTargetsRequest}
}

// DescribeAssessmentTargetsRequest is the request type for the
// DescribeAssessmentTargets API operation.
type DescribeAssessmentTargetsRequest struct {
	*aws.Request
	Input *DescribeAssessmentTargetsInput
	Copy  func(*DescribeAssessmentTargetsInput) DescribeAssessmentTargetsRequest
}

// Send marshals and sends the DescribeAssessmentTargets API request.
func (r DescribeAssessmentTargetsRequest) Send(ctx context.Context) (*DescribeAssessmentTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAssessmentTargetsResponse{
		DescribeAssessmentTargetsOutput: r.Request.Data.(*DescribeAssessmentTargetsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAssessmentTargetsResponse is the response type for the
// DescribeAssessmentTargets API operation.
type DescribeAssessmentTargetsResponse struct {
	*DescribeAssessmentTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAssessmentTargets request.
func (r *DescribeAssessmentTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
