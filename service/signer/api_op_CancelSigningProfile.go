// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type CancelSigningProfileInput struct {
	_ struct{} `type:"structure"`

	// The name of the signing profile to be canceled.
	//
	// ProfileName is a required field
	ProfileName *string `location:"uri" locationName:"profileName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelSigningProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelSigningProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelSigningProfileInput"}

	if s.ProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfileName"))
	}
	if s.ProfileName != nil && len(*s.ProfileName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfileName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelSigningProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProfileName != nil {
		v := *s.ProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CancelSigningProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelSigningProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelSigningProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCancelSigningProfile = "CancelSigningProfile"

// CancelSigningProfileRequest returns a request value for making API operation for
// AWS Signer.
//
// Changes the state of an ACTIVE signing profile to CANCELED. A canceled profile
// is still viewable with the ListSigningProfiles operation, but it cannot perform
// new signing jobs, and is deleted two years after cancelation.
//
//    // Example sending a request using CancelSigningProfileRequest.
//    req := client.CancelSigningProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/CancelSigningProfile
func (c *Client) CancelSigningProfileRequest(input *CancelSigningProfileInput) CancelSigningProfileRequest {
	op := &aws.Operation{
		Name:       opCancelSigningProfile,
		HTTPMethod: "DELETE",
		HTTPPath:   "/signing-profiles/{profileName}",
	}

	if input == nil {
		input = &CancelSigningProfileInput{}
	}

	req := c.newRequest(op, input, &CancelSigningProfileOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CancelSigningProfileRequest{Request: req, Input: input, Copy: c.CancelSigningProfileRequest}
}

// CancelSigningProfileRequest is the request type for the
// CancelSigningProfile API operation.
type CancelSigningProfileRequest struct {
	*aws.Request
	Input *CancelSigningProfileInput
	Copy  func(*CancelSigningProfileInput) CancelSigningProfileRequest
}

// Send marshals and sends the CancelSigningProfile API request.
func (r CancelSigningProfileRequest) Send(ctx context.Context) (*CancelSigningProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelSigningProfileResponse{
		CancelSigningProfileOutput: r.Request.Data.(*CancelSigningProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelSigningProfileResponse is the response type for the
// CancelSigningProfile API operation.
type CancelSigningProfileResponse struct {
	*CancelSigningProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelSigningProfile request.
func (r *CancelSigningProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
