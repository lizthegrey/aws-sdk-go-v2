// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to create a user pool client.
type CreateUserPoolClientInput struct {
	_ struct{} `type:"structure"`

	// The allowed OAuth flows.
	//
	// Set to code to initiate a code grant flow, which provides an authorization
	// code as the response. This code can be exchanged for access tokens with the
	// token endpoint.
	//
	// Set to implicit to specify that the client should get the access token (and,
	// optionally, ID token, based on scopes) directly.
	//
	// Set to client_credentials to specify that the client should get the access
	// token (and, optionally, ID token, based on scopes) from the token endpoint
	// using a combination of client and client_secret.
	AllowedOAuthFlows []OAuthFlowType `type:"list"`

	// Set to true if the client is allowed to follow the OAuth protocol when interacting
	// with Cognito user pools.
	AllowedOAuthFlowsUserPoolClient *bool `type:"boolean"`

	// The allowed OAuth scopes. Possible values provided by OAuth are: phone, email,
	// openid, and profile. Possible values provided by AWS are: aws.cognito.signin.user.admin.
	// Custom scopes created in Resource Servers are also supported.
	AllowedOAuthScopes []string `type:"list"`

	// The Amazon Pinpoint analytics configuration for collecting metrics for this
	// user pool.
	AnalyticsConfiguration *AnalyticsConfigurationType `type:"structure"`

	// A list of allowed redirect (callback) URLs for the identity providers.
	//
	// A redirect URI must:
	//
	//    * Be an absolute URI.
	//
	//    * Be registered with the authorization server.
	//
	//    * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2).
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	CallbackURLs []string `type:"list"`

	// The client name for the user pool client you would like to create.
	//
	// ClientName is a required field
	ClientName *string `min:"1" type:"string" required:"true"`

	// The default redirect URI. Must be in the CallbackURLs list.
	//
	// A redirect URI must:
	//
	//    * Be an absolute URI.
	//
	//    * Be registered with the authorization server.
	//
	//    * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2).
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectURI *string `min:"1" type:"string"`

	// The authentication flows that are supported by the user pool clients. Flow
	// names without the ALLOW_ prefix are deprecated in favor of new names with
	// the ALLOW_ prefix. Note that values with ALLOW_ prefix cannot be used along
	// with values without ALLOW_ prefix.
	//
	// Valid values include:
	//
	//    * ALLOW_ADMIN_USER_PASSWORD_AUTH: Enable admin based user password authentication
	//    flow ADMIN_USER_PASSWORD_AUTH. This setting replaces the ADMIN_NO_SRP_AUTH
	//    setting. With this authentication flow, Cognito receives the password
	//    in the request instead of using the SRP (Secure Remote Password protocol)
	//    protocol to verify passwords.
	//
	//    * ALLOW_CUSTOM_AUTH: Enable Lambda trigger based authentication.
	//
	//    * ALLOW_USER_PASSWORD_AUTH: Enable user password-based authentication.
	//    In this flow, Cognito receives the password in the request instead of
	//    using the SRP protocol to verify passwords.
	//
	//    * ALLOW_USER_SRP_AUTH: Enable SRP based authentication.
	//
	//    * ALLOW_REFRESH_TOKEN_AUTH: Enable authflow to refresh tokens.
	ExplicitAuthFlows []ExplicitAuthFlowsType `type:"list"`

	// Boolean to specify whether you want to generate a secret for the user pool
	// client being created.
	GenerateSecret *bool `type:"boolean"`

	// A list of allowed logout URLs for the identity providers.
	LogoutURLs []string `type:"list"`

	// Use this setting to choose which errors and responses are returned by Cognito
	// APIs during authentication, account confirmation, and password recovery when
	// the user does not exist in the user pool. When set to ENABLED and the user
	// does not exist, authentication returns an error indicating either the username
	// or password was incorrect, and account confirmation and password recovery
	// return a response indicating a code was sent to a simulated destination.
	// When set to LEGACY, those APIs will return a UserNotFoundException exception
	// if the user does not exist in the user pool.
	//
	// Valid values include:
	//
	//    * ENABLED - This prevents user existence-related errors.
	//
	//    * LEGACY - This represents the old behavior of Cognito where user existence
	//    related errors are not prevented.
	//
	// This setting affects the behavior of following APIs:
	//
	//    * AdminInitiateAuth
	//
	//    * AdminRespondToAuthChallenge
	//
	//    * InitiateAuth
	//
	//    * RespondToAuthChallenge
	//
	//    * ForgotPassword
	//
	//    * ConfirmForgotPassword
	//
	//    * ConfirmSignUp
	//
	//    * ResendConfirmationCode
	//
	// After February 15th 2020, the value of PreventUserExistenceErrors will default
	// to ENABLED for newly created user pool clients if no value is provided.
	PreventUserExistenceErrors PreventUserExistenceErrorTypes `type:"string" enum:"true"`

	// The read attributes.
	ReadAttributes []string `type:"list"`

	// The time limit, in days, after which the refresh token is no longer valid
	// and cannot be used.
	RefreshTokenValidity *int64 `type:"integer"`

	// A list of provider names for the identity providers that are supported on
	// this client. The following are supported: COGNITO, Facebook, Google and LoginWithAmazon.
	SupportedIdentityProviders []string `type:"list"`

	// The user pool ID for the user pool where you want to create a user pool client.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user pool attributes that the app client can write to.
	//
	// If your app client allows users to sign in through an identity provider,
	// this array must include all attributes that are mapped to identity provider
	// attributes. Amazon Cognito updates mapped attributes when users sign in to
	// your application through an identity provider. If your app client lacks write
	// access to a mapped attribute, Amazon Cognito throws an error when it attempts
	// to update the attribute. For more information, see Specifying Identity Provider
	// Attribute Mappings for Your User Pool (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html).
	WriteAttributes []string `type:"list"`
}

