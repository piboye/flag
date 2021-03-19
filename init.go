package flag

import "os"

func init() {
	if CommandLine.values == nil {
		CommandLine.values = make(map[string]string)
	}
	values := CommandLine.values

	CommandLine.preParse(os.Args[1:])

	preParseFile(values)
	preParseEnv(values)

	/*
		preParseFlagFile(values)
		preParseToml(values)
		preParseJson(values)
		preParseYaml(values)
	*/
}
