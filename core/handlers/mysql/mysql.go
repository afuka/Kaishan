package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"kaishan/core/handlers/conf"
)

var (
	Clients map[string]*gorm.DB
)

func InitMysql() {
	configs := conf.Viper.GetStringMap("mysql")
	Clients = make(map[string]*gorm.DB, len(configs))
	for key, _ := range configs {
		dsn := fmt.Sprintf("%s:%s@%s/%s?%s",
			conf.Viper.GetString(fmt.Sprint("mysql.", key, ".user")),
			conf.Viper.GetString(fmt.Sprint("mysql.", key, ".password")),
			conf.Viper.GetString(fmt.Sprint("mysql.", key, ".path")),
			conf.Viper.GetString(fmt.Sprint("mysql.", key, ".name")),
			conf.Viper.GetString(fmt.Sprint("mysql.", key, ".config")),
		)
		client, err := newMysqlClient(dsn)
		if err != nil {
			panic(fmt.Errorf("fatal err connect mysql: %s", key))
		}

		Clients[key] = client
	}
}

// newMysqlClient 新建一个具柄
func newMysqlClient(dsn string) (*gorm.DB, error) {
	fmt.Println(dsn)
	client, err := gorm.Open("mysql", dsn)
	if err!= nil{
		panic(err)
	}

	return client, err
}

// Close 关闭所有的mysql具柄
func Close()  {
	for _, client := range Clients {
		client.Close()
	}
}