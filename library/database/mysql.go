package database

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	_"github.com/go-sql-driver/mysql"
	"luxuewen.com/config"
)

var err error
var engin *gorose.Engin
func init() {
	conf := config.InitConfig("./config", "database", "yaml")
	if clusters := conf.GetBool("mysql.clusters"); !clusters {
		fmt.Println("mysql")
		goroseConfig := &gorose.Config{
			Driver: "mysql",
			Dsn: conf.GetStringSlice("mysql.master")[0],
		}
		engin, err = gorose.Open(goroseConfig)

	} else {
		fmt.Println("cluster")
		var masters, slaves []gorose.Config
		for _, master := range conf.GetStringSlice("mysql.master") {
			masters = append(masters, gorose.Config{
				Driver: "mysql",
				Dsn: master,
			})
		}
		
		for _, slave := range conf.GetStringSlice("mysql.slave") {
			slaves = append(slaves, gorose.Config{
				Driver: "mysql",
				Dsn: slave,
			})
		}

		goroseConfig := &gorose.ConfigCluster{
			Master: masters,
			Slave:  slaves,
			Driver: "mysql",
			Prefix: "",
		}
		engin, err = gorose.Open(goroseConfig)
	}
	if err != nil {
		panic(err.Error())
	}
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}