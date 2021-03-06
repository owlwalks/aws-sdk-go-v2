// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListVocabularyFiltersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of filters to return in the response. If there are fewer
	// results in the list, this response contains only the actual results.
	MaxResults *int64 `min:"1" type:"integer"`

	// Filters the response so that it only contains vocabulary filters whose name
	// contains the specified string.
	NameContains *string `min:"1" type:"string"`

	// If the result of the previous request to ListVocabularyFilters was truncated,
	// include the NextToken to fetch the next set of collections.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListVocabularyFiltersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVocabularyFiltersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVocabularyFiltersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListVocabularyFiltersOutput struct {
	_ struct{} `type:"structure"`

	// The ListVocabularyFilters operation returns a page of collections at a time.
	// The maximum size of the page is set by the MaxResults parameter. If there
	// are more jobs in the list than the page size, Amazon Transcribe returns the
	// NextPage token. Include the token in the next request to the ListVocabularyFilters
	// operation to return in the next page of jobs.
	NextToken *string `type:"string"`

	// The list of vocabulary filters. It will contain at most MaxResults number
	// of filters. If there are more filters, call the ListVocabularyFilters operation
	// again with the NextToken parameter in the request set to the value of the
	// NextToken field in the response.
	VocabularyFilters []VocabularyFilterInfo `type:"list"`
}

// String returns the string representation
func (s ListVocabularyFiltersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListVocabularyFilters = "ListVocabularyFilters"

// ListVocabularyFiltersRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Gets information about vocabulary filters.
//
//    // Example sending a request using ListVocabularyFiltersRequest.
//    req := client.ListVocabularyFiltersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/ListVocabularyFilters
func (c *Client) ListVocabularyFiltersRequest(input *ListVocabularyFiltersInput) ListVocabularyFiltersRequest {
	op := &aws.Operation{
		Name:       opListVocabularyFilters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListVocabularyFiltersInput{}
	}

	req := c.newRequest(op, input, &ListVocabularyFiltersOutput{})
	return ListVocabularyFiltersRequest{Request: req, Input: input, Copy: c.ListVocabularyFiltersRequest}
}

// ListVocabularyFiltersRequest is the request type for the
// ListVocabularyFilters API operation.
type ListVocabularyFiltersRequest struct {
	*aws.Request
	Input *ListVocabularyFiltersInput
	Copy  func(*ListVocabularyFiltersInput) ListVocabularyFiltersRequest
}

// Send marshals and sends the ListVocabularyFilters API request.
func (r ListVocabularyFiltersRequest) Send(ctx context.Context) (*ListVocabularyFiltersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVocabularyFiltersResponse{
		ListVocabularyFiltersOutput: r.Request.Data.(*ListVocabularyFiltersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVocabularyFiltersRequestPaginator returns a paginator for ListVocabularyFilters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVocabularyFiltersRequest(input)
//   p := transcribe.NewListVocabularyFiltersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVocabularyFiltersPaginator(req ListVocabularyFiltersRequest) ListVocabularyFiltersPaginator {
	return ListVocabularyFiltersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListVocabularyFiltersInput
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

// ListVocabularyFiltersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVocabularyFiltersPaginator struct {
	aws.Pager
}

func (p *ListVocabularyFiltersPaginator) CurrentPage() *ListVocabularyFiltersOutput {
	return p.Pager.CurrentPage().(*ListVocabularyFiltersOutput)
}

// ListVocabularyFiltersResponse is the response type for the
// ListVocabularyFilters API operation.
type ListVocabularyFiltersResponse struct {
	*ListVocabularyFiltersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVocabularyFilters request.
func (r *ListVocabularyFiltersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
