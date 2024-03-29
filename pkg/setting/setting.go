package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("fail to parse 'conf/app.ini' :%v", err)
	}
	LoadApp()
	LoadServer()
	LoadBase()

}
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("faild to get section 'server' :%v", err)

	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("faild to get app 'app' :%v", err)
	}

	PageSize = sec.Key("PAGE_SECRET").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")

}
