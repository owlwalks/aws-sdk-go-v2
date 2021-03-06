// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopReplicationTaskInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name(ARN) of the replication task to be stopped.
	//
	// ReplicationTaskArn is a required field
	ReplicationTaskArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopReplicationTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopReplicationTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopReplicationTaskInput"}

	if s.ReplicationTaskArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationTaskArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopReplicationTaskOutput struct {
	_ struct{} `type:"structure"`

	// The replication task stopped.
	ReplicationTask *ReplicationTask `type:"structure"`
}

// String returns the string representation
func (s StopReplicationTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopReplicationTask = "StopReplicationTask"

// StopReplicationTaskRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Stops the replication task.
//
//    // Example sending a request using StopReplicationTaskRequest.
//    req := client.StopReplicationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/StopReplicationTask
func (c *Client) StopReplicationTaskRequest(input *StopReplicationTaskInput) StopReplicationTaskRequest {
	op := &aws.Operation{
		Name:       opStopReplicationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopReplicationTaskInput{}
	}

	req := c.newRequest(op, input, &StopReplicationTaskOutput{})
	return StopReplicationTaskRequest{Request: req, Input: input, Copy: c.StopReplicationTaskRequest}
}

// StopReplicationTaskRequest is the request type for the
// StopReplicationTask API operation.
type StopReplicationTaskRequest struct {
	*aws.Request
	Input *StopReplicationTaskInput
	Copy  func(*StopReplicationTaskInput) StopReplicationTaskRequest
}

// Send marshals and sends the StopReplicationTask API request.
func (r StopReplicationTaskRequest) Send(ctx context.Context) (*StopReplicationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopReplicationTaskResponse{
		StopReplicationTaskOutput: r.Request.Data.(*StopReplicationTaskOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopReplicationTaskResponse is the response type for the
// StopReplicationTask API operation.
type StopReplicationTaskResponse struct {
	*StopReplicationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopReplicationTask request.
func (r *StopReplicationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
