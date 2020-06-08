package appconf

import (
	"github.com/appconf/engine/render"
)

//Storage 后端存储的相关配置
type Storage struct {
	//Type 存储类型
	Type string `toml:"type"`
	//Config 在根据Type创建Storage时，会将Config传递Storage的构造函数
	Config map[string]interface{} `toml:"config"`
}

//Log 定义Log相关的配置
type Log struct {
	Level string `toml:"level"`
}

//Config 配置文件内容
type Config struct {
	Storage   Storage          `toml:"storage"`
	Templates []*render.Config `toml:"templates"`
	Log       Log              `toml:"log"`
}
