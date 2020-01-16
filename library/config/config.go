package config

import (
	"github.com/spf13/viper"
)

/*
	path	配置文件所在目录路径
	fileName	配置文件名字
	fileType	配置文件类型
	return config viger对象
*/
func InitConfig ( path string, fileName string, fileType string) (config *viper.Viper) {
	config = viper.New()
	config.AddConfigPath(path)
	config.SetConfigName(fileName)
	config.SetConfigType(fileType)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	} 
	return config
}