package main

import (
	"log"
	"github.com/MiSingularity/logging/fe"
	"github.com/MiSingularity/logging/p"
	"time"
)

func main() {
	log.Println("init bi client")
	fe.Init("127.0.0.1", "8999")
	log.Println("add a bi log")
	fe.Bi(&p.BiLog{
		ProjectName: "deepshare",
		ActionName:  "userlink",
		Timestamp:   time.Now().Unix(),
		Detail:      []byte("detail~~~~~~"),
	})
}