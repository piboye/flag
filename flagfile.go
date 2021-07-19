package flag

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func readFlagString(data string) map[string]string {
	lines := strings.Split(data, "\n")
	result := make(map[string]string, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || trimmed[0] == '#' {
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) < 2 {
			log.Fatalf("Invalid config line: %s", line)
		}
		result[strings.TrimSpace(parts[0])] = strings.TrimSpace(strings.Join(parts[1:], "="))
	}
	return result
}

func tryParseFlagFile(filename string, values map[string]string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	str := string(data)

	new_vals := readFlagString(str)
	//log.Printf("flagfile map:%+v", new_vals)

	for k, v := range new_vals {
		_, ok := values[k]
		if ok {
			continue
		}
		values[k] = v
	}
	return nil
}

func preParseFlagFile(values map[string]string) bool {
	fp := String("flagfile", ".flagfile", "flagfile")
	filename := getRawFlagValue("flagfile")
	//log.Printf("filename=%s", filename)
	if len(filename) > 0 {
		err := tryParseFlagFile(filename, values)
		if err != nil {
			log.Fatalf("flagfile read file failed, [file=%s][err=%+v]", filename, err)
		}
		return true
	}

	//flagfile := CommandLine.actual["flagfile"]

	//filename = flagfile.Value.String()
	filename = *fp
	if len(filename) <= 0 {
		return false
	}

	err := tryParseFlagFile(filename, values)
	return err == nil
}

func tryParseFile(filename string, values map[string]string) error {
	ext := filepath.Ext(filename)
	//log.Printf("ext:%s", ext)
	switch ext {
	case ".ini":
		return tryParseIni(filename, values)
	case ".json":
		return tryParseJson(filename, values)
	case ".jsonnet":
		return tryParseJsonnet(filename, values)
	case ".cue":
		return tryParseCue(filename, values)
	case ".yaml":
		return tryParseYaml(filename, values)
	case ".toml":
		return tryParseToml(filename, values)
	default:
		if filename == "-" {
			return tryParseJson(filename, values)
		}
		return tryParseFlagFile(filename, values)
	}
}

func preParseFile(values map[string]string) bool {
	filename := getRawFlagValue("flagfile")

	if len(filename) > 0 {
		for _, fn := range strings.Split(filename, ",") {
			if len(fn) <= 0 {
				continue
			}
			err := tryParseFile(fn, values)
			if err != nil {
				log.Fatalf("flagfile parse file failed, [file=%s][err=%+v]", fn, err)
			}
		}
		return true
	}

	filename = *flagfile
	if len(filename) <= 0 {
		return false
	}

	err := tryParseFile(filename, values)
	return err == nil
}

func dumpRawFlag(cfg map[string]interface{}) {
	for k, v := range cfg {
		fmt.Printf("%s=%+v\n", k, v)
	}
}
