package proxy
	
	import (
		"context"
	
		bplus "gitlab.intelligentb.com/devops/bplus/http"
		api "gitlab.intelligentb.com/examples/stringdemo/stringdemoapi/api"
		e "gitlab.intelligentb.com/examples/stringdemo/stringdemoapi/internal/err"
	)

	type Stringdemo struct {}


// Uppercase - proxies the Uppercase and calls the server using HTTP
func (Stringdemo) Uppercase(ctx context.Context,ucr *api.UpperCaseRequest)(api.UpperCaseResponse,error){
	resp, err := bplus.ProxyRequest(ctx, "stringdemo", "Uppercase" ,ucr)
	if err != nil {
		return api.UpperCaseResponse{}, err
	}
	r, ok := resp.(*api.UpperCaseResponse)
	if ok {
		return *r,nil
	}

	return api.UpperCaseResponse{}, e.MakeBplusError(ctx, e.CannotCastResponse,map[string]interface{}{
		"Response":resp, })

}

// Count - proxies the Count and calls the server using HTTP
func (Stringdemo) Count(ctx context.Context,cr *api.CountRequest)(api.CountResponse,error){
	resp, err := bplus.ProxyRequest(ctx, "stringdemo", "Count" ,cr)
	if err != nil {
		return api.CountResponse{}, err
	}
	r, ok := resp.(*api.CountResponse)
	if ok {
		return *r,nil
	}

	return api.CountResponse{}, e.MakeBplusError(ctx, e.CannotCastResponse,map[string]interface{}{
		"Response":resp, })

}

// AddNumbers - proxies the AddNumbers and calls the server using HTTP
func (Stringdemo) AddNumbers(ctx context.Context,arg1 int,arg2 int)(api.AddNumbersResponse,error){
	resp, err := bplus.ProxyRequest(ctx, "stringdemo", "AddNumbers" ,arg1,arg2)
	if err != nil {
		return api.AddNumbersResponse{}, err
	}
	r, ok := resp.(*api.AddNumbersResponse)
	if ok {
		return *r,nil
	}

	return api.AddNumbersResponse{}, e.MakeBplusError(ctx, e.CannotCastResponse,map[string]interface{}{
		"Response":resp, })

}

