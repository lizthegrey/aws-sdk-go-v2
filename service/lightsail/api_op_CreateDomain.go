// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDomainInput struct {
	_ struct{} `type:"structure"`

	// The domain name to manage (e.g., example.com).
	//
	// You cannot register a new domain name using Lightsail. You must register
	// a domain name using Amazon Route 53 or another domain name registrar. If
	// you have already registered your domain, you can enter its name in this parameter
	// to manage the DNS records for that domain.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDomainOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the time stamp of the request, and the resources affected
	// by the request.
	Operation *Operation `locationName:"operation" type:"structure"`
}

// String returns the string representation
func (s CreateDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDomain = "CreateDomain"

// CreateDomainRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates a domain resource for the specified domain (e.g., example.com).
//
// The create domain operation supports tag-based access control via request
// tags. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CreateDomainRequest.
//    req := client.CreateDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateDomain
func (c *Client) CreateDomainRequest(input *CreateDomainInput) CreateDomainRequest {
	op := &aws.Operation{
		Name:       opCreateDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDomainInput{}
	}

	req := c.newRequest(op, input, &CreateDomainOutput{})
	return CreateDomainRequest{Request: req, Input: input, Copy: c.CreateDomainRequest}
}

// CreateDomainRequest is the request type for the
// CreateDomain API operation.
type CreateDomainRequest struct {
	*aws.Request
	Input *CreateDomainInput
	Copy  func(*CreateDomainInput) CreateDomainRequest
}

// Send marshals and sends the CreateDomain API request.
func (r CreateDomainRequest) Send(ctx context.Context) (*CreateDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDomainResponse{
		CreateDomainOutput: r.Request.Data.(*CreateDomainOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDomainResponse is the response type for the
// CreateDomain API operation.
type CreateDomainResponse struct {
	*CreateDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDomain request.
func (r *CreateDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
