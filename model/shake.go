package model

import (
	"time"
	"strconv"
	"log"
)

type Shake struct {
	No int   `json:"no,omitempty"`
	Id int64 `json:"id,omitempty"`
	X float64 	`json:"x,omitempty"`
	Y float64 	`json:"y,omitempty"`
	Z float64 	`json:"z,omitempty"`
	S float64	`json:"s,omitempty"`	//speed
	AtTime time.Time `json:"at_time,omitempty"`
}


//func (self *Shake)Insert()(num int64,err error){
//	num,err=Db.Insert(self)
//	return
//}
//
//func (self *Shake)GetListByPage(page int)(result []*Shake,err error){
//
//	rows,err:=Db.OrderBy("id desc").Limit(10,10*page).Rows(self)
//	if err!=nil{
//		return
//	}
//	defer rows.Close()
//	for rows.Next(){
//		o:=new(Shake)
//		err=rows.Scan(o)
//		if err!=nil{
//			return
//		}
//		result=append(result,o)
//	}
//	return
//}

func (self *Shake)InsertRedis()(result int64,err error){
	reConn:=res.Get()
	defer reConn.Close()

	_,err=reConn.Do("incr","kye_"+strconv.Itoa(self.No))
	if err!=nil{
		log.Println(err)
		return
	}
	return
}