// String returns the string representation
func (s CreateUserPoolClientInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserPoolClientInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserPoolClientInput"}

	if s.ClientName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientName"))
	}
	if s.ClientName != nil && len(*s.ClientName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientName", 1))
	}
	if s.DefaultRedirectURI != nil && len(*s.DefaultRedirectURI) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultRedirectURI", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.AnalyticsConfiguration != nil {
		if err := s.AnalyticsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("AnalyticsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server to create a user pool client.
type CreateUserPoolClientOutput struct {
	_ struct{} `type:"structure"`

	// The user pool client that was just created.
	UserPoolClient *UserPoolClientType `type:"structure"`
}

// String returns the string representation
func (s CreateUserPoolClientOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateUserPoolClient = "CreateUserPoolClient"

// CreateUserPoolClientRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Creates the user pool client.
//
//    // Example sending a request using CreateUserPoolClientRequest.
//    req := client.CreateUserPoolClientRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/CreateUserPoolClient
func (c *Client) CreateUserPoolClientRequest(input *CreateUserPoolClientInput) CreateUserPoolClientRequest {
	op := &aws.Operation{
		Name:       opCreateUserPoolClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateUserPoolClientInput{}
	}

	req := c.newRequest(op, input, &CreateUserPoolClientOutput{})
	return CreateUserPoolClientRequest{Request: req, Input: input, Copy: c.CreateUserPoolClientRequest}
}

// CreateUserPoolClientRequest is the request type for the
// CreateUserPoolClient API operation.
type CreateUserPoolClientRequest struct {
	*aws.Request
	Input *CreateUserPoolClientInput
	Copy  func(*CreateUserPoolClientInput) CreateUserPoolClientRequest
}

// Send marshals and sends the CreateUserPoolClient API request.
func (r CreateUserPoolClientRequest) Send(ctx context.Context) (*CreateUserPoolClientResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserPoolClientResponse{
		CreateUserPoolClientOutput: r.Request.Data.(*CreateUserPoolClientOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserPoolClientResponse is the response type for the
// CreateUserPoolClient API operation.
type CreateUserPoolClientResponse struct {
	*CreateUserPoolClientOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUserPoolClient request.
func (r *CreateUserPoolClientResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
