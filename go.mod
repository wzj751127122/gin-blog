module gin-blog

go 1.16

replace (
	github.com/wzj751127122/gin-blog/conf => ./gin-blog/pkg/conf
	github.com/wzj751127122/gin-blog/middleware => ./gin-blog/middleware
	github.com/wzj751127122/gin-blog/models => ./gin-blog/models
	github.com/wzj751127122/gin-blog/pkg/e => ./gin-blog/pkg/e
	github.com/wzj751127122/gin-blog/pkg/setting => ./gin-blog/pkg/setting
	github.com/wzj751127122/gin-blog/pkg/util => ./gin-blog/pkg/util
	github.com/wzj751127122/gin-blog/routers => ./gin-blog/routers
	github.com/wzj751127122/gin-blog/routers/api/V1 => ./gin-blog/routers/api/V1
)

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.2
	github.com/jinzhu/gorm v1.9.16
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/unknwon/com v1.0.1
)
