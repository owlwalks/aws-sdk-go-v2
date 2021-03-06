// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddApplicationReferenceDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The name of an existing application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The version of the application for which you are adding the reference data
	// source. You can use the DescribeApplication operation to get the current
	// application version. If the version specified is not the current version,
	// the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The reference data source can be an object in your Amazon S3 bucket. Kinesis
	// Data Analytics reads the object and copies the data into the in-application
	// table that is created. You provide an S3 bucket, object key name, and the
	// resulting in-application table that is created.
	//
	// ReferenceDataSource is a required field
	ReferenceDataSource *ReferenceDataSource `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddApplicationReferenceDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationReferenceDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationReferenceDataSourceInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if s.ReferenceDataSource == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReferenceDataSource"))
	}
	if s.ReferenceDataSource != nil {
		if err := s.ReferenceDataSource.Validate(); err != nil {
			invalidParams.AddNested("ReferenceDataSource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddApplicationReferenceDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string `min:"1" type:"string"`

	// The updated application version ID. Amazon Kinesis Data Analytics increments
	// this ID when the application is updated.
	ApplicationVersionId *int64 `min:"1" type:"long"`

	// Describes reference data sources configured for the application.
	ReferenceDataSourceDescriptions []ReferenceDataSourceDescription `type:"list"`
}

// String returns the string representation
func (s AddApplicationReferenceDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddApplicationReferenceDataSource = "AddApplicationReferenceDataSource"

// AddApplicationReferenceDataSourceRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Adds a reference data source to an existing SQL-based Amazon Kinesis Data
// Analytics application.
//
// Kinesis Data Analytics reads reference data (that is, an Amazon S3 object)
// and creates an in-application table within your application. In the request,
// you provide the source (S3 bucket name and object key name), name of the
// in-application table to create, and the necessary mapping information that
// describes how data in an Amazon S3 object maps to columns in the resulting
// in-application table.
//
//    // Example sending a request using AddApplicationReferenceDataSourceRequest.
//    req := client.AddApplicationReferenceDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationReferenceDataSource
func (c *Client) AddApplicationReferenceDataSourceRequest(input *AddApplicationReferenceDataSourceInput) AddApplicationReferenceDataSourceRequest {
	op := &aws.Operation{
		Name:       opAddApplicationReferenceDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddApplicationReferenceDataSourceInput{}
	}

	req := c.newRequest(op, input, &AddApplicationReferenceDataSourceOutput{})
	return AddApplicationReferenceDataSourceRequest{Request: req, Input: input, Copy: c.AddApplicationReferenceDataSourceRequest}
}

// AddApplicationReferenceDataSourceRequest is the request type for the
// AddApplicationReferenceDataSource API operation.
type AddApplicationReferenceDataSourceRequest struct {
	*aws.Request
	Input *AddApplicationReferenceDataSourceInput
	Copy  func(*AddApplicationReferenceDataSourceInput) AddApplicationReferenceDataSourceRequest
}

// Send marshals and sends the AddApplicationReferenceDataSource API request.
func (r AddApplicationReferenceDataSourceRequest) Send(ctx context.Context) (*AddApplicationReferenceDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddApplicationReferenceDataSourceResponse{
		AddApplicationReferenceDataSourceOutput: r.Request.Data.(*AddApplicationReferenceDataSourceOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddApplicationReferenceDataSourceResponse is the response type for the
// AddApplicationReferenceDataSource API operation.
type AddApplicationReferenceDataSourceResponse struct {
	*AddApplicationReferenceDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddApplicationReferenceDataSource request.
func (r *AddApplicationReferenceDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
