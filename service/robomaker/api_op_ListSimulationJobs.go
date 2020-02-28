// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListSimulationJobsInput struct {
	_ struct{} `type:"structure"`

	// Optional filters to limit results.
	//
	// The filter names status and simulationApplicationName and robotApplicationName
	// are supported. When filtering, you must use the complete value of the filtered
	// item. You can use up to three filters, but they must be for the same named
	// item. For example, if you are looking for items with the status Preparing
	// or the status Running.
	Filters []Filter `locationName:"filters" min:"1" type:"list"`

	// When this parameter is used, ListSimulationJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListSimulationJobs
	// request with the returned nextToken value. This value can be between 1 and
	// 1000. If this parameter is not used, then ListSimulationJobs returns up to
	// 1000 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListSimulationJobs
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListSimulationJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSimulationJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSimulationJobsInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSimulationJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListSimulationJobsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListSimulationJobs request. When
	// the results of a ListRobot request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// A list of simulation job summaries that meet the criteria of the request.
	//
	// SimulationJobSummaries is a required field
	SimulationJobSummaries []SimulationJobSummary `locationName:"simulationJobSummaries" type:"list" required:"true"`
}

// String returns the string representation
func (s ListSimulationJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSimulationJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SimulationJobSummaries != nil {
		v := s.SimulationJobSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "simulationJobSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListSimulationJobs = "ListSimulationJobs"

// ListSimulationJobsRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Returns a list of simulation jobs. You can optionally provide filters to
// retrieve specific simulation jobs.
//
//    // Example sending a request using ListSimulationJobsRequest.
//    req := client.ListSimulationJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListSimulationJobs
func (c *Client) ListSimulationJobsRequest(input *ListSimulationJobsInput) ListSimulationJobsRequest {
	op := &aws.Operation{
		Name:       opListSimulationJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/listSimulationJobs",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSimulationJobsInput{}
	}

	req := c.newRequest(op, input, &ListSimulationJobsOutput{})
	return ListSimulationJobsRequest{Request: req, Input: input, Copy: c.ListSimulationJobsRequest}
}

// ListSimulationJobsRequest is the request type for the
// ListSimulationJobs API operation.
type ListSimulationJobsRequest struct {
	*aws.Request
	Input *ListSimulationJobsInput
	Copy  func(*ListSimulationJobsInput) ListSimulationJobsRequest
}

// Send marshals and sends the ListSimulationJobs API request.
func (r ListSimulationJobsRequest) Send(ctx context.Context) (*ListSimulationJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSimulationJobsResponse{
		ListSimulationJobsOutput: r.Request.Data.(*ListSimulationJobsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSimulationJobsRequestPaginator returns a paginator for ListSimulationJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSimulationJobsRequest(input)
//   p := robomaker.NewListSimulationJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSimulationJobsPaginator(req ListSimulationJobsRequest) ListSimulationJobsPaginator {
	return ListSimulationJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSimulationJobsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListSimulationJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSimulationJobsPaginator struct {
	aws.Pager
}

func (p *ListSimulationJobsPaginator) CurrentPage() *ListSimulationJobsOutput {
	return p.Pager.CurrentPage().(*ListSimulationJobsOutput)
}

// ListSimulationJobsResponse is the response type for the
// ListSimulationJobs API operation.
type ListSimulationJobsResponse struct {
	*ListSimulationJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSimulationJobs request.
func (r *ListSimulationJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
