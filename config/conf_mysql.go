package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
	Loglevel string `yaml:"log_level"`
	Config   string `yaml:"config"` //高级配置,如charset
}

func (msql *Mysql) Dsn() string {
	return msql.UserName + ":" + msql.PassWord + "@tcp(" + msql.Host + ":" + strconv.Itoa(msql.Port) + ")/" + msql.DB + "?" + msql.Config
}
