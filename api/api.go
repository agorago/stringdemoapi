package api

import "context"

// UppercaseRequest - the payload for Uppercase service
type UppercaseRequest struct {
	S string `json:"s"`
}

// UppercaseResponse - the  Uppercase service response
type UppercaseResponse struct {
	V string `json:"v"`
}

// CountRequest - the payload for Count service
type CountRequest struct {
	S string `json:"s"`
}

// CountResponse - the  Count service response
type CountResponse struct {
	V int `json:"v"`
}

// AddNumbersResponse - the  AddNumbers service response
type AddNumbersResponse struct {
	Sum int `json:"sum"`
}

// StringDemoService - the interface that is going to be implemented by the string demo service
// This has methods to illustrate features of the WeGO framework
type StringDemoService interface {
	// Uppercase - Converts the input string into upper case
	Uppercase( // the context
		ctx context.Context,
		// The upper case request
		ucr *UppercaseRequest) (
		// the upper case response
		UppercaseResponse, error)
	// Count - returns the length of the input string
	Count(ctx context.Context, cr *CountRequest) (CountResponse, error)
	// AddNumbers - adds two numbers and returns the result
	// This method illustrates a GET method implementation in WeGO since there is no request payload required
	AddNumbers(ctx context.Context, arg1 int, arg2 int) (AddNumbersResponse, error)
}
