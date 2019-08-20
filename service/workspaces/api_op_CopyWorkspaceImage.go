// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/CopyWorkspaceImageRequest
type CopyWorkspaceImageInput struct {
	_ struct{} `type:"structure"`

	// A description of the image.
	Description *string `min:"1" type:"string"`

	// The name of the image.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The identifier of the source image.
	//
	// SourceImageId is a required field
	SourceImageId *string `type:"string" required:"true"`

	// The identifier of the source Region.
	//
	// SourceRegion is a required field
	SourceRegion *string `min:"1" type:"string" required:"true"`

	// The tags for the image.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CopyWorkspaceImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyWorkspaceImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyWorkspaceImageInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SourceImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceImageId"))
	}

	if s.SourceRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceRegion"))
	}
	if s.SourceRegion != nil && len(*s.SourceRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SourceRegion", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/CopyWorkspaceImageResult
type CopyWorkspaceImageOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the image.
	ImageId *string `type:"string"`
}

// String returns the string representation
func (s CopyWorkspaceImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopyWorkspaceImage = "CopyWorkspaceImage"

// CopyWorkspaceImageRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Copies the specified image from the specified Region to the current Region.
//
//    // Example sending a request using CopyWorkspaceImageRequest.
//    req := client.CopyWorkspaceImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/CopyWorkspaceImage
func (c *Client) CopyWorkspaceImageRequest(input *CopyWorkspaceImageInput) CopyWorkspaceImageRequest {
	op := &aws.Operation{
		Name:       opCopyWorkspaceImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyWorkspaceImageInput{}
	}

	req := c.newRequest(op, input, &CopyWorkspaceImageOutput{})
	return CopyWorkspaceImageRequest{Request: req, Input: input, Copy: c.CopyWorkspaceImageRequest}
}

// CopyWorkspaceImageRequest is the request type for the
// CopyWorkspaceImage API operation.
type CopyWorkspaceImageRequest struct {
	*aws.Request
	Input *CopyWorkspaceImageInput
	Copy  func(*CopyWorkspaceImageInput) CopyWorkspaceImageRequest
}

// Send marshals and sends the CopyWorkspaceImage API request.
func (r CopyWorkspaceImageRequest) Send(ctx context.Context) (*CopyWorkspaceImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyWorkspaceImageResponse{
		CopyWorkspaceImageOutput: r.Request.Data.(*CopyWorkspaceImageOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyWorkspaceImageResponse is the response type for the
// CopyWorkspaceImage API operation.
type CopyWorkspaceImageResponse struct {
	*CopyWorkspaceImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyWorkspaceImage request.
func (r *CopyWorkspaceImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}