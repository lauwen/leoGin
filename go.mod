module leoGin

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/spf13/viper v1.6.1
	luxuewen.com/config v0.0.0-00010101000000-000000000000
	luxuewen.com/message v0.0.0-00010101000000-000000000000 // indirect
	zx.ccbaidu.com/middleware v0.0.0-00010101000000-000000000000
)

replace (
	luxuewen.com/config => ./library/config
	luxuewen.com/message => ./library/message
	zx.ccbaidu.com/middleware => ./middleware
)
