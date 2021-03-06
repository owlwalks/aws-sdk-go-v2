// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstanceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the instance that you want to get information about.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The ID of the service that the instance is associated with.
	//
	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.ServiceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInstanceOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about a specified instance.
	Instance *Instance `type:"structure"`
}

// String returns the string representation
func (s GetInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstance = "GetInstance"

// GetInstanceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Gets information about a specified instance.
//
//    // Example sending a request using GetInstanceRequest.
//    req := client.GetInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/GetInstance
func (c *Client) GetInstanceRequest(input *GetInstanceInput) GetInstanceRequest {
	op := &aws.Operation{
		Name:       opGetInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInstanceInput{}
	}

	req := c.newRequest(op, input, &GetInstanceOutput{})
	return GetInstanceRequest{Request: req, Input: input, Copy: c.GetInstanceRequest}
}

// GetInstanceRequest is the request type for the
// GetInstance API operation.
type GetInstanceRequest struct {
	*aws.Request
	Input *GetInstanceInput
	Copy  func(*GetInstanceInput) GetInstanceRequest
}

// Send marshals and sends the GetInstance API request.
func (r GetInstanceRequest) Send(ctx context.Context) (*GetInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstanceResponse{
		GetInstanceOutput: r.Request.Data.(*GetInstanceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstanceResponse is the response type for the
// GetInstance API operation.
type GetInstanceResponse struct {
	*GetInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstance request.
func (r *GetInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
