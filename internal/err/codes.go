package err

import (
	"context"
	wegoe "github.com/agorago/wego/err"
	"net/http"
)

// It is recommended that each module define its own error file

func internalMakeBplusError(ctx context.Context, ll wegoe.LogLevel, e BPlusErrorCode, httpErrorCode int, args map[string]interface{}) wegoe.WeGOError {
	return wegoe.MakeErrWithHTTPCode(ctx, ll, int(e), e.String(), httpErrorCode, args)
}

// MakeBplusError - returns a customized CAFUError for BPlus
func MakeBplusError(ctx context.Context, e BPlusErrorCode, args map[string]interface{}) wegoe.WeGOError {
	return internalMakeBplusError(ctx, wegoe.Error, e, http.StatusInternalServerError, args)

}

// MakeBplusWarning - returns a customized CAFUError for BPlus
func MakeBplusWarning(ctx context.Context, e BPlusErrorCode, args map[string]interface{}) wegoe.WeGOError {
	return internalMakeBplusError(ctx, wegoe.Warning, e, http.StatusInternalServerError, args)

}

// MakeBplusErrorWithErrorCode - returns a customized CAFUError for BPlus
func MakeBplusErrorWithErrorCode(ctx context.Context, httpErrorCode int, e BPlusErrorCode, args map[string]interface{}) wegoe.WeGOError {
	return internalMakeBplusError(ctx, wegoe.Error, e, httpErrorCode, args)

}

// MakeBplusWarningWithErrorCode - returns a customized CAFUError for BPlus
func MakeBplusWarningWithErrorCode(ctx context.Context, httpErrorCode int, e BPlusErrorCode, args map[string]interface{}) wegoe.WeGOError {
	return internalMakeBplusError(ctx, wegoe.Warning, e, httpErrorCode, args)

}

// BPlusErrorCode - A BPlus error code
type BPlusErrorCode int

// enumeration for B Plus Error codes
const (
	CannotInvokeOperation         BPlusErrorCode = iota + 200000 // stringdemoapi.errors.CannotInvokeOperation
	CannotCastResponse                                           // stringdemoapi.errors.CannotCastResponse
	UnexpectedProxyInputParameter                                // stringdemoapi.errors.UnexpectedProxyInputParameter
)

//go:generate stringer -linecomment -type=BPlusErrorCode
