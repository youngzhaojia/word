package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	PageSize  int
	JwtSecret string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("fail to load 'conf/app.ini' : %v", err)
	}

	LoadBase()
	LoadApp()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("fail to get app config : %v", err)
	}

	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@#!@#!@#!#!")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("fail to get server config : %v", err)
	}

	HttpPort = sec.Key("HTTP_PORT").MustInt(9000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
