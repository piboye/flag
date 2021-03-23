package flag

import (
	"fmt"
	"strings"
)

var g_flagdump = String("flagdump", "flag", "dump all flags value to stdout, support json|toml|yaml|flag|env")

func dumpRawFlag(cfg map[string]string) {
	for k, v := range cfg {
		fmt.Printf("%s=%s\n", k, v)
	}
}
func get_path_node(root map[string]interface{}, names []string) map[string]interface{} {
	if len(names) < 2 {
		return root
	}

	name := names[0]

	if node, ok := root[name]; ok {
		return get_path_node(node.(map[string]interface{}), names[1:])
	} else {
		node := make(map[string]interface{})
		root[name] = node
		return get_path_node(node, names[1:])
	}
}

func pathToMap(cfg map[string]string) map[string]interface{} {
	root := make(map[string]interface{})
	for k, v := range cfg {
		names := strings.Split(k, ".")
		n := get_path_node(root, names)

		n[names[len(names)-1]] = v
	}
	return root
}

func dumpFlag() {
	cfg := make(map[string]string)
	CommandLine.VisitAll(func(f *Flag) {
		name := f.Name
		if name == "flagdump" {
			return
		} else if name == "flagenv" {
			return
		} else if name == "flagfile" {
			return
		}
		v := f.Value.String()
		if *g_flagdump == "env" {
			cfg[f.Name] = v
			return
		}
		if strings.HasPrefix(v, "\"") {
			v = v[1:]
		}

		if strings.HasSuffix(v, "\"") {
			v = v[0 : len(v)-1]
		}
		cfg[f.Name] = v
	})

	switch *g_flagdump {
	case "toml":
		dumpTomlFlag(cfg)
	case "yaml":
		dumpYamlFlag(cfg)
	case "json":
		dumpJsonFlag(cfg)
	case "flag":
		fallthrough
	default:
		dumpRawFlag(cfg)
	}

}
