// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLifecyclePolicyPreviewInput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that filters results based on image tag status and
	// all tags, if tagged.
	Filter *LifecyclePolicyPreviewFilter `locationName:"filter" type:"structure"`

	// The list of imageIDs to be included.
	ImageIds []ImageIdentifier `locationName:"imageIds" min:"1" type:"list"`

	// The maximum number of repository results returned by GetLifecyclePolicyPreviewRequest
	// in paginated output. When this parameter is used, GetLifecyclePolicyPreviewRequest
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another GetLifecyclePolicyPreviewRequest request with the returned nextToken
	// value. This value can be between 1 and 1000. If this parameter is not used,
	// then GetLifecyclePolicyPreviewRequest returns up to 100 results and a nextToken
	// value, if applicable. This option cannot be used when you specify images
	// with imageIds.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated GetLifecyclePolicyPreviewRequest
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return. This option cannot be used when you specify images with imageIds.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The AWS account ID associated with the registry that contains the repository.
	// If you do not specify a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLifecyclePolicyPreviewInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLifecyclePolicyPreviewInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLifecyclePolicyPreviewInput"}
	if s.ImageIds != nil && len(s.ImageIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageIds", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}
	if s.ImageIds != nil {
		for i, v := range s.ImageIds {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ImageIds", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLifecyclePolicyPreviewOutput struct {
	_ struct{} `type:"structure"`

	// The JSON lifecycle policy text.
	LifecyclePolicyText *string `locationName:"lifecyclePolicyText" min:"100" type:"string"`

	// The nextToken value to include in a future GetLifecyclePolicyPreview request.
	// When the results of a GetLifecyclePolicyPreview request exceed maxResults,
	// this value can be used to retrieve the next page of results. This value is
	// null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The results of the lifecycle policy preview request.
	PreviewResults []LifecyclePolicyPreviewResult `locationName:"previewResults" type:"list"`

	// The registry ID associated with the request.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository name associated with the request.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`

	// The status of the lifecycle policy preview request.
	Status LifecyclePolicyPreviewStatus `locationName:"status" type:"string" enum:"true"`

	// The list of images that is returned as a result of the action.
	Summary *LifecyclePolicyPreviewSummary `locationName:"summary" type:"structure"`
}

// String returns the string representation
func (s GetLifecyclePolicyPreviewOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLifecyclePolicyPreview = "GetLifecyclePolicyPreview"

// GetLifecyclePolicyPreviewRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Retrieves the results of the lifecycle policy preview request for the specified
// repository.
//
//    // Example sending a request using GetLifecyclePolicyPreviewRequest.
//    req := client.GetLifecyclePolicyPreviewRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/GetLifecyclePolicyPreview
func (c *Client) GetLifecyclePolicyPreviewRequest(input *GetLifecyclePolicyPreviewInput) GetLifecyclePolicyPreviewRequest {
	op := &aws.Operation{
		Name:       opGetLifecyclePolicyPreview,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetLifecyclePolicyPreviewInput{}
	}

	req := c.newRequest(op, input, &GetLifecyclePolicyPreviewOutput{})
	return GetLifecyclePolicyPreviewRequest{Request: req, Input: input, Copy: c.GetLifecyclePolicyPreviewRequest}
}

// GetLifecyclePolicyPreviewRequest is the request type for the
// GetLifecyclePolicyPreview API operation.
type GetLifecyclePolicyPreviewRequest struct {
	*aws.Request
	Input *GetLifecyclePolicyPreviewInput
	Copy  func(*GetLifecyclePolicyPreviewInput) GetLifecyclePolicyPreviewRequest
}

// Send marshals and sends the GetLifecyclePolicyPreview API request.
func (r GetLifecyclePolicyPreviewRequest) Send(ctx context.Context) (*GetLifecyclePolicyPreviewResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLifecyclePolicyPreviewResponse{
		GetLifecyclePolicyPreviewOutput: r.Request.Data.(*GetLifecyclePolicyPreviewOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetLifecyclePolicyPreviewRequestPaginator returns a paginator for GetLifecyclePolicyPreview.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetLifecyclePolicyPreviewRequest(input)
//   p := ecr.NewGetLifecyclePolicyPreviewRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetLifecyclePolicyPreviewPaginator(req GetLifecyclePolicyPreviewRequest) GetLifecyclePolicyPreviewPaginator {
	return GetLifecyclePolicyPreviewPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetLifecyclePolicyPreviewInput
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

// GetLifecyclePolicyPreviewPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetLifecyclePolicyPreviewPaginator struct {
	aws.Pager
}

func (p *GetLifecyclePolicyPreviewPaginator) CurrentPage() *GetLifecyclePolicyPreviewOutput {
	return p.Pager.CurrentPage().(*GetLifecyclePolicyPreviewOutput)
}

// GetLifecyclePolicyPreviewResponse is the response type for the
// GetLifecyclePolicyPreview API operation.
type GetLifecyclePolicyPreviewResponse struct {
	*GetLifecyclePolicyPreviewOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLifecyclePolicyPreview request.
func (r *GetLifecyclePolicyPreviewResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
