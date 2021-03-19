package flag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
