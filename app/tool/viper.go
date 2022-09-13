package tool

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

func init() {
	fmt.Println("[init] viper")
	Viper = viper.New()
	Viper.SetConfigName("config")
	Viper.AddConfigPath(".")
	Viper.SetConfigType("yaml")
	if err := Viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("CONFIG FILE NOT FOUND: %w", err))
		} else {
			panic(fmt.Errorf("CONFIG FILE WAS FOUND BUT ANOTHER ERROR WAS PRODUCED: %w", err))
		}
	}
}

func Env(name string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(name, defaultValue[0])
	}
	return Get(name)
}

func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

func Get(key string, defaultValue ...interface{}) interface{} {
	value := Viper.Get(key)
	if value == nil {
		return defaultValue[0]
	}
	return value
}

func GetString(key string, defaultValue ...interface{}) string {
	return cast.ToString(Get(key, defaultValue...))
}

func GetInt(key string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(key, defaultValue...))
}

func GetUint(key string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(key, defaultValue...))
}

func GetInt64(key string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(key, defaultValue...))
}

func GetBool(key string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(key, defaultValue...))
}
