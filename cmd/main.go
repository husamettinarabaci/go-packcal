package main

import (
	"sync"

	"github.com/golobby/container/v3"

	pa "github.com/husamettinarabaci/go-packcal/core/application/presentation/adapter"
	as "github.com/husamettinarabaci/go-packcal/core/application/service"
	ds "github.com/husamettinarabaci/go-packcal/core/domain/service"
	ia "github.com/husamettinarabaci/go-packcal/pkg/infrastructure/adapter"
	cr "github.com/husamettinarabaci/go-packcal/pkg/presentation/controller/rest"
	cw "github.com/husamettinarabaci/go-packcal/pkg/presentation/controller/web"
	tconfig "github.com/husamettinarabaci/go-packcal/tool/config"
)

var restConfig tconfig.RestConfig
var webConfig tconfig.WebConfig

func main() {
	restConfig.ReadConfig()
	webConfig.ReadConfig()
	var err error
	cont := container.New()

	//Domain PackCal Service
	err = cont.Singleton(func() ds.Service {
		return ds.NewService()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure PackCal Calc Adapter
	err = cont.Singleton(func() ia.CalcAdapter {
		return ia.NewCalcAdapter()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure PackCal Log Adapter
	err = cont.Singleton(func() ia.LogAdapter {
		return ia.NewLogAdapter()
	})
	if err != nil {
		panic(err)
	}

	//Application PackCal Service
	err = cont.Singleton(func(s ds.Service, i ia.CalcAdapter, l ia.LogAdapter) as.Service {
		return as.NewService(s, i, l)
	})
	if err != nil {
		panic(err)
	}

	//Application PackCal Query Adapter
	err = cont.Singleton(func(s as.Service) pa.QueryAdapter {
		return pa.NewQueryAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	//Application PackCal Calc Adapter
	err = cont.Singleton(func(s as.Service) pa.CalcAdapter {
		return pa.NewCalcAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	var queryHandler pa.QueryAdapter
	err = cont.Resolve(&queryHandler)
	if err != nil {
		panic(err)
	}

	var calcHandler pa.CalcAdapter
	err = cont.Resolve(&calcHandler)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go cw.NewWebServer(queryHandler, calcHandler).Serve(webConfig.Debug, webConfig.WebServer.Port)
	wg.Add(1)
	go cr.NewRestAPI(queryHandler, calcHandler).Serve(restConfig.Debug, restConfig.Restapi.Port)
	wg.Wait()
}
