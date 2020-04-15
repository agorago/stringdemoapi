package proxy

import (
	"context"
	"fmt"
	bplusc "gitlab.intelligentb.com/devops/bplus/context"
	"gitlab.intelligentb.com/devops/bplus/fw"
	"gitlab.intelligentb.com/examples/stringdemo/stringdemoapi/api"
	e "gitlab.intelligentb.com/examples/stringdemo/stringdemoapi/internal/err"
)

func InterceptCount(ctx context.Context,chain *fw.MiddlewareChain)context.Context{
	// Intercept at the proxy side when the argument is "Count".  Return 8 instead of 5 for count
	cr,ok :=  bplusc.GetPayload(ctx).(*api.CountRequest)
	if !ok {
		ctx = bplusc.SetError(ctx,e.MakeBplusError(ctx,e.UnexpectedProxyInputParameter,nil))
		return ctx
	}
	if cr.S == "Count"{
		// this snippet of code will completely bypass the remote http call.
		// can be great to implement circuit breakers with default values
		fmt.Printf("count is encountered\n")
		ctx = bplusc.SetResponsePayload(ctx,&api.CountResponse{V: 8,})
		return ctx
	}
	ctx = chain.DoContinue(ctx)
	return ctx
}
