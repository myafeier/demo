package handler

import (
	"net/http"
	"html/template"
	"github.com/myafeier/meeting/config"
)

func Login(w http.ResponseWriter,r *http.Request){

	if r.Method==http.MethodGet{
		tp:=template.Must(template.ParseFiles(config.SiteCommonConfig.TemplatePath+"/login.html"))
		tp.Execute(w,nil)
		return
	}else{


		cookie:=new(http.Cookie)
		cookie.Name="username"
		cookie.Value="xiafei"
		cookie.Path="/"
		cookie.MaxAge=1000
		http.SetCookie(w,cookie)
		http.Redirect(w,r,"/admin",302)
		return


	}
}
