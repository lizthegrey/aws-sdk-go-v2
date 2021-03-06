// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEngineDefaultClusterParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster parameter group family to return engine parameter
	// information for.
	//
	// DBParameterGroupFamily is a required field
	DBParameterGroupFamily *string `type:"string" required:"true"`

	// This parameter isn't currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeEngineDefaultClusterParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeEngineDefaultClusterParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEngineDefaultClusterParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEngineDefaultClusterParametersInput"}

	if s.DBParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupFamily"))
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

type DescribeEngineDefaultClusterParametersOutput struct {
	_ struct{} `type:"structure"`

	// Contains the result of a successful invocation of the DescribeEngineDefaultParameters
	// action.
	EngineDefaults *EngineDefaults `type:"structure"`
}

// String returns the string representation
func (s DescribeEngineDefaultClusterParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEngineDefaultClusterParameters = "DescribeEngineDefaultClusterParameters"

// DescribeEngineDefaultClusterParametersRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns the default engine and system parameter information for the cluster
// database engine.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
//    // Example sending a request using DescribeEngineDefaultClusterParametersRequest.
//    req := client.DescribeEngineDefaultClusterParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeEngineDefaultClusterParameters
func (c *Client) DescribeEngineDefaultClusterParametersRequest(input *DescribeEngineDefaultClusterParametersInput) DescribeEngineDefaultClusterParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeEngineDefaultClusterParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEngineDefaultClusterParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeEngineDefaultClusterParametersOutput{})
	return DescribeEngineDefaultClusterParametersRequest{Request: req, Input: input, Copy: c.DescribeEngineDefaultClusterParametersRequest}
}

// DescribeEngineDefaultClusterParametersRequest is the request type for the
// DescribeEngineDefaultClusterParameters API operation.
type DescribeEngineDefaultClusterParametersRequest struct {
	*aws.Request
	Input *DescribeEngineDefaultClusterParametersInput
	Copy  func(*DescribeEngineDefaultClusterParametersInput) DescribeEngineDefaultClusterParametersRequest
}

// Send marshals and sends the DescribeEngineDefaultClusterParameters API request.
func (r DescribeEngineDefaultClusterParametersRequest) Send(ctx context.Context) (*DescribeEngineDefaultClusterParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEngineDefaultClusterParametersResponse{
		DescribeEngineDefaultClusterParametersOutput: r.Request.Data.(*DescribeEngineDefaultClusterParametersOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEngineDefaultClusterParametersResponse is the response type for the
// DescribeEngineDefaultClusterParameters API operation.
type DescribeEngineDefaultClusterParametersResponse struct {
	*DescribeEngineDefaultClusterParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEngineDefaultClusterParameters request.
func (r *DescribeEngineDefaultClusterParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
