// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateSkillGroupFromRoomInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the room from which the skill group is to be disassociated. Required.
	RoomArn *string `type:"string"`

	// The ARN of the skill group to disassociate from a room. Required.
	SkillGroupArn *string `type:"string"`
}

// String returns the string representation
func (s DisassociateSkillGroupFromRoomInput) String() string {
	return awsutil.Prettify(s)
}

type DisassociateSkillGroupFromRoomOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateSkillGroupFromRoomOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateSkillGroupFromRoom = "DisassociateSkillGroupFromRoom"

// DisassociateSkillGroupFromRoomRequest returns a request value for making API operation for
// Alexa For Business.
//
// Disassociates a skill group from a specified room. This disables all skills
// in the skill group on all devices in the room.
//
//    // Example sending a request using DisassociateSkillGroupFromRoomRequest.
//    req := client.DisassociateSkillGroupFromRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DisassociateSkillGroupFromRoom
func (c *Client) DisassociateSkillGroupFromRoomRequest(input *DisassociateSkillGroupFromRoomInput) DisassociateSkillGroupFromRoomRequest {
	op := &aws.Operation{
		Name:       opDisassociateSkillGroupFromRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateSkillGroupFromRoomInput{}
	}

	req := c.newRequest(op, input, &DisassociateSkillGroupFromRoomOutput{})
	return DisassociateSkillGroupFromRoomRequest{Request: req, Input: input, Copy: c.DisassociateSkillGroupFromRoomRequest}
}

// DisassociateSkillGroupFromRoomRequest is the request type for the
// DisassociateSkillGroupFromRoom API operation.
type DisassociateSkillGroupFromRoomRequest struct {
	*aws.Request
	Input *DisassociateSkillGroupFromRoomInput
	Copy  func(*DisassociateSkillGroupFromRoomInput) DisassociateSkillGroupFromRoomRequest
}

// Send marshals and sends the DisassociateSkillGroupFromRoom API request.
func (r DisassociateSkillGroupFromRoomRequest) Send(ctx context.Context) (*DisassociateSkillGroupFromRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateSkillGroupFromRoomResponse{
		DisassociateSkillGroupFromRoomOutput: r.Request.Data.(*DisassociateSkillGroupFromRoomOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateSkillGroupFromRoomResponse is the response type for the
// DisassociateSkillGroupFromRoom API operation.
type DisassociateSkillGroupFromRoomResponse struct {
	*DisassociateSkillGroupFromRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateSkillGroupFromRoom request.
func (r *DisassociateSkillGroupFromRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
