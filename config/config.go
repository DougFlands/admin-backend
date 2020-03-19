package config

import (
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/spf13/viper"
)

type Config struct {
	Kafka struct{
		Topic string
		Host string
	}
	Consul struct{
		Addr string
	}
	Oss struct{
		Path string
		AccessKeyId string
		AccessKeySecret string
	}
	Mysql struct{
		Host string
	}
	Redis struct{
		Host string
		Passwd string
	}
	Service struct{
		ApiGatewayHost string
		UploadHost string
	}
	PwdSalt string

}

var (
	// Kafka 配置
	KafkaTopic string
	KafkaHosts string

	// Consul 配置
	ConsulAddr string
	ConsulReg registry.Registry

	// Oss 配置
	OSSPath            string
	OSSAccessKeyID     string
	OSSAccessKeySecret string

	// mysql
	MySqlHost string

	// redis
	RedisHost   string
	RedisPasswd string

	// User 配置
	PwdSalt string

	// Service 配置
	// 网关服务地址
	ApiGatewayServiceHost string
	// 上传服务地址
	UploadServiceHost string

	//文件存放路径
	FilePath = "./data/file/"
)

func init() {
	var config Config
	viper.SetConfigName("config")   // 设置配置文件名 (不带后缀)
	viper.AddConfigPath("./data/config")        // 第一个搜索路径
	err := viper.ReadInConfig()     // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&config)        // 将配置信息绑定到结构体上
	KafkaTopic = config.Kafka.Topic
	KafkaHosts = config.Kafka.Host
	ConsulAddr = config.Consul.Addr
	OSSPath = config.Oss.Path
	OSSAccessKeyID     = config.Oss.AccessKeyId
	OSSAccessKeySecret = config.Oss.AccessKeySecret
	MySqlHost = config.Mysql.Host
	RedisHost   = config.Redis.Host
	RedisPasswd = config.Redis.Passwd
	ApiGatewayServiceHost = config.Service.ApiGatewayHost
	UploadServiceHost = config.Service.UploadHost
	PwdSalt = config.PwdSalt

	ConsulReg = consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			ConsulAddr,
		}
	})
}