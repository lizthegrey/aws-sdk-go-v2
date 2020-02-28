// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutLifecyclePolicyInput struct {
	_ struct{} `type:"structure"`

	// The JSON repository policy text to apply to the repository.
	//
	// LifecyclePolicyText is a required field
	LifecyclePolicyText *string `locationName:"lifecyclePolicyText" min:"100" type:"string" required:"true"`

	// The AWS account ID associated with the registry that contains the repository.
	// If you do not specify a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository to receive the policy.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PutLifecyclePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLifecyclePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLifecyclePolicyInput"}

	if s.LifecyclePolicyText == nil {
		invalidParams.Add(aws.NewErrParamRequired("LifecyclePolicyText"))
	}
	if s.LifecyclePolicyText != nil && len(*s.LifecyclePolicyText) < 100 {
		invalidParams.Add(aws.NewErrParamMinLen("LifecyclePolicyText", 100))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutLifecyclePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The JSON repository policy text.
	LifecyclePolicyText *string `locationName:"lifecyclePolicyText" min:"100" type:"string"`

	// The registry ID associated with the request.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository name associated with the request.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`
}

// String returns the string representation
func (s PutLifecyclePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutLifecyclePolicy = "PutLifecyclePolicy"

// PutLifecyclePolicyRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Creates or updates the lifecycle policy for the specified repository. For
// more information, see Lifecycle Policy Template (https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).
//
//    // Example sending a request using PutLifecyclePolicyRequest.
//    req := client.PutLifecyclePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/PutLifecyclePolicy
func (c *Client) PutLifecyclePolicyRequest(input *PutLifecyclePolicyInput) PutLifecyclePolicyRequest {
	op := &aws.Operation{
		Name:       opPutLifecyclePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutLifecyclePolicyInput{}
	}

	req := c.newRequest(op, input, &PutLifecyclePolicyOutput{})
	return PutLifecyclePolicyRequest{Request: req, Input: input, Copy: c.PutLifecyclePolicyRequest}
}

// PutLifecyclePolicyRequest is the request type for the
// PutLifecyclePolicy API operation.
type PutLifecyclePolicyRequest struct {
	*aws.Request
	Input *PutLifecyclePolicyInput
	Copy  func(*PutLifecyclePolicyInput) PutLifecyclePolicyRequest
}

// Send marshals and sends the PutLifecyclePolicy API request.
func (r PutLifecyclePolicyRequest) Send(ctx context.Context) (*PutLifecyclePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLifecyclePolicyResponse{
		PutLifecyclePolicyOutput: r.Request.Data.(*PutLifecyclePolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLifecyclePolicyResponse is the response type for the
// PutLifecyclePolicy API operation.
type PutLifecyclePolicyResponse struct {
	*PutLifecyclePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLifecyclePolicy request.
func (r *PutLifecyclePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
