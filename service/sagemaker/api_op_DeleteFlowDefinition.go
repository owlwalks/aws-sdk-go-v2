// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFlowDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The name of the flow definition you are deleting.
	//
	// FlowDefinitionName is a required field
	FlowDefinitionName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFlowDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFlowDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFlowDefinitionInput"}

	if s.FlowDefinitionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowDefinitionName"))
	}
	if s.FlowDefinitionName != nil && len(*s.FlowDefinitionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FlowDefinitionName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFlowDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFlowDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFlowDefinition = "DeleteFlowDefinition"

// DeleteFlowDefinitionRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Deletes the specified flow definition.
//
//    // Example sending a request using DeleteFlowDefinitionRequest.
//    req := client.DeleteFlowDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteFlowDefinition
func (c *Client) DeleteFlowDefinitionRequest(input *DeleteFlowDefinitionInput) DeleteFlowDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteFlowDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFlowDefinitionInput{}
	}

	req := c.newRequest(op, input, &DeleteFlowDefinitionOutput{})
	return DeleteFlowDefinitionRequest{Request: req, Input: input, Copy: c.DeleteFlowDefinitionRequest}
}

// DeleteFlowDefinitionRequest is the request type for the
// DeleteFlowDefinition API operation.
type DeleteFlowDefinitionRequest struct {
	*aws.Request
	Input *DeleteFlowDefinitionInput
	Copy  func(*DeleteFlowDefinitionInput) DeleteFlowDefinitionRequest
}

// Send marshals and sends the DeleteFlowDefinition API request.
func (r DeleteFlowDefinitionRequest) Send(ctx context.Context) (*DeleteFlowDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFlowDefinitionResponse{
		DeleteFlowDefinitionOutput: r.Request.Data.(*DeleteFlowDefinitionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFlowDefinitionResponse is the response type for the
// DeleteFlowDefinition API operation.
type DeleteFlowDefinitionResponse struct {
	*DeleteFlowDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFlowDefinition request.
func (r *DeleteFlowDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
