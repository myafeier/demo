package handler

import (
	"net/http"
	"github.com/honsty/hqr"
	"strconv"
	"log"
)

func QrCodeShowHandler(w http.ResponseWriter,r *http.Request){
	code:="http://test.diakor.cc"
	c,err:=hqr.Encode(code,hqr.L)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	pngData:=c.PNG()

	//w.WriteHeader(200)
	w.Header().Set("Content-Type","image/png")
	w.Header().Set("Cache-Control","no-cache")
	length,err:=w.Write(pngData)
	w.Header().Set("Content-Length",strconv.Itoa(length))
}
