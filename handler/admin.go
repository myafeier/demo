package handler

import (
	"net/http"
	"html/template"
	"github.com/myafeier/meeting/config"
)

func AdminIndex(w http.ResponseWriter,r *http.Request){
	if r.Method==http.MethodGet{
		tp:=template.Must(template.ParseFiles(config.SiteCommonConfig.TemplatePath+"/admin/index.html"))
		data:=make(map[string]string)
		data["username"]=r.Context().Value("username").(string)
		tp.Execute(w,data)
		return
	}


}
