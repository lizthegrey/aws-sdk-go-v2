// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateSimulationApplicationVersionInput struct {
	_ struct{} `type:"structure"`

	// The application information for the simulation application.
	//
	// Application is a required field
	Application *string `locationName:"application" min:"1" type:"string" required:"true"`

	// The current revision id for the simulation application. If you provide a
	// value and it matches the latest revision ID, a new version will be created.
	CurrentRevisionId *string `locationName:"currentRevisionId" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateSimulationApplicationVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSimulationApplicationVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSimulationApplicationVersionInput"}

	if s.Application == nil {
		invalidParams.Add(aws.NewErrParamRequired("Application"))
	}
	if s.Application != nil && len(*s.Application) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Application", 1))
	}
	if s.CurrentRevisionId != nil && len(*s.CurrentRevisionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentRevisionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSimulationApplicationVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Application != nil {
		v := *s.Application

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "application", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CurrentRevisionId != nil {
		v := *s.CurrentRevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currentRevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateSimulationApplicationVersionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the simulation application.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the simulation application
	// was last updated.
	LastUpdatedAt *time.Time `locationName:"lastUpdatedAt" type:"timestamp"`

	// The name of the simulation application.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The rendering engine for the simulation application.
	RenderingEngine *RenderingEngine `locationName:"renderingEngine" type:"structure"`

	// The revision ID of the simulation application.
	RevisionId *string `locationName:"revisionId" min:"1" type:"string"`

	// Information about the robot software suite (ROS distribution).
	RobotSoftwareSuite *RobotSoftwareSuite `locationName:"robotSoftwareSuite" type:"structure"`

	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *SimulationSoftwareSuite `locationName:"simulationSoftwareSuite" type:"structure"`

	// The sources of the simulation application.
	Sources []Source `locationName:"sources" type:"list"`

	// The version of the simulation application.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateSimulationApplicationVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSimulationApplicationVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RenderingEngine != nil {
		v := s.RenderingEngine

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "renderingEngine", v, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RobotSoftwareSuite != nil {
		v := s.RobotSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "robotSoftwareSuite", v, metadata)
	}
	if s.SimulationSoftwareSuite != nil {
		v := s.SimulationSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "simulationSoftwareSuite", v, metadata)
	}
	if s.Sources != nil {
		v := s.Sources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "sources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateSimulationApplicationVersion = "CreateSimulationApplicationVersion"

// CreateSimulationApplicationVersionRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Creates a simulation application with a specific revision id.
//
//    // Example sending a request using CreateSimulationApplicationVersionRequest.
//    req := client.CreateSimulationApplicationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateSimulationApplicationVersion
func (c *Client) CreateSimulationApplicationVersionRequest(input *CreateSimulationApplicationVersionInput) CreateSimulationApplicationVersionRequest {
	op := &aws.Operation{
		Name:       opCreateSimulationApplicationVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/createSimulationApplicationVersion",
	}

	if input == nil {
		input = &CreateSimulationApplicationVersionInput{}
	}

	req := c.newRequest(op, input, &CreateSimulationApplicationVersionOutput{})
	return CreateSimulationApplicationVersionRequest{Request: req, Input: input, Copy: c.CreateSimulationApplicationVersionRequest}
}

// CreateSimulationApplicationVersionRequest is the request type for the
// CreateSimulationApplicationVersion API operation.
type CreateSimulationApplicationVersionRequest struct {
	*aws.Request
	Input *CreateSimulationApplicationVersionInput
	Copy  func(*CreateSimulationApplicationVersionInput) CreateSimulationApplicationVersionRequest
}

// Send marshals and sends the CreateSimulationApplicationVersion API request.
func (r CreateSimulationApplicationVersionRequest) Send(ctx context.Context) (*CreateSimulationApplicationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSimulationApplicationVersionResponse{
		CreateSimulationApplicationVersionOutput: r.Request.Data.(*CreateSimulationApplicationVersionOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSimulationApplicationVersionResponse is the response type for the
// CreateSimulationApplicationVersion API operation.
type CreateSimulationApplicationVersionResponse struct {
	*CreateSimulationApplicationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSimulationApplicationVersion request.
func (r *CreateSimulationApplicationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
