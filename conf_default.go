package cream_conf

import "fmt"

var defaultConf = &Configuration{}

func Initialize() error {
	fmt.Println("开始初始化【conf】基础组件")
	return defaultConf.Initialize(
		InitWithLoader(&ConfLoaderFlag{}),
		InitWithLoader(&ConfLoaderEnv{}),
	)
}

func Get(key string) interface{} {
	return defaultConf.Get(key)
}
func GetString(key string) string {
	return defaultConf.GetString(key)
}
func GetStringd(key, dv string) string {
	return defaultConf.GetStringd(key, dv)
}
func GetBool(key string) bool {
	return defaultConf.GetBool(key)
}
func GetBoold(key string, dv bool) bool {
	return defaultConf.GetBoold(key, dv)
}
func GetInt(key string) int {
	return defaultConf.GetInt(key)
}
func GetIntd(key string, dv int) int {
	return defaultConf.GetIntd(key, dv)
}
func GetInt64(key string) int64 {
	return defaultConf.GetInt64(key)
}
func GetInt64d(key string, dv int64) int64 {
	return defaultConf.GetInt64d(key, dv)
}
func GetFloat64(key string) float64 {
	return defaultConf.GetFloat64(key)
}
func GetFloat64d(key string, dv float64) float64 {
	return defaultConf.GetFloat64d(key, dv)
}
func GetStringSlice(key string) []string {
	return defaultConf.GetStringSlice(key)
}
func GetChildd(key string) *ConfigDataSourceItems {
	return defaultConf.GetChildd(key)
}
