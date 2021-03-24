package flag

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-ini/ini"
)

func tryParseIni(filename string, values map[string]string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	file, err := ini.Load(data)

	if err != nil {
		return err
	}

	for _, sec := range file.Sections() {
		sname := sec.Name()
		if sname == "__flag__" {
			sname = ""
		}
		for _, key := range sec.Keys() {
			kname := key.Name()
			name := sname + "." + kname
			values[name] = key.Value()
		}
	}

	return nil
}

func dumpIniFlag(cfg map[string]interface{}) {
	file := ini.Empty()
	for k, v := range cfg {
		names := strings.SplitN(k, ".", 2)
		var s *ini.Section
		var err error
		kname := ""
		if len(names) < 2 {
			s, err = file.GetSection("__flag__")
			if err != nil {
				s, err = file.NewSection("__flag__")
			}
			kname = k
		} else {
			s, err = file.GetSection(names[0])
			if err != nil {
				s, err = file.NewSection(names[0])
			}

			kname = names[1]
		}

		s.NewKey(kname, v.(Value).String())
	}

	file.WriteTo(os.Stdout)
}
