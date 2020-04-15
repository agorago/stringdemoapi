package register

import (
	"context"
	api "gitlab.intelligentb.com/examples/stringdemo/stringdemoapi/api"
	_ "gitlab.intelligentb.com/devops/bplus" // initialize BPlus first to make sure
	"gitlab.intelligentb.com/examples/stringdemo/stringdemoapi/proxy"

	// that all BPLUS modules are loaded
	bplus "gitlab.intelligentb.com/devops/bplus/fw"
	"reflect"
)

func init() {
	var sd = GetServiceDescriptor()

	bplus.RegisterService("stringdemo", sd)
}

func GetServiceDescriptor() bplus.ServiceDescriptor {
	return bplus.ServiceDescriptor{
		Name:            "stringdemo",
		Description:     " StringDemoService - the interface that is going to be implemented by the string demo service This has methods to illustrate features of the BPlus framework",
		Operations:      OperationDescriptors(),
	}
}

func OperationDescriptors() []bplus.OperationDescriptor {
	return []bplus.OperationDescriptor{

		bplus.OperationDescriptor{
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

		bplus.OperationDescriptor{
			Name:                "Count",
			Description:         " Count - returns the length of the input string",
			URL:                 "/count",
			HTTPMethod:          "POST",
			RequestDescription:  " CountRequest - the payload for Count service",
			ResponseDescription: " CountResponse - the  Count service response",
			OpRequestMaker:      makeCountRequest,
			OpResponseMaker:     makeCountResponse,
			ProxyMiddleware:     []bplus.Middleware{proxy.InterceptCount,},
			Params:              CountPD(),
		},

		bplus.OperationDescriptor{
			Name:                "AddNumbers",
			Description:         " AddNumbers - adds two numbers and returns the result This method illustrates a GET method implementation in BPlus since there is no request payload required",
			URL:                 "/add-numbers",
			HTTPMethod:          "GET",
			RequestDescription:  "",
			ResponseDescription: " AddNumbersResponse - the  AddNumbers service response",

			OpResponseMaker: makeAddNumbersResponse,
			Params:          AddNumbersPD(),
		},

		bplus.OperationDescriptor{
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

func UppercasePD() []bplus.ParamDescriptor {

	return []bplus.ParamDescriptor{

		bplus.ParamDescriptor{
			Name:        "ctx",
			Description: "",
			ParamOrigin: bplus.CONTEXT,
		},

		bplus.ParamDescriptor{
			Name:        "ucr",
			Description: "",
			ParamOrigin: bplus.PAYLOAD,
		},
	}
}

func CountPD() []bplus.ParamDescriptor {

	return []bplus.ParamDescriptor{

		bplus.ParamDescriptor{
			Name:        "ctx",
			Description: "",
			ParamOrigin: bplus.CONTEXT,
		},

		bplus.ParamDescriptor{
			Name:        "cr",
			Description: "",
			ParamOrigin: bplus.PAYLOAD,
		},
	}
}

func AddNumbersPD() []bplus.ParamDescriptor {

	return []bplus.ParamDescriptor{

		bplus.ParamDescriptor{
			Name:        "ctx",
			Description: "",
			ParamOrigin: bplus.CONTEXT,
		},

		bplus.ParamDescriptor{
			Name:        "Arg1",
			Description: "",
			ParamOrigin: bplus.HEADER,
			ParamKind:   reflect.Int,
		},

		bplus.ParamDescriptor{
			Name:        "Arg2",
			Description: "",
			ParamOrigin: bplus.HEADER,
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
