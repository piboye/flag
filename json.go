package flag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-jsonnet"
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

		/*
			v, ok := values[name]
			if ok && len(v) > 0 {
				continue
			}
		*/

		values[name] = str
	}
}

func tryParseJson(filename string, values map[string]string) error {
	var data []byte
	var err error
	if filename == "-" {
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
	}

	root := make(map[string]interface{})

	err = json.Unmarshal(data, &root)
	if err != nil {
		return err
	}

	dftJson(root, "", values)
	return nil
}

func renderJsonnet(file string) (string, error) {

	// Create a JSonnet VM
	vm := jsonnet.MakeVM()

	// render the jsonnet
	out, err := vm.EvaluateFile(file)

	if err != nil {
		log.Panic("Error evaluating jsonnet snippet: ", err)
		return "", err
	}

	return out, nil

}

func tryParseJsonnet(filename string, values map[string]string) error {
	data, err := renderJsonnet(filename)
	if err != nil {
		return err
	}

	root := make(map[string]interface{})

	err = json.Unmarshal([]byte(data), &root)
	if err != nil {
		return err
	}

	dftJson(root, "", values)
	return nil
}

func dumpJsonFlag(cfg map[string]interface{}) {
	root := pathToMap(cfg)
	txt, _ := json.Marshal(root)
	fmt.Printf("%s\n", txt)
}
