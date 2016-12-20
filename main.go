package main

import (
	"github.com/myafeier/negroni"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/myafeier/meeting/handler"
	"github.com/myafeier/meeting/config"

	"github.com/myafeier/meeting/model"
)



func init()  {
	config.InitConfig()
	model.InitRedis()

}

func main()  {


	logger:=negroni.NewLogger()
	recover:=&negroni.Recovery{
		Logger:     logger,
		PrintStack: false,
		StackAll:   false,
		StackSize:  1024 * 8,
	}
	recover.Logger=logger
	static:=&negroni.Static{
		Dir:       http.Dir("static"),
		Prefix:    "/static",
		IndexFile: "index.html",
	}
	n:=negroni.New(logger,recover,static)
	mux:=mux.NewRouter()
	//路由设置
	{
	//
		mux.HandleFunc("/",handler.ScreenIndex)

		mux.HandleFunc("/qr",handler.QrCodeShowHandler)
		mux.HandleFunc("/login",handler.Login)
		mux.Handle("/admin",handler.AuthHandler(http.HandlerFunc(handler.AdminIndex)))
		mux.HandleFunc("/ws",handler.Websocket)

	}

	n.UseHandler(mux)
	n.Run(":"+config.SiteCommonConfig.ListenPort)

}


