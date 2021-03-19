package flag

import (
	"testing"
	//flag "github.com/piboye/flag"
)

var test = Bool("test", false, "test")
var flagfile = String("flagfile", "flagfile", "flagfile")

func init() {
	//log.Printf("test=%+v", *test)
	//log.Printf("flagfile=%+v", *flagfile)
}

func TestBool(t *testing.T) {
	var boolt = false
	BoolVar(&boolt, "boolt", true, "boolt")

	if boolt == false {
		t.Errorf("bool default value failed!, [expect=true] [cur=%+v]", boolt)
	}
}

func TestString(t *testing.T) {

	if *flagfile != "flagfile" {
		t.Errorf("bool default value failed!, [expect=flagfile] [cur=%+v]", *flagfile)
	}
}
