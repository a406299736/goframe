package configs

import (
	"time"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/env"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/tools"

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
	} `toml:"app"`

	Center struct{
		ClassUrl string `toml:"classUrl"`
	} `toml:"center"`
}

func init() {
	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath(tools.GetProjectAbsolutePath() + "/configs")

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

func (c Config) LogPath() string {
	date := carbon.Now()
	return c.App.LogPath + ProjectName + date.ToDateString() + ".log"
}
