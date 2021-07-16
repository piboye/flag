package flag

import (
	"testing"
)

var test = Bool("test", false, "test")

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
	var pf = String("pf", "flagfile", "flagfile")

	if *pf != "flagfile" {
		t.Errorf("bool default value failed!, [expect=flagfile] [cur=%+v]", *pf)
	}
}
