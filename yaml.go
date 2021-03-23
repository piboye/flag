package flag

import (
	"fmt"
	"io/ioutil"

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

func dumpYamlFlag(cfg map[string]string) {
	root := pathToMap(cfg)
	txt, _ := yaml.Marshal(root)
	fmt.Printf("%s\n", txt)
}
