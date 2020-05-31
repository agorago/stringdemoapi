package register

import (
	"context"
	api "github.com/agorago/stringdemoapi/api"
	
	// that all WEGO modules are loaded
	fw "github.com/agorago/wego/fw"
	"reflect"
)

func RegisterService(wegoService fw.RegistrationService,proxyMiddlewares []fw.Middleware){
	wegoService.RegisterService("stringdemo",GetServiceDescriptor(proxyMiddlewares))
}
func GetServiceDescriptor(proxyMiddlewares []fw.Middleware) fw.ServiceDescriptor {
	return fw.ServiceDescriptor{
		Name:        "stringdemo",
		Description: " StringDemoService - the interface that is going to be implemented by the string demo service This has methods to illustrate features of the Wego framework",
		Operations:  operationDescriptors(proxyMiddlewares),
	}
}

func operationDescriptors(proxyMiddlewares []fw.Middleware) []fw.OperationDescriptor {
	return []fw.OperationDescriptor{

		{
			Name:                "Uppercase",
			Description:         " Uppercase - Converts the input string into upper case",
			URL:                 "/uppercase",
			HTTPMethod:          "POST",
			RequestDescription:  " UppercaseRequest - the payload for Uppercase service",
			ResponseDescription: " UppercaseResponse - the  Uppercase service response",
			OpRequestMaker:      makeUppercaseRequest,
			OpResponseMaker:     makeUppercaseResponse,
			Params:              uppercasePD(),
		},

		{
			Name:                "Count",
			Description:         " Count - returns the length of the input string",
			URL:                 "/count",
			HTTPMethod:          "POST",
			RequestDescription:  " CountRequest - the payload for Count service",
			ResponseDescription: " CountResponse - the  Count service response",
			OpRequestMaker:      makeCountRequest,
			OpResponseMaker:     makeCountResponse,
			ProxyMiddleware:     proxyMiddlewares,
			Params:              countPD(),
		},

		{
			Name:                "AddNumbers",
			Description:         " AddNumbers - adds two numbers and returns the result This method illustrates a GET method implementation in WeGO since there is no request payload required",
			URL:                 "/add-numbers",
			HTTPMethod:          "GET",
			RequestDescription:  "",
			ResponseDescription: " AddNumbersResponse - the  AddNumbers service response",

			OpResponseMaker: makeAddNumbersResponse,
			Params:          addNumbersPD(),
		},

		{
			Name:                "AddNumbers",
			Description:         " AddNumbers - adds two numbers and returns the result This method illustrates a GET method implementation in WeGO since there is no request payload required",
			URL:                 "/add-numbers-path/{Arg1}/{Arg2}",
			HTTPMethod:          "GET",
			RequestDescription:  "",
			ResponseDescription: " AddNumbersResponse - the  AddNumbers service response",

			OpResponseMaker: makeAddNumbersResponse,
			Params:          addNumbersPD(),
		},

	}
}

func uppercasePD() []fw.ParamDescriptor {

	return []fw.ParamDescriptor{

		{
			Name:        "ctx",
			Description: "",
			ParamOrigin: fw.CONTEXT,
		},

		{
			Name:        "ucr",
			Description: "",
			ParamOrigin: fw.PAYLOAD,
		},
	}
}

func countPD() []fw.ParamDescriptor {

	return []fw.ParamDescriptor{

		{
			Name:        "ctx",
			Description: "",
			ParamOrigin: fw.CONTEXT,
		},

		{
			Name:        "cr",
			Description: "",
			ParamOrigin: fw.PAYLOAD,
		},
	}
}

func addNumbersPD() []fw.ParamDescriptor {

	return []fw.ParamDescriptor{

		{
			Name:        "ctx",
			Description: "",
			ParamOrigin: fw.CONTEXT,
		},

		{
			Name:        "Arg1",
			Description: "",
			ParamOrigin: fw.HEADER,
			ParamKind:   reflect.Int,
		},

		{
			Name:        "Arg2",
			Description: "",
			ParamOrigin: fw.HEADER,
			ParamKind:   reflect.Int,
		},
	}
}

func makeUppercaseRequest(context.Context) (interface{}, error) {
	return &api.UppercaseRequest{}, nil
}

func makeUppercaseResponse(context.Context) (interface{}, error) {
	return &api.UppercaseResponse{}, nil
}
func makeCountRequest(context.Context) (interface{}, error) {
	return &api.CountRequest{}, nil
}

func makeCountResponse(context.Context) (interface{}, error) {
	return &api.CountResponse{}, nil
}

func makeAddNumbersResponse(context.Context) (interface{}, error) {
	return &api.AddNumbersResponse{}, nil
}
