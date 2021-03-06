// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteClusterInput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster to be deleted.
	//
	// ClusterName is a required field
	ClusterName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteClusterInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteClusterOutput struct {
	_ struct{} `type:"structure"`

	// A description of the DAX cluster that is being deleted.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s DeleteClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCluster = "DeleteCluster"

// DeleteClusterRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Deletes a previously provisioned DAX cluster. DeleteCluster deletes all associated
// nodes, node endpoints and the DAX cluster itself. When you receive a successful
// response from this action, DAX immediately begins deleting the cluster; you
// cannot cancel or revert this action.
//
//    // Example sending a request using DeleteClusterRequest.
//    req := client.DeleteClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DeleteCluster
func (c *Client) DeleteClusterRequest(input *DeleteClusterInput) DeleteClusterRequest {
	op := &aws.Operation{
		Name:       opDeleteCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteClusterInput{}
	}

	req := c.newRequest(op, input, &DeleteClusterOutput{})
	return DeleteClusterRequest{Request: req, Input: input, Copy: c.DeleteClusterRequest}
}

// DeleteClusterRequest is the request type for the
// DeleteCluster API operation.
type DeleteClusterRequest struct {
	*aws.Request
	Input *DeleteClusterInput
	Copy  func(*DeleteClusterInput) DeleteClusterRequest
}

// Send marshals and sends the DeleteCluster API request.
func (r DeleteClusterRequest) Send(ctx context.Context) (*DeleteClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteClusterResponse{
		DeleteClusterOutput: r.Request.Data.(*DeleteClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteClusterResponse is the response type for the
// DeleteCluster API operation.
type DeleteClusterResponse struct {
	*DeleteClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCluster request.
func (r *DeleteClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
