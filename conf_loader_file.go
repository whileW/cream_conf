package cream_conf

import (
	"errors"
	"fmt"
	"github.com/whilew/cream_conf/common/xcodec"
	"os"
	"path"
	"path/filepath"
)

// todo 监听文件变化

type ConfLoaderFile struct {
	filePath   string
	codec      string
	sourceData []byte
	config     *Configuration
}

func (f *ConfLoaderFile) Initialize(c *Configuration) error {
	fmt.Println("-----------------------开始初始化File加载器-----------------------")
	f.config = c
	f.filePath = c.GetString("config.file.path")
	if f.filePath == "" {
		fmt.Println("未找到配置项[config.file.path]")
		fmt.Println("开始尝试搜索配置文件")
		if f.filePath = tryFindConfigFilePath(c); f.filePath == "" {
			fmt.Println("未找到配置文件，停止初始化File配置加载器")
			return ErrUnNeedLoad
		}
		fmt.Println(fmt.Sprintf("自动搜索到配置文件[%s]", f.filePath))
	}
	return f.loadConfigFile()
}

// tryFindConfigFilePath 尝试寻找配置文件
func tryFindConfigFilePath(c *Configuration) string {
	filePath := c.GetStringd("config.file.name", "config.yaml")
	for i := 0; i < 10; i++ {
		// 检查配置文件是否存在
		if _, err := os.ReadFile(filePath); err == nil {
			return filePath
		}
		filePath = "../" + filePath
	}
	return ""
}

func (f *ConfLoaderFile) loadConfigFile() error {
	absolutePath, err := filepath.Abs(f.filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("get file config absolute path error. path:%s, err: %v", f.filePath, err))
	}
	fmt.Println(fmt.Sprintf("开始加载配置文件[%s]", absolutePath))
	contentData, err := os.ReadFile(absolutePath)
	if err != nil {
		return errors.New(fmt.Sprintf("read file by absolute path error. absolute_path:%s, err: %v", absolutePath, err))
	}
	f.sourceData = contentData
	// todo 判断是否需要输出配置文件内容
	switch path.Ext(f.filePath) {
	case ".yaml":
		f.codec = xcodec.XcodecYaml
	case ".json":
		f.codec = xcodec.XcodecJson
	default:
		f.codec = xcodec.XcodecJson
	}
	conf := map[string]interface{}{}
	if err := xcodec.Decode(f.codec, contentData, &conf); err != nil {
		return errors.New(fmt.Sprintf("decode file config data failed: %v，codec: %s", err, f.codec))
	}
	f.setConfigByMap(conf)
	fmt.Println("File配置加载器加载完成")
	return nil
}

func (f *ConfLoaderFile) setConfigByMap(data map[string]interface{}) {
	var recursionSetConfig func(key string, d map[string]interface{})
	recursionSetConfig = func(key string, d map[string]interface{}) {
		for k, v := range d {
			if vv, ok := v.(map[string]interface{}); ok {
				recursionSetConfig(key+k+".", vv)
			} else {
				f.config.SetConfig(key+k, v)
			}
		}
	}
	recursionSetConfig("", data)
}
