// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstanceSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetInstanceSnapshots request. If
	// your results are paginated, the response will return a next page token that
	// you can specify as the page token in a subsequent request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetInstanceSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

type GetInstanceSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the results of your
	// get instance snapshots request.
	InstanceSnapshots []InstanceSnapshot `locationName:"instanceSnapshots" type:"list"`

	// The token to advance to the next page of resutls from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetInstanceSnapshots request
	// and specify the next page token using the pageToken parameter.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetInstanceSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstanceSnapshots = "GetInstanceSnapshots"

// GetInstanceSnapshotsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns all instance snapshots for the user's account.
//
//    // Example sending a request using GetInstanceSnapshotsRequest.
//    req := client.GetInstanceSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstanceSnapshots
func (c *Client) GetInstanceSnapshotsRequest(input *GetInstanceSnapshotsInput) GetInstanceSnapshotsRequest {
	op := &aws.Operation{
		Name:       opGetInstanceSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInstanceSnapshotsInput{}
	}

	req := c.newRequest(op, input, &GetInstanceSnapshotsOutput{})
	return GetInstanceSnapshotsRequest{Request: req, Input: input, Copy: c.GetInstanceSnapshotsRequest}
}

// GetInstanceSnapshotsRequest is the request type for the
// GetInstanceSnapshots API operation.
type GetInstanceSnapshotsRequest struct {
	*aws.Request
	Input *GetInstanceSnapshotsInput
	Copy  func(*GetInstanceSnapshotsInput) GetInstanceSnapshotsRequest
}

// Send marshals and sends the GetInstanceSnapshots API request.
func (r GetInstanceSnapshotsRequest) Send(ctx context.Context) (*GetInstanceSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstanceSnapshotsResponse{
		GetInstanceSnapshotsOutput: r.Request.Data.(*GetInstanceSnapshotsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstanceSnapshotsResponse is the response type for the
// GetInstanceSnapshots API operation.
type GetInstanceSnapshotsResponse struct {
	*GetInstanceSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstanceSnapshots request.
func (r *GetInstanceSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
