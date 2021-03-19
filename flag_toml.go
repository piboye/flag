package flag

import (
	"fmt"
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

func dftToml(root *toml.Tree, prefix string, values map[string]string) {
	if root == nil {
		return
	}

	for _, key := range root.Keys() {

		name := key
		if len(prefix) > 0 {
			name = prefix + "." + key
		}

		//log.Printf("[key=%s]", name) //[value=%s]", key, str)

		val := root.Get(key)

		child, ok := val.(*toml.Tree)
		if ok {
			dftToml(child, name, values)
			continue
		}

		//log.Printf("val:%+v", val)

		str := fmt.Sprintf("%+v", val)

		//log.Printf("[key=%s][value=%s]", name, str) //[value=%s]", key, str)

		v, ok := values[name]
		if ok && len(v) > 0 {
			continue
		}

		values[name] = str
	}
}

func tryParseToml(filename string, values map[string]string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	str := string(data)

	conf, err := toml.Load(str)
	if err != nil {
		return err
	}

	dftToml(conf, "", values)
	return nil
}
