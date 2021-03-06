// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartContentModerationInput struct {
	_ struct{} `type:"structure"`

	// Idempotent token used to identify the start request. If you use the same
	// token with multiple StartContentModeration requests, the same JobId is returned.
	// Use ClientRequestToken to prevent the same job from being accidently started
	// more than once.
	ClientRequestToken *string `min:"1" type:"string"`

	// An identifier you specify that's returned in the completion notification
	// that's published to your Amazon Simple Notification Service topic. For example,
	// you can use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string `min:"1" type:"string"`

	// Specifies the minimum confidence that Amazon Rekognition must have in order
	// to return a moderated content label. Confidence represents how certain Amazon
	// Rekognition is that the moderated content is correctly identified. 0 is the
	// lowest confidence. 100 is the highest confidence. Amazon Rekognition doesn't
	// return any moderated content labels with a confidence level lower than this
	// specified value. If you don't specify MinConfidence, GetContentModeration
	// returns labels with confidence values greater than or equal to 50 percent.
	MinConfidence *float64 `type:"float"`

	// The Amazon SNS topic ARN that you want Amazon Rekognition Video to publish
	// the completion status of the unsafe content analysis to.
	NotificationChannel *NotificationChannel `type:"structure"`

	// The video in which you want to detect unsafe content. The video must be stored
	// in an Amazon S3 bucket.
	//
	// Video is a required field
	Video *Video `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartContentModerationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartContentModerationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartContentModerationInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.JobTag != nil && len(*s.JobTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobTag", 1))
	}

	if s.Video == nil {
		invalidParams.Add(aws.NewErrParamRequired("Video"))
	}
	if s.NotificationChannel != nil {
		if err := s.NotificationChannel.Validate(); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(aws.ErrInvalidParams))
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			invalidParams.AddNested("Video", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartContentModerationOutput struct {
	_ struct{} `type:"structure"`

	// The identifier for the unsafe content analysis job. Use JobId to identify
	// the job in a subsequent call to GetContentModeration.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartContentModerationOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartContentModeration = "StartContentModeration"

// StartContentModerationRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Starts asynchronous detection of unsafe content in a stored video.
//
// Amazon Rekognition Video can moderate content in a video stored in an Amazon
// S3 bucket. Use Video to specify the bucket name and the filename of the video.
// StartContentModeration returns a job identifier (JobId) which you use to
// get the results of the analysis. When unsafe content analysis is finished,
// Amazon Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic that you specify in NotificationChannel.
//
// To get the results of the unsafe content analysis, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED. If so, call GetContentModeration
// and pass the job identifier (JobId) from the initial call to StartContentModeration.
//
// For more information, see Detecting Unsafe Content in the Amazon Rekognition
// Developer Guide.
//
//    // Example sending a request using StartContentModerationRequest.
//    req := client.StartContentModerationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartContentModerationRequest(input *StartContentModerationInput) StartContentModerationRequest {
	op := &aws.Operation{
		Name:       opStartContentModeration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartContentModerationInput{}
	}

	req := c.newRequest(op, input, &StartContentModerationOutput{})
	return StartContentModerationRequest{Request: req, Input: input, Copy: c.StartContentModerationRequest}
}

// StartContentModerationRequest is the request type for the
// StartContentModeration API operation.
type StartContentModerationRequest struct {
	*aws.Request
	Input *StartContentModerationInput
	Copy  func(*StartContentModerationInput) StartContentModerationRequest
}

// Send marshals and sends the StartContentModeration API request.
func (r StartContentModerationRequest) Send(ctx context.Context) (*StartContentModerationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartContentModerationResponse{
		StartContentModerationOutput: r.Request.Data.(*StartContentModerationOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartContentModerationResponse is the response type for the
// StartContentModeration API operation.
type StartContentModerationResponse struct {
	*StartContentModerationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartContentModeration request.
func (r *StartContentModerationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
