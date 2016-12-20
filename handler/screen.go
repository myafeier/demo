package handler

import (
	"net/http"
	"html/template"
	"github.com/myafeier/meeting/config"
)

func ScreenIndex(w http.ResponseWriter,r *http.Request){
	tp:=template.Must(template.ParseFiles(config.SiteCommonConfig.TemplatePath+"/index/index.html"))
	tp.Execute(w,r.Context())
}
