// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssooidc

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StartDeviceAuthorizationInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier string for the client that is registered with AWS SSO.
	// This value should come from the persisted result of the RegisterClient API
	// operation.
	//
	// ClientId is a required field
	ClientId *string `locationName:"clientId" type:"string" required:"true"`

	// A secret string that is generated for the client. This value should come
	// from the persisted result of the RegisterClient API operation.
	//
	// ClientSecret is a required field
	ClientSecret *string `locationName:"clientSecret" type:"string" required:"true"`

	// The URL for the AWS SSO user portal. For more information, see Using the
	// User Portal (https://docs.aws.amazon.com/singlesignon/latest/userguide/using-the-portal.html)
	// in the AWS Single Sign-On User Guide.
	//
	// StartUrl is a required field
	StartUrl *string `locationName:"startUrl" type:"string" required:"true"`
}

// String returns the string representation
func (s StartDeviceAuthorizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDeviceAuthorizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDeviceAuthorizationInput"}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}

	if s.ClientSecret == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientSecret"))
	}

	if s.StartUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartDeviceAuthorizationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClientId != nil {
		v := *s.ClientId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientSecret != nil {
		v := *s.ClientSecret

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientSecret", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartUrl != nil {
		v := *s.StartUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartDeviceAuthorizationOutput struct {
	_ struct{} `type:"structure"`

	// The short-lived code that is used by the device when polling for a session
	// token.
	DeviceCode *string `locationName:"deviceCode" type:"string"`

	// Indicates the number of seconds in which the verification code will become
	// invalid.
	ExpiresIn *int64 `locationName:"expiresIn" type:"integer"`

	// Indicates the number of seconds the client must wait between attempts when
	// polling for a session.
	Interval *int64 `locationName:"interval" type:"integer"`

	// A one-time user verification code. This is needed to authorize an in-use
	// device.
	UserCode *string `locationName:"userCode" type:"string"`

	// The URI of the verification page that takes the userCode to authorize the
	// device.
	VerificationUri *string `locationName:"verificationUri" type:"string"`

	// An alternate URL that the client can use to automatically launch a browser.
	// This process skips the manual step in which the user visits the verification
	// page and enters their code.
	VerificationUriComplete *string `locationName:"verificationUriComplete" type:"string"`
}

// String returns the string representation
func (s StartDeviceAuthorizationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartDeviceAuthorizationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeviceCode != nil {
		v := *s.DeviceCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deviceCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExpiresIn != nil {
		v := *s.ExpiresIn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expiresIn", protocol.Int64Value(v), metadata)
	}
	if s.Interval != nil {
		v := *s.Interval

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "interval", protocol.Int64Value(v), metadata)
	}
	if s.UserCode != nil {
		v := *s.UserCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "userCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VerificationUri != nil {
		v := *s.VerificationUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "verificationUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VerificationUriComplete != nil {
		v := *s.VerificationUriComplete

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "verificationUriComplete", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opStartDeviceAuthorization = "StartDeviceAuthorization"

// StartDeviceAuthorizationRequest returns a request value for making API operation for
// AWS SSO OIDC.
//
// Initiates device authorization by requesting a pair of verification codes
// from the authorization service.
//
//    // Example sending a request using StartDeviceAuthorizationRequest.
//    req := client.StartDeviceAuthorizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sso-oidc-2019-06-10/StartDeviceAuthorization
func (c *Client) StartDeviceAuthorizationRequest(input *StartDeviceAuthorizationInput) StartDeviceAuthorizationRequest {
	op := &aws.Operation{
		Name:       opStartDeviceAuthorization,
		HTTPMethod: "POST",
		HTTPPath:   "/device_authorization",
	}

	if input == nil {
		input = &StartDeviceAuthorizationInput{}
	}

	req := c.newRequest(op, input, &StartDeviceAuthorizationOutput{})
	req.Config.Credentials = aws.AnonymousCredentials
	return StartDeviceAuthorizationRequest{Request: req, Input: input, Copy: c.StartDeviceAuthorizationRequest}
}

// StartDeviceAuthorizationRequest is the request type for the
// StartDeviceAuthorization API operation.
type StartDeviceAuthorizationRequest struct {
	*aws.Request
	Input *StartDeviceAuthorizationInput
	Copy  func(*StartDeviceAuthorizationInput) StartDeviceAuthorizationRequest
}

// Send marshals and sends the StartDeviceAuthorization API request.
func (r StartDeviceAuthorizationRequest) Send(ctx context.Context) (*StartDeviceAuthorizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDeviceAuthorizationResponse{
		StartDeviceAuthorizationOutput: r.Request.Data.(*StartDeviceAuthorizationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDeviceAuthorizationResponse is the response type for the
// StartDeviceAuthorization API operation.
type StartDeviceAuthorizationResponse struct {
	*StartDeviceAuthorizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDeviceAuthorization request.
func (r *StartDeviceAuthorizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
