package apollo

import (
	"errors"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/tools"
	"strings"
	"time"
)

type apollo struct {
	config *config.AppConfig
	namespace string
}

const defaultNamespace = "application"

type options func(apl *apollo)

func New(opt ...options) (*storage.Config, error) {
	apl := &apollo{}
	for _, v := range opt {
		v(apl)
	}

	if apl.config == nil {
		apl.config = defaultConfig()
	}
	if apl.namespace == "" {
		apl.namespace = defaultNamespace
	}

	start, err := Start(apl.config)
	if err != nil {
		return nil, err
	}

	conf := start.GetConfig(apl.namespace)
	if conf == nil {
		return nil, errors.New("namespace 不存在")
	}

	return conf, nil
}

func WithConfig(config *config.AppConfig) options {
	return func(apl *apollo) {
		apl.config = config
	}
}

func WithNamespace(namespace string) options {
	return func(apl *apollo) {
		apl.namespace = namespace
	}
}

func Start(apolloConfig *config.AppConfig) (*agollo.Client, error) {
	return agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return apolloConfig, nil
	})
}

// 启动检测
func CheckStart() error {
	apolloConfig := defaultConfig()
	client, err := Start(apolloConfig)
	if err != nil {
		return err
	}

	split := strings.Split(apolloConfig.NamespaceName, ",")
	for _, n := range split {
		checkKey(n,client)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func checkKey(namespace string,client *agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count:=0
	cache.Range(func(key, value interface{}) bool {
		//fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count < 1{
		panic("config key can not be null")
	}
}

func defaultConfig() *config.AppConfig {
	conf := configs.Get()
	return &config.AppConfig{AppID: conf.Apollo.AppId,
		Cluster: conf.Apollo.Cluster, NamespaceName: conf.Apollo.NamespaceName,
		IP: conf.Apollo.Ip, IsBackupConfig: true,
		BackupConfigPath: tools.GetProjectAbsolutePath() +  "/configs",
		Secret: "", SyncServerTimeout: 3}
}