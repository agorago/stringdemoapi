package proxy

import (
	"context"
	api "github.com/agorago/stringdemoapi/api"
	e "github.com/agorago/stringdemoapi/internal/err"
	wegohttp "github.com/agorago/wego/http"
)

type stringdemoProxy struct{
	WegoProxy wegohttp.ProxyService
}

func MakeStringdemoProxy(wegoproxy wegohttp.ProxyService) api.StringDemoService{
	// make sure that the proxy info is first initialized
	return stringdemoProxy{
		WegoProxy: wegoproxy,
	}
}
// Uppercase - proxies the Uppercase and calls the server using HTTP
func (proxy stringdemoProxy) Uppercase(ctx context.Context, ucr *api.UpperCaseRequest) (api.UpperCaseResponse, error) {
	resp, err := proxy.WegoProxy.ProxyRequest(ctx, "stringdemo", "Uppercase", ucr)
	if err != nil {
		return api.UpperCaseResponse{}, err
	}
	r, ok := resp.(*api.UpperCaseResponse)
	if ok {
		return *r, nil
	}

	return api.UpperCaseResponse{}, e.MakeBplusError(ctx, e.CannotCastResponse, map[string]interface{}{
		"Response": resp})

}

// Count - proxies the Count and calls the server using HTTP
func (proxy stringdemoProxy) Count(ctx context.Context, cr *api.CountRequest) (api.CountResponse, error) {
	resp, err :=proxy.WegoProxy.ProxyRequest(ctx, "stringdemo", "Count", cr)
	if err != nil {
		return api.CountResponse{}, err
	}
	r, ok := resp.(*api.CountResponse)
	if ok {
		return *r, nil
	}

	return api.CountResponse{}, e.MakeBplusError(ctx, e.CannotCastResponse, map[string]interface{}{
		"Response": resp})

}

// AddNumbers - proxies the AddNumbers and calls the server using HTTP
func (proxy stringdemoProxy) AddNumbers(ctx context.Context, arg1 int, arg2 int) (api.AddNumbersResponse, error) {
	resp, err := proxy.WegoProxy.ProxyRequest(ctx, "stringdemo", "AddNumbers", arg1, arg2)
	if err != nil {
		return api.AddNumbersResponse{}, err
	}
	r, ok := resp.(*api.AddNumbersResponse)
	if ok {
		return *r, nil
	}

	return api.AddNumbersResponse{}, e.MakeBplusError(ctx, e.CannotCastResponse, map[string]interface{}{
		"Response": resp})

}
