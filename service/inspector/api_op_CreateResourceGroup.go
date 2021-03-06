// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateResourceGroupInput struct {
	_ struct{} `type:"structure"`

	// A collection of keys and an array of possible values, '[{"key":"key1","values":["Value1","Value2"]},{"key":"Key2","values":["Value3"]}]'.
	//
	// For example,'[{"key":"Name","values":["TestEC2Instance"]}]'.
	//
	// ResourceGroupTags is a required field
	ResourceGroupTags []ResourceGroupTag `locationName:"resourceGroupTags" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateResourceGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResourceGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateResourceGroupInput"}

	if s.ResourceGroupTags == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupTags"))
	}
	if s.ResourceGroupTags != nil && len(s.ResourceGroupTags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupTags", 1))
	}
	if s.ResourceGroupTags != nil {
		for i, v := range s.ResourceGroupTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceGroupTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateResourceGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the resource group that is created.
	//
	// ResourceGroupArn is a required field
	ResourceGroupArn *string `locationName:"resourceGroupArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateResourceGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateResourceGroup = "CreateResourceGroup"

// CreateResourceGroupRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Creates a resource group using the specified set of tags (key and value pairs)
// that are used to select the EC2 instances to be included in an Amazon Inspector
// assessment target. The created resource group is then used to create an Amazon
// Inspector assessment target. For more information, see CreateAssessmentTarget.
//
//    // Example sending a request using CreateResourceGroupRequest.
//    req := client.CreateResourceGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/CreateResourceGroup
func (c *Client) CreateResourceGroupRequest(input *CreateResourceGroupInput) CreateResourceGroupRequest {
	op := &aws.Operation{
		Name:       opCreateResourceGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateResourceGroupInput{}
	}

	req := c.newRequest(op, input, &CreateResourceGroupOutput{})
	return CreateResourceGroupRequest{Request: req, Input: input, Copy: c.CreateResourceGroupRequest}
}

// CreateResourceGroupRequest is the request type for the
// CreateResourceGroup API operation.
type CreateResourceGroupRequest struct {
	*aws.Request
	Input *CreateResourceGroupInput
	Copy  func(*CreateResourceGroupInput) CreateResourceGroupRequest
}

// Send marshals and sends the CreateResourceGroup API request.
func (r CreateResourceGroupRequest) Send(ctx context.Context) (*CreateResourceGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateResourceGroupResponse{
		CreateResourceGroupOutput: r.Request.Data.(*CreateResourceGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateResourceGroupResponse is the response type for the
// CreateResourceGroup API operation.
type CreateResourceGroupResponse struct {
	*CreateResourceGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateResourceGroup request.
func (r *CreateResourceGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
