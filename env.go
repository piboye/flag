package flag

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func preParseEnv(values map[string]string) bool {
	if !*g_flagenv {
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

		if !strings.Contains(k, "__") {
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
		if strings.Contains(k, ".") {
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
