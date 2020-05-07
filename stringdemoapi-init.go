package stringdemoapi

import (
	"github.com/agorago/stringdemoapi/api"
	"github.com/agorago/stringdemoapi/proxy"
	"github.com/agorago/stringdemoapi/register"
	"github.com/agorago/wego"
	"github.com/agorago/wego/fw"
)

const (
	StringDemoApi = "StringDemoApi"
	StringDemoProxy = "StringDemoProxy"
)

func MakeStringDemoApiInitializer()fw.Initializer{
	return stringDemoApiInitializer{}
}

type stringDemoApiInitializer struct{}

func (stringDemoApiInitializer)ModuleName() string{
	return StringDemoApi
}
// The stringdemoapi initializer
func (stringDemoApiInitializer)Initialize(commandCatalog fw.CommandCatalog)(fw.CommandCatalog,error){
	wegoService,err := wego.GetWego(commandCatalog)
	if err != nil {
		return commandCatalog,nil
	}
	wegoService.RegisterService("stringdemo",register.GetServiceDescriptor())
	// create a proxy
	proxyService,err := wego.GetProxyService(commandCatalog)
	stringdemoProxy := proxy.MakeStringdemoProxy(proxyService)
	commandCatalog.RegisterCommand(StringDemoProxy,stringdemoProxy)
	return commandCatalog,nil
}

func GetStringDemoProxy(commandCatalog fw.CommandCatalog)(api.StringDemoService,error){
	return commandCatalog.Command(StringDemoProxy).(api.StringDemoService),nil
}