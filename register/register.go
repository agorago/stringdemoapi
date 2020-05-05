package register

import (
	"context"
	_ "github.com/agorago/wego" // initialize Wego first to make sure
	api "github.com/agorago/stringdemoapi/api"
	"github.com/agorago/stringdemoapi/proxy"

	// that all WEGO modules are loaded
	fw "github.com/agorago/wego/fw"
	"reflect"
)

type StringDemoRegistration interface{
	Register(sd fw.ServiceDescriptor)
}

type stringDemoRegistration struct{
	Wego fw.RegistrationService
}

func GetServiceDescriptor() fw.ServiceDescriptor {
	return fw.ServiceDescriptor{
		Name:        "stringdemo",
		Description: " StringDemoService - the interface that is going to be implemented by the string demo service This has methods to illustrate features of the BPlus framework",
		Operations:  OperationDescriptors(),
	}
}

func (sdr stringDemoRegistration)Register(sd fw.ServiceDescriptor){
	sdr.Wego.RegisterService("stringdemo",sd)
}

func OperationDescriptors() []fw.OperationDescriptor {
	return []fw.OperationDescriptor{

		fw.OperationDescriptor{
			Name:                "Uppercase",
			Description:         " Uppercase - Converts the input string into upper case",
			URL:                 "/uppercase",
			HTTPMethod:          "POST",
			RequestDescription:  " UpperCaseRequest - the payload for Uppercase service",
			ResponseDescription: " UpperCaseResponse - the  Uppercase service response",
			OpRequestMaker:      makeUppercaseRequest,
			OpResponseMaker:     makeUppercaseResponse,
			Params:              UppercasePD(),
		},

		fw.OperationDescriptor{
			Name:                "Count",
			Description:         " Count - returns the length of the input string",
			URL:                 "/count",
			HTTPMethod:          "POST",
			RequestDescription:  " CountRequest - the payload for Count service",
			ResponseDescription: " CountResponse - the  Count service response",
			OpRequestMaker:      makeCountRequest,
			OpResponseMaker:     makeCountResponse,
			ProxyMiddleware:     []fw.Middleware{proxy.InterceptCount},
			Params:              CountPD(),
		},

		fw.OperationDescriptor{
			Name:                "AddNumbers",
			Description:         " AddNumbers - adds two numbers and returns the result This method illustrates a GET method implementation in BPlus since there is no request payload required",
			URL:                 "/add-numbers",
			HTTPMethod:          "GET",
			RequestDescription:  "",
			ResponseDescription: " AddNumbersResponse - the  AddNumbers service response",

			OpResponseMaker: makeAddNumbersResponse,
			Params:          AddNumbersPD(),
		},

		fw.OperationDescriptor{
			Name:                "AddNumbers",
			Description:         " AddNumbers - adds two numbers and returns the result This method illustrates a GET method implementation in BPlus since there is no request payload required",
			URL:                 "/add-numbers-path/{Arg1}/{Arg2}",
			HTTPMethod:          "GET",
			RequestDescription:  "",
			ResponseDescription: " AddNumbersResponse - the  AddNumbers service response",

			OpResponseMaker: makeAddNumbersResponse,
			Params:          AddNumbersPD(),
		},
	}
}

func UppercasePD() []fw.ParamDescriptor {

	return []fw.ParamDescriptor{

		fw.ParamDescriptor{
			Name:        "ctx",
			Description: "",
			ParamOrigin: fw.CONTEXT,
		},

		fw.ParamDescriptor{
			Name:        "ucr",
			Description: "",
			ParamOrigin: fw.PAYLOAD,
		},
	}
}

func CountPD() []fw.ParamDescriptor {

	return []fw.ParamDescriptor{

		fw.ParamDescriptor{
			Name:        "ctx",
			Description: "",
			ParamOrigin: fw.CONTEXT,
		},

		fw.ParamDescriptor{
			Name:        "cr",
			Description: "",
			ParamOrigin: fw.PAYLOAD,
		},
	}
}

func AddNumbersPD() []fw.ParamDescriptor {

	return []fw.ParamDescriptor{

		fw.ParamDescriptor{
			Name:        "ctx",
			Description: "",
			ParamOrigin: fw.CONTEXT,
		},

		fw.ParamDescriptor{
			Name:        "Arg1",
			Description: "",
			ParamOrigin: fw.HEADER,
			ParamKind:   reflect.Int,
		},

		fw.ParamDescriptor{
			Name:        "Arg2",
			Description: "",
			ParamOrigin: fw.HEADER,
			ParamKind:   reflect.Int,
		},
	}
}

func makeUppercaseRequest(ctx context.Context) (interface{}, error) {
	return &api.UpperCaseRequest{}, nil
}

func makeUppercaseResponse(ctx context.Context) (interface{}, error) {
	return &api.UpperCaseResponse{}, nil
}
func makeCountRequest(ctx context.Context) (interface{}, error) {
	return &api.CountRequest{}, nil
}

func makeCountResponse(ctx context.Context) (interface{}, error) {
	return &api.CountResponse{}, nil
}

func makeAddNumbersResponse(ctx context.Context) (interface{}, error) {
	return &api.AddNumbersResponse{}, nil
}
