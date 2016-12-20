package handler

import (
	"net/http"
	"log"
	"context"
)

func AuthHandler(next http.Handler)http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter,r *http.Request) {
			cookieUsername,err:=r.Cookie("username")
			if err!=nil{
				if err!=http.ErrNoCookie{
					log.Println(err)
					w.WriteHeader(500)
					return
				}
				http.Redirect(w,r,"/login",302)
				return
			}

			if cookieUsername.Value==""{
				http.Redirect(w,r,"/login",302)
				return
			}

			key:="username"
			ctx:=context.WithValue(r.Context(),key,cookieUsername.Value)
			log.Println(ctx.Value(key))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
