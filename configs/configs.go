package configs

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/golang-module/carbon"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	MySQL struct {
		Read struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"read"`
		Write struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"write"`
		// 连接conn1读库
		Conn1read struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"conn1read"`
		// 连接conn1写库
		Conn1write struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"conn1write"`
		Base struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
		} `toml:"base"`
	} `toml:"mysql"`

	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"redis"`

	JWT struct {
		Secret         string        `toml:"secret"`
		ExpireDuration time.Duration `toml:"expireDuration"`
	} `toml:"jwt"`

	URLToken struct {
		Secret         string        `toml:"secret"`
		ExpireDuration time.Duration `toml:"expireDuration"`
	} `toml:"urlToken"`

	HashIds struct {
		Secret string `toml:"secret"`
		Length int    `toml:"length"`
	} `toml:"hashids"`

	Language struct {
		Local string `toml:"local"`
	} `toml:"language"`

	Rocket struct {
		HttpEndpoint string `toml:"httpEndpoint"`
		AccessKey    string `toml:"accessKey"`
		SecretKey    string `toml:"secretKey"`
		InstanceId   string `toml:"instanceId"`
		Topic        string `toml:"topic"`
		GroupId      string `toml:"groupId"`
	} `toml:"rocket"`

	App struct {
		LogPath string `toml:"logPath"`
		Env     string `toml:"env"`
		Debug   bool   `toml:"debug"`
	} `toml:"app"`

	Apollo struct {
		AppId         string `toml:"appId"`
		Cluster       string `toml:"cluster"`
		NamespaceName string `toml:"namespaceName"`
		Ip            string `toml:"ip"`
	}
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}

func IsPro() bool {
	return config.App.Env == "pro"
}

func (c Config) LogPath() string {
	date := carbon.Now()
	return c.App.LogPath + ProjectName + date.ToDateString() + ".log"
}
