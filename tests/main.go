package main

import (
	"github.com/xiangrui2019/xgetenv"
	"log"
)

func main() {
	log.Println(xgetenv.XGetenv("GOPArTH", "u"))
}
