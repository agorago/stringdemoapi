package proxy

import (
	"context"
	"fmt"
	"github.com/agorago/stringdemoapi/api"
	e "github.com/agorago/stringdemoapi/internal/err"
	wegocontext "github.com/agorago/wego/context"
	"github.com/agorago/wego/fw"
)

func InterceptCount(ctx context.Context, chain *fw.MiddlewareChain) context.Context {
	// Intercept at the proxy side when the argument is "Count".  Return 8 instead of 5 for count
	cr, ok := wegocontext.GetPayload(ctx).(*api.CountRequest)
	if !ok {
		ctx = wegocontext.SetError(ctx, e.Error(ctx, e.UnexpectedProxyInputParameter, nil))
		return ctx
	}
	if cr.S == "Count" {
		// this snippet of code will completely bypass the remote http call.
		// can be great to implement circuit breakers with default values
		fmt.Printf("count is encountered\n")
		ctx = wegocontext.SetResponsePayload(ctx, &api.CountResponse{V: 8})
		return ctx
	}
	ctx = chain.DoContinue(ctx)
	return ctx
}
