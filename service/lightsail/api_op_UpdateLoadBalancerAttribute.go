// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateLoadBalancerAttributeInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute you want to update. Valid values are below.
	//
	// AttributeName is a required field
	AttributeName LoadBalancerAttributeName `locationName:"attributeName" type:"string" required:"true" enum:"true"`

	// The value that you want to specify for the attribute name.
	//
	// AttributeValue is a required field
	AttributeValue *string `locationName:"attributeValue" min:"1" type:"string" required:"true"`

	// The name of the load balancer that you want to modify (e.g., my-load-balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateLoadBalancerAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLoadBalancerAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLoadBalancerAttributeInput"}
	if len(s.AttributeName) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AttributeName"))
	}

	if s.AttributeValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeValue"))
	}
	if s.AttributeValue != nil && len(*s.AttributeValue) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AttributeValue", 1))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateLoadBalancerAttributeOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the time stamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s UpdateLoadBalancerAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateLoadBalancerAttribute = "UpdateLoadBalancerAttribute"

// UpdateLoadBalancerAttributeRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Updates the specified attribute for a load balancer. You can only update
// one attribute at a time.
//
// The update load balancer attribute operation supports tag-based access control
// via resource tags applied to the resource identified by load balancer name.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using UpdateLoadBalancerAttributeRequest.
//    req := client.UpdateLoadBalancerAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/UpdateLoadBalancerAttribute
func (c *Client) UpdateLoadBalancerAttributeRequest(input *UpdateLoadBalancerAttributeInput) UpdateLoadBalancerAttributeRequest {
	op := &aws.Operation{
		Name:       opUpdateLoadBalancerAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateLoadBalancerAttributeInput{}
	}

	req := c.newRequest(op, input, &UpdateLoadBalancerAttributeOutput{})
	return UpdateLoadBalancerAttributeRequest{Request: req, Input: input, Copy: c.UpdateLoadBalancerAttributeRequest}
}

// UpdateLoadBalancerAttributeRequest is the request type for the
// UpdateLoadBalancerAttribute API operation.
type UpdateLoadBalancerAttributeRequest struct {
	*aws.Request
	Input *UpdateLoadBalancerAttributeInput
	Copy  func(*UpdateLoadBalancerAttributeInput) UpdateLoadBalancerAttributeRequest
}

// Send marshals and sends the UpdateLoadBalancerAttribute API request.
func (r UpdateLoadBalancerAttributeRequest) Send(ctx context.Context) (*UpdateLoadBalancerAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLoadBalancerAttributeResponse{
		UpdateLoadBalancerAttributeOutput: r.Request.Data.(*UpdateLoadBalancerAttributeOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLoadBalancerAttributeResponse is the response type for the
// UpdateLoadBalancerAttribute API operation.
type UpdateLoadBalancerAttributeResponse struct {
	*UpdateLoadBalancerAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLoadBalancerAttribute request.
func (r *UpdateLoadBalancerAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
