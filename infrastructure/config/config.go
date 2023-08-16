package config

import (
	"github.com/spf13/viper"
)

var (
	// viper 설정파일 로더
	v *viper.Viper
)

// Load 설정파일을 로드한다.
func Load(cfgPath string) error {
	v = viper.New()
	v.SetConfigFile(cfgPath)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return nil
}
