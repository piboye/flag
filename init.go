package flag

import (
	"fmt"
	"os"
)

func (fs *FlagSet) Config() []string {
	cfg := make([]string, 0, 10)
	fs.VisitAll(func(f *Flag) {
		name := f.Name
		if name == "flagdump" {
			return
		} else if name == "flagenv" {
			return
		} else if name == "flagfile" {
			return
		}
		cfg = append(cfg, fmt.Sprintf("%s:%q", f.Name, f.Value.String()))
	})

	return cfg
}

var flagfile = String("flagfile", "flagfile", "parse flag from file,  support file-format: flagifle|conf.json|conf.toml|conf.yaml")
var g_flagdump = String("flagdump", "flag", "dump all flags value to stdout, support json|toml|yaml|flag|env")
var g_flagenv = Bool("flagenv", true, "parse flag from env var or .env file")

func init() {

	if CommandLine.values == nil {
		CommandLine.values = make(map[string]string)
	}

	values := CommandLine.values

	preParseFile(values)
	preParseEnv(values)
	CommandLine.preParse(os.Args[1:])

	/*
		for _, arg := range os.Args[1:] {
			if arg == "-h" || arg == "-help" {
				commandLineUsage()
				return
			}
		}
	*/

}
