package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	restRestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	BASEURL string
	User struct{
		EncodedPassword string
		Password string
	}
}
