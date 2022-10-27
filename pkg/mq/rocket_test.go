package mq

import (
	"testing"

	"github.com/a406299736/goframe/configs"
	"github.com/a406299736/goframe/pkg/logger"
)

func TestPush(t *testing.T) {
	log, _ := logger.NewJSONLogger(logger.WithFileP(configs.Get().LogPath()), logger.WithTrace())

	rocket := New(nil)
	conf := configs.Get().Rocket
	t.Log(rocket, log, conf)

	//producer
	//rocket.Producer(conf.InstanceId, conf.Topic).Push("producer test", "", nil, log)
}
