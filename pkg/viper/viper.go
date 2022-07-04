package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("migrate-data.toml")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// Conf DBManager配置
type Conf map[string]GroupConf

// GroupConf DB主从配置
type GroupConf struct {
	Master MasterConf
	Slave  SlaveConf
}

// MasterConf Master配置
type MasterConf struct {
	Dialect         string // 类型
	DSN             string // 地址
	MaxOpenConns    int    // 最大连接数
	MaxIdleConns    int    // 最大闲置连接数
	ConnMaxLifetime int    // 最大空闲回收时间
	LogMode         bool   // 是否开启debug日志
	LogSQL          bool   // 是否显示日志中的sql
	SlowThreshold   int    // 慢日志阈值
}

// SlaveConf Slave配置
type SlaveConf struct {
	Dialect         string   // 类型
	DSN             []string // 地址
	MaxOpenConns    int      // 最大连接数
	MaxIdleConns    int      // 最大闲置连接数
	ConnMaxLifetime int      // 最大空闲回收时间
	LogMode         bool     // 是否开启debug日志
	LogSQL          bool     // 是否显示日志中的sql
	SlowThreshold   int      // 慢日志阈值
}

// clickhouse库配置
func CHConf(name string) (conf *GroupConf, err error) {
	conf = &GroupConf{}
	err = viper.UnmarshalKey(fmt.Sprintf("clickhouse.%s", name), conf)
	return
}

func main() {
	conf, err := CHConf("udcp_log_ch")
	if err != nil {
		return
	}
	fmt.Println(conf.Slave.DSN)
}
