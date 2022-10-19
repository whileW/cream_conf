package cream_conf

import (
	"os"
	"strings"
)

type ConfLoaderEnv struct {
}

func (f *ConfLoaderEnv) LoadConf(c *Configuration) error {
	for _, t := range os.Environ() {
		kv := strings.Split(t, "=")
		if len(kv) != 2 {
			continue
		}
		c.SetConfig(kv[0], kv[1])
	}
	return nil
}
