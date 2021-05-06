package conf

import(
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

var Viper = viper.New()

func InitConfig(path string)  {
	if path == "" {
		Viper.SetConfigName("app") // 默认配置文件为app
		_, fn, _, _ := runtime.Caller(0)
		confDir := filepath.Dir(fn)
		path = filepath.Join(confDir, "../../../conf")
		Viper.AddConfigPath(path)
	} else {
		Viper.SetConfigFile(path)
	}

	err := Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal err config file: %s", err))
	}
}