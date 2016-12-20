package handler

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"io"
)

func Websocket(w http.ResponseWriter,r *http.Request){
	upgrader:=&websocket.Upgrader{
		ReadBufferSize:1024,
		WriteBufferSize:1024,
	}
	conn,err:=upgrader.Upgrade(w,r,nil)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	defer conn.Close()

	for{
		_,p,err:=conn.ReadMessage()
		if err!=nil{
			if err!=io.EOF{
				log.Println("Connection Disconnect!",err)
				w.WriteHeader(400)
				return
			}
			continue
		}
		log.Println(p)


	}
}
