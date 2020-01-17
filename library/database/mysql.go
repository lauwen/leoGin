package database

import (
	"github.com/gohouse/gorose/v2"
	_"github.com/go-sql-driver/mysql"
)

var err error
var engin *gorose.Engin
func init() {
	var master = &gorose.Config{
		Driver: "mysql", 
		Dsn: "root:mscc2019distributed``@tcp(172.16.116.40:3306)/go?charset=utf8&parseTime=true",
	}
	var slave0 = &gorose.Config{
		Driver: "mysql", 
		Dsn: "root:mscc2019distributed``@tcp(172.16.116.42:3306)/go?charset=utf8&parseTime=true",
	}
	var slave1 = &gorose.Config{
		Driver: "mysql", 
		Dsn: "root:mscc2019distributed``@tcp(172.16.116.43:3306)/go?charset=utf8&parseTime=true",
	}
    var config = &gorose.ConfigCluster{
		Master:       []&gorose.Config{}{master},
		Slave:        []&gorose.Config{}{slave0, slave1},
		Prefix:       "pre_",
		Driver:       "mysql",
	}
	engin, err = gorose.Open(config)
}
func DB() gorose.IOrm {
	return engin.NewOrm()
}