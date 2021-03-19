package flag

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func dftYaml(root map[string]interface{}, prefix string, values map[string]string) {
	if root == nil {
		return
	}

	for key, val := range root {
		name := key
		if len(prefix) > 0 {
			name = prefix + "." + key
		}

		if child, ok := val.(map[string]interface{}); ok {
			dftYaml(child, name, values)
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

func tryParseYaml(filename string, values map[string]string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	root := make(map[string]interface{})

	err = yaml.Unmarshal(data, root)
	if err != nil {
		return err
	}

	dftJson(root, "", values)
	return nil
}

func preParseYaml(values map[string]string) bool {
	var fp = String("flagyaml", "", "yaml file for flag")

	filename := getRawFlagValue("flagyaml")

	if len(filename) > 0 {
		err := tryParseYaml(filename, values)
		if err != nil {
			log.Fatalf("flagyaml load yaml file failed, [file=%s][err=%+v]", filename, err)
		}
		return true
	}

	filename = *fp
	if len(filename) <= 0 {
		return false
	}

	err := tryParseYaml(filename, values)
	return err == nil

}
