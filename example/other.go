package main

import (
	"log"

	//"github.com/piboye/flag"
	"flag"
)

var hookable = true

func init() {
	flag.BoolVar(&hookable, "hookable", true, "test bool")
	log.Printf("hookable %t", hookable)
}
