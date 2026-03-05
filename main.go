package main

import (
	"flag"
	"net/http"
	"log"

)


func main(){
	var addr=flag.String("addr",":8080","http service address")
	flag.Parse()
	r:=newroom()
	go r.run()
	 log.Printl("staring webserver on:",*addr)
	 if err:= http.ListenAndServe(*addr,bil);err !=nil{
		log.Fatal("l/s",err)
	 }
}
