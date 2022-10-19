package cream_conf

import "errors"

// IConfLoader 配置适配器接口
type IConfLoader interface {
	/*
		Initialize 初始化配置适配器
		参数:
			Configuration: 在初始化配置适配器前已经得到的配置，例如: 从环境变量中获取的配置
	*/
	LoadConf(c *Configuration) error
}

var (
	// ErrUnNeedLoad 当未满足该加载器加载条件时返回异常
	ErrUnNeedLoad = errors.New("不需要加载该配置")
)
