package cream_conf

import (
	"github.com/whilew/cream_conf/common/xcast"
	"strings"
)

type (
	// ConfDataSource 配置数据源
	ConfDataSource struct {
		items *ConfigDataSourceItems
	}
	ConfigDataSourceItems []*ConfDataSourceItem
	// ConfDataSourceItem 配置数据源项
	ConfDataSourceItem struct {
		key    string
		value  interface{}
		childs ConfigDataSourceItems
	}
)

func (ds *ConfigDataSourceItems) SetConfig(key []string, value interface{}) {
	var find_key = func(s *ConfDataSourceItem) {
		if len(key) > 1 {
			if s.childs == nil {
				s.childs = ConfigDataSourceItems{}
			}
			s.childs.SetConfig(key[1:], value)
		} else {
			s.value = value
		}
	}
	for _, t := range *ds {
		if t.key == key[0] {
			find_key(t)
			return
		}
	}
	s := &ConfDataSourceItem{
		key: key[0],
	}
	find_key(s)
	*ds = append(*ds, s)
}
func (ds *ConfigDataSourceItems) find(keys []string) interface{} {
	for _, t := range *ds {
		if t.key == keys[0] {
			if len(keys) > 1 {
				return t.childs.find(keys[1:])
			} else {
				return t.value
			}
		}
	}
	return nil
}
func (s *ConfigDataSourceItems) Get(key string) interface{} {
	keys := strings.Split(strings.ToLower(key), ".")
	return s.find(keys)
}
func (s *ConfigDataSourceItems) GetString(key string) string {
	return xcast.ToString(s.Get(key))
}
func (s *ConfigDataSourceItems) GetStringd(key, dv string) string {
	v := s.GetString(key)
	if v == "" {
		return dv
	}
	return v
}
func (s *ConfigDataSourceItems) GetBool(key string) bool {
	return xcast.ToBool(s.Get(key))
}
func (s *ConfigDataSourceItems) GetBoold(key string, dv bool) bool {
	if v := s.Get(key); v == nil {
		return dv
	} else {
		return xcast.ToBool(dv)
	}
}
func (s *ConfigDataSourceItems) GetInt(key string) int {
	return xcast.ToInt(s.Get(key))
}
func (s *ConfigDataSourceItems) GetIntd(key string, dv int) int {
	if v := s.Get(key); v == nil {
		return dv
	} else {
		return xcast.ToInt(v)
	}
}
func (s *ConfigDataSourceItems) GetInt64(key string) int64 {
	return xcast.ToInt64(s.Get(key))
}
func (s *ConfigDataSourceItems) GetInt64d(key string, dv int64) int64 {
	if v := s.Get(key); v == nil {
		return dv
	} else {
		return xcast.ToInt64(v)
	}
}
func (s *ConfigDataSourceItems) GetFloat64(key string) float64 {
	return xcast.ToFloat64(s.Get(key))
}
func (s *ConfigDataSourceItems) GetFloat64d(key string, dv float64) float64 {
	if v := s.Get(key); v == nil {
		return dv
	} else {
		return xcast.ToFloat64(v)
	}
}
func (s *ConfigDataSourceItems) GetStringSlice(key string) []string {
	return xcast.ToStringSlice(key)
}
func (s *ConfigDataSourceItems) GetChildd(key string) *ConfigDataSourceItems {
	if v := s.Get(key); v == nil {
		return &ConfigDataSourceItems{}
	} else {
		if v, ok := v.(*ConfigDataSourceItems); ok {
			return v
		} else {
			return &ConfigDataSourceItems{}
		}
	}
}

func (s *ConfDataSource) SetConfig(key string, value interface{}) {
	s.items.SetConfig(splitConfigKey(key), value)
}
func (s *ConfDataSource) find(keys []string) interface{} {
	return s.items.find(keys)
}
func (s *ConfDataSource) Get(key string) interface{} {
	keys := strings.Split(strings.ToLower(key), ".")
	return s.find(keys)
}
func (s *ConfDataSource) GetString(key string) string {
	return xcast.ToString(s.Get(key))
}
func (s *ConfDataSource) GetStringd(key, dv string) string {
	v := s.GetString(key)
	if v == "" {
		return dv
	}
	return v
}
func (s *ConfDataSource) GetBool(key string) bool {
	return xcast.ToBool(s.Get(key))
}
func (s *ConfDataSource) GetBoold(key string, dv bool) bool {
	if v := s.Get(key); v == nil {
		return dv
	} else {
		return xcast.ToBool(dv)
	}
}
func (s *ConfDataSource) GetInt(key string) int {
	return xcast.ToInt(s.Get(key))
}
func (s *ConfDataSource) GetIntd(key string, dv int) int {
	if v := s.Get(key); v == nil {
		return dv
	} else {
		return xcast.ToInt(v)
	}
}
func (s *ConfDataSource) GetInt64(key string) int64 {
	return xcast.ToInt64(s.Get(key))
}
func (s *ConfDataSource) GetInt64d(key string, dv int64) int64 {
	if v := s.Get(key); v == nil {
		return dv
	} else {
		return xcast.ToInt64(v)
	}
}
func (s *ConfDataSource) GetFloat64(key string) float64 {
	return xcast.ToFloat64(s.Get(key))
}
func (s *ConfDataSource) GetFloat64d(key string, dv float64) float64 {
	if v := s.Get(key); v == nil {
		return dv
	} else {
		return xcast.ToFloat64(v)
	}
}
func (s *ConfDataSource) GetStringSlice(key string) []string {
	return xcast.ToStringSlice(key)
}
func (s *ConfDataSource) GetChildd(key string) *ConfigDataSourceItems {
	if v := s.Get(key); v == nil {
		return &ConfigDataSourceItems{}
	} else {
		if v, ok := v.(*ConfigDataSourceItems); ok {
			return v
		} else {
			return &ConfigDataSourceItems{}
		}
	}
}
