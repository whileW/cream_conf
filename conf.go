package cream_conf

import "sync"

// Configuration 管理所有的配置
type Configuration struct {
	ConfDataSource
	initOnce       sync.Once     // 控制只初始化一次
	needInitLoader []IConfLoader // 需要初始化的配置加载器
}

// Initialize 初始化配置
func (c *Configuration) Initialize(opts ...func(c *Configuration)) error {
	var err error
	c.initOnce.Do(func() {
		for _, opt := range opts {
			opt(c)
		}

		// 初始化配置适配器
		for _, t := range c.needInitLoader {
			if err = t.LoadConf(c); err != nil {
				if err == ErrUnNeedLoad {
					continue
				} else {
					return
				}
			}
		}
	})
	return err
}

func InitWithLoader(a IConfLoader) func(c *Configuration) {
	return func(c *Configuration) {
		c.needInitLoader = append(c.needInitLoader, a)
	}
}
