// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetImageInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image that you want to retrieve.
	//
	// ImageBuildVersionArn is a required field
	ImageBuildVersionArn *string `location:"querystring" locationName:"imageBuildVersionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetImageInput"}

	if s.ImageBuildVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageBuildVersionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ImageBuildVersionArn != nil {
		v := *s.ImageBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "imageBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetImageOutput struct {
	_ struct{} `type:"structure"`

	// The image object.
	Image *Image `locationName:"image" type:"structure"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s GetImageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Image != nil {
		v := s.Image

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "image", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetImage = "GetImage"

// GetImageRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Gets an image.
//
//    // Example sending a request using GetImageRequest.
//    req := client.GetImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/GetImage
func (c *Client) GetImageRequest(input *GetImageInput) GetImageRequest {
	op := &aws.Operation{
		Name:       opGetImage,
		HTTPMethod: "GET",
		HTTPPath:   "/GetImage",
	}

	if input == nil {
		input = &GetImageInput{}
	}

	req := c.newRequest(op, input, &GetImageOutput{})
	return GetImageRequest{Request: req, Input: input, Copy: c.GetImageRequest}
}

// GetImageRequest is the request type for the
// GetImage API operation.
type GetImageRequest struct {
	*aws.Request
	Input *GetImageInput
	Copy  func(*GetImageInput) GetImageRequest
}

// Send marshals and sends the GetImage API request.
func (r GetImageRequest) Send(ctx context.Context) (*GetImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetImageResponse{
		GetImageOutput: r.Request.Data.(*GetImageOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetImageResponse is the response type for the
// GetImage API operation.
type GetImageResponse struct {
	*GetImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetImage request.
func (r *GetImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
