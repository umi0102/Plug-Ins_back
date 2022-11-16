package redis

import "time"

type Redis struct {
	Host        string
	Password    string
	DB          int
	IdleTimeout time.Duration
}
