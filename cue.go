package flag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"cuelang.org/go/cue"
)

func CheckFlag() {
	/*
		var r cue.Runtime
		i, err := r.Compile("")
		if err {

		}

		var codec = gocodec.New(r, nil)

		codec.Validate(i, i)
	*/
}

var g_cue_r cue.Runtime

func tryParseCue(filename string, values map[string]string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	str := string(data)

	conf, err := g_cue_r.Compile("test", str)
	if err != nil {
		return err
	}

	txt, err := conf.Value().MarshalJSON()
	if err != nil {
		return err
	}

	root := make(map[string]interface{})
	json.Unmarshal(txt, &root)

	dftJson(root, "", values)
	return nil
}

func dumpCueFlag(cfg map[string]interface{}) {
	root := pathToMap(cfg)
	txt, _ := json.Marshal(root)
	fmt.Printf("%s\n", txt)
}
