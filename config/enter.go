package config

type Config struct {
	Mysql config.Mysql  `yaml:"mysql"`
	System System `yaml:"system"`
	Logger Logger `yaml:"logger"`
}
