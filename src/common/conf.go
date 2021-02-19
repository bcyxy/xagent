package common

import (
	"path"
	"xagent/src/glbval"

	"github.com/Unknwon/goconfig"
)

var confCache *goconfig.ConfigFile

//InitConf xxx
func InitConf() error {
	var err error

	confFPath := path.Join(glbval.RootPath, "conf/conf.ini")
	confCache, err = goconfig.LoadConfigFile(confFPath)
	return err
}

//GetConfStr 获取string类型的配置
func GetConfStr(section, confKey, defVal string) string {
	confVal, err := confCache.GetValue(section, confKey)
	if err != nil {
		return defVal
	}
	return confVal
}

//GetConfInt 获取Int类型的配置
func GetConfInt(section, confKey string, defVal int) int {
	confVal, err := confCache.Int(section, confKey)
	if err != nil {
		return defVal
	}
	return confVal
}

//GetConfFloat 获取float类型的配置
func GetConfFloat(section, confKey string, defVal float64) float64 {
	confVal, err := confCache.Float64(section, confKey)
	if err != nil {
		return defVal
	}
	return confVal
}
