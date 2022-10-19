package cream_conf

import (
	"os"
	"strings"
)

type ConfLoaderFlag struct {
}

func (f *ConfLoaderFlag) LoadConf(c *Configuration) error {
	for _, t := range os.Args[1:] {
		kv := strings.Split(t, "=")
		if len(kv) != 2 {
			continue
		}
		c.SetConfig(kv[0], kv[1])
	}
	return nil
}
