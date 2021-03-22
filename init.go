package flag

import (
	"fmt"
	"os"
)

func (fs *FlagSet) Config() []string {
	cfg := make([]string, 0, 10)
	fs.VisitAll(func(f *Flag) {
		cfg = append(cfg, fmt.Sprintf("%s:%q", f.Name, f.Value.String()))
	})

	return cfg
}

func init() {

	if CommandLine.values == nil {
		CommandLine.values = make(map[string]string)
	}
	values := CommandLine.values
	CommandLine.preParse(os.Args[1:])

	preParseFile(values)
	preParseEnv(values)

	/*
		for _, arg := range os.Args[1:] {
			if arg == "-h" || arg == "-help" {
				commandLineUsage()
				return
			}
		}
	*/

}
