package flag

import (
	"io/ioutil"
	"os"
	"strings"
)

func preParseEnv(values map[string]string) bool {
	var flagenvvar = Bool("flagenv", true, "parse flag from env var or .env file")
	if !*flagenvvar {
		return false
	}

	data, err := ioutil.ReadFile(".env")
	if err == nil {
		str := string(data)
		new_vals := readFlagString(str)
		//log.Printf("map env %+v", new_vals)

		for k, v := range new_vals {
			os.Setenv(k, v)
			k = strings.ReplaceAll(k, ".", "__")
			os.Setenv(k, v)
		}
	}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 2 {
			continue
		}

		k := pair[0]

		v := pair[1]
		_, ok := values[k]
		if ok {
			continue
		}
		values[k] = v

		if strings.Index(k, "__") < 0 {
			continue
		}

		k = strings.ReplaceAll(k, "__", ".")
		_, ok = values[k]
		if ok {
			continue
		}
		values[k] = v
	}

	return true
}
