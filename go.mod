module leoGin

go 1.13

require (
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet v2.1.2+incompatible // indirect
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/cmd/xorm v0.0.0-20190426080617-f87981e709a1 // indirect
	github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/gohouse/gorose/v2 v2.1.3
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.2+incompatible // indirect
	github.com/spf13/viper v1.6.1
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/xormplus/builder v0.0.0-20190724032102-0ee351fedce9 // indirect
	github.com/xormplus/core v0.0.0-20190724072625-00f5a85ad6e0 // indirect
	github.com/xormplus/xorm v0.0.0-20190926190243-42377f593eb1
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	luxuewen.com/config v0.0.0-00010101000000-000000000000
	luxuewen.com/database v0.0.0-00010101000000-000000000000
	luxuewen.com/message v0.0.0-00010101000000-000000000000 // indirect
	xorm.io/cmd/xorm v0.0.0-20191108140657-006dbf24bb9b // indirect
	zx.ccbaidu.com/middleware v0.0.0-00010101000000-000000000000
)

replace (
	luxuewen.com/config => ./library/config
	luxuewen.com/database => ./library/database
	luxuewen.com/message => ./library/message
	zx.ccbaidu.com/middleware => ./middleware
)
