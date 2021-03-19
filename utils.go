package flag

import "os"

func getRawFlagValue(target string) (result string) {

	args := os.Args[1:]
	result = ""

R:
	for i, s := range args {

		if len(s) < 2 || s[0] != '-' {
			continue
		}

		numMinuses := 1
		if s[1] == '-' {
			numMinuses++
			if len(s) == 2 { // "--" terminates the flags
				break
			}
		}
		name := s[numMinuses:]
		if len(name) == 0 || name[0] == '-' || name[0] == '=' {
			// bad flag syntax
			continue
		}

		name2 := name
		for j := 1; j < len(name); j++ { // equals cannot be first
			if name[j] == '=' {
				name2 = name[0:j]
				if name2 == target {
					return name[j+1:]
				}
				continue R
			}
		}

		i += 1
		if i >= len(args) {
			break
		}

		s = args[i]

		if len(s) < 1 || s[0] == '-' {
			break
		}
		result = s
		break
	}

	return result
}
