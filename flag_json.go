package flag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func dftJson(root map[string]interface{}, prefix string, values map[string]string) {
	if root == nil {
		return
	}

	for key, val := range root {
		name := key
		if len(prefix) > 0 {
			name = prefix + "." + key
		}

		if child, ok := val.(map[string]interface{}); ok {
			dftJson(child, name, values)
			continue
		}

		str := fmt.Sprintf("%+v", val)

		v, ok := values[name]
		if ok && len(v) > 0 {
			continue
		}

		values[name] = str
	}
}

func tryParseJson(filename string, values map[string]string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	root := make(map[string]interface{})

	err = json.Unmarshal(data, &root)
	if err != nil {
		return err
	}

	dftJson(root, "", values)
	return nil
}

func preParseJson(values map[string]string) bool {
	var fp = String("flagjson", "", "json file for flag")
	//var fp = String("flagtoml", "config.toml", "tomlfile for flag")

	filename := getRawFlagValue("flagjson")

	if len(filename) > 0 {
		err := tryParseJson(filename, values)
		if err != nil {
			log.Fatalf("flagjson load json file failed, [file=%s][err=%+v]", filename, err)
		}
		return true
	}

	//flagfile := CommandLine.actual["flagtoml"]

	//filename = flagfile.Value.String()
	filename = *fp
	if len(filename) <= 0 {
		return false
	}

	err := tryParseJson(filename, values)
	return err == nil

}
