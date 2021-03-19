package main

import (
	"log"

	"github.com/piboye/flag"
)

var testb = flag.Bool("testb", true, "test bool")

func init() {
	var addr = ""
	flag.StringVar(&addr, "redis.addr", "empty", "redis address")
	log.Printf("redis addr:%s", addr)
	log.Printf("testb %t", *testb)
}

func main() {

	//log.Printf("testb %t", *testb)
	flag.Parse()
	//log.Printf("testb %t", *testb)
}
