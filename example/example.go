package main

import (
	"log"

	"github.com/piboye/flag"
)

var testb = flag.Bool("testb", true, "test bool")
var addr = ""

func init() {
	flag.StringVar(&addr, "redis.addr", "empty", "redis address")
}

func main() {
	flag.Parse()
	log.Printf("redis addr:%s", addr)
	log.Printf("testb %t", *testb)
}
