package econfig

import (
	"github.com/spf13/viper"
)

var v *viper.Viper

func Init() {
	v = viper.New()
}

func SetConfigName(name string) {
	v.SetConfigName(name)
}

func AddConfigPath(path string) {
	v.AddConfigPath(path)
}

func Get(key string) interface{} {
	return v.Get(key)
}

func GetString(key string) string {
	return v.GetString(key)
}

func ReadInConfig() error {
	return v.ReadInConfig()
}
