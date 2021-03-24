package flag

import (
	"strings"
)

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

func pathToMap(cfg map[string]interface{}) map[string]interface{} {
	root := make(map[string]interface{})
	for k, v := range cfg {
		names := strings.Split(k, ".")
		n := get_path_node(root, names)

		n[names[len(names)-1]] = v
	}
	return root
}

func dumpFlag() {
	cfg := make(map[string]interface{})
	CommandLine.VisitAll(func(f *Flag) {
		name := f.Name
		if name == "flagdump" {
			return
		} else if name == "flagenv" {
			return
		} else if name == "flagfile" {
			return
		}
		cfg[f.Name] = f.Value
	})

	switch *g_flagdump {
	case "env":
		dumpEnvFlag(cfg)
	case "toml":
		dumpTomlFlag(cfg)
	case "yaml":
		dumpYamlFlag(cfg)
	case "json":
		dumpJsonFlag(cfg)
	case "ini":
		dumpIniFlag(cfg)
	case "flag":
		fallthrough
	default:
		dumpRawFlag(cfg)
	}

}
