package err

import (
	"context"
	wegoe "github.com/agorago/wego/err"
	"net/http"
)

// It is recommended that each module define its own error file

func internalMakeWegoError(ctx context.Context, ll wegoe.LogLevel, e WegoErrorCode, httpErrorCode int, args map[string]interface{}) wegoe.WeGOError {
	return wegoe.MakeErrWithHTTPCode(ctx, ll, int(e), e.String(), httpErrorCode, args)
}

// Error - returns a customized CAFUError for WeGO
func Error(ctx context.Context, e WegoErrorCode, args map[string]interface{}) wegoe.WeGOError {
	return internalMakeWegoError(ctx, wegoe.Error, e, http.StatusInternalServerError, args)

}

// Warning - returns a customized CAFUError for WeGO
func Warning(ctx context.Context, e WegoErrorCode, args map[string]interface{}) wegoe.WeGOError {
	return internalMakeWegoError(ctx, wegoe.Warning, e, http.StatusInternalServerError, args)

}

// HTTPError - returns a customized CAFUError for WeGO
func HTTPError(ctx context.Context, httpErrorCode int, e WegoErrorCode, args map[string]interface{}) wegoe.WeGOError {
	return internalMakeWegoError(ctx, wegoe.Error, e, httpErrorCode, args)

}

// HTTPWarning - returns a customized CAFUError for WeGO
func HTTPWarning(ctx context.Context, httpErrorCode int, e WegoErrorCode, args map[string]interface{}) wegoe.WeGOError {
	return internalMakeWegoError(ctx, wegoe.Warning, e, httpErrorCode, args)

}

// WegoErrorCode - A WeGO error code
type WegoErrorCode int

// enumeration for WeGO Error codes
const (
	CannotInvokeOperation         WegoErrorCode = iota + 200000 // stringdemoapi.errors.CannotInvokeOperation
	CannotCastResponse                                          // stringdemoapi.errors.CannotCastResponse
	UnexpectedProxyInputParameter                               // stringdemoapi.errors.UnexpectedProxyInputParameter
)

//go:generate stringer -linecomment -type=WegoErrorCode
