package flag

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var flagenv = Bool("flagenv", true, "parse flag from env var or .env file")

func preParseEnv(values map[string]string) bool {
	if !*flagenv {
		return false
	}

	godotenv.Load(".env")

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

func dumpEnvFlag(cfg map[string]interface{}) {
	out := make(map[string]string)
	for k, v := range cfg {

		v1 := v.(Value).String()
		if strings.Index(k, ".") >= 0 {
			k = strings.ReplaceAll(k, ".", "__")
		}
		out[k] = v1
	}
	txt, err := godotenv.Marshal(out)
	if err != nil {
		return
	}

	fmt.Printf("%s\n", txt)
}
