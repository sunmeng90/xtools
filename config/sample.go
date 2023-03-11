package config

import "time"

type SampleConfig struct {
	Custom CustomConfig
	Db     DbConfig
}

type CustomConfig struct {
	Name    string
	Prop1   int32
	Prop2   bool
	Time1   time.Duration
	Time2   time.Duration
	Time3   time.Duration
	Time4   time.Duration
	Time5   time.Duration
	StrArr1 []string `yaml:"str_arr1"` // by default mapstructure: "str_arr1"`
	IntArr1 []int32  `yaml:"int_arr1"`
	Map1    map[string]interface{}
	Redis   RedisConfig
}

type DbConfig struct {
	Redis RedisConfig
	Mysql MysqlConfig
}

type RedisConfig struct {
	Host string
	Port int
}

type MysqlConfig struct {
	Host string
	Port int
}
