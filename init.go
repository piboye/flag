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
		for _, arg := range os.Args[1:] {
			if arg == "-h" || arg == "-help" {
				commandLineUsage()
				return
			}
		}
	*/

}
