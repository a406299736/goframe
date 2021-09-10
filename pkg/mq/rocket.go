package mq

import (
	"strings"
	"time"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"

	mq "github.com/aliyunmq/mq-http-go-sdk"
	"github.com/gogap/errors"
	"github.com/golang-module/carbon"
	"go.uber.org/zap"
)

// use r := New(nil)
// r.Producer(...).Push(...)
// or r.Consumer(...).pull(doFunc)

type connectConfig struct {
	endpoint, accessKey, accessSecret, token string
}

type mqConnectConfer interface {
	config() connectConfig
}

type defaultConnectConfig struct{}

type producer interface {
	Push(ctxStr, key string, properties map[string]string) (resp mq.PublishMessageResponse, err error)
}

type DoConsumeFun func(ctxStr string)

type consumer interface {
	Pull(fun DoConsumeFun)
}

func (d defaultConnectConfig) config() (c connectConfig) {
	conf := configs.Get().Rocket

	c.endpoint = conf.HttpEndpoint
	c.accessKey = conf.AccessKey
	c.accessSecret = conf.SecretKey

	return
}

type InstanceConfig2C struct {
	InstanceConfig2P
	GroupId, MsgTag string
}

type InstanceConfig2P struct {
	InstanceId, TopicId string
}

type rocketClient struct {
	client mq.MQClient
}

type product struct {
	pro mq.MQProducer
}

type consume struct {
	con mq.MQConsumer
}

func (c consume) Pull(doFunc DoConsumeFun, logger *zap.Logger) {
	for {
		endChan := make(chan int)
		respChan := make(chan mq.ConsumeMessageResponse)
		errChan := make(chan error)
		go func() {
			select {
			case resp := <-respChan:
				{
					// doFunc 处理业务逻辑
					var handles []string
					logger.Info("Consume messages---->\n", zap.Any("count", len(resp.Messages)))
					for _, v := range resp.Messages {
						handles = append(handles, v.ReceiptHandle)
						logger.Info("msg info", zap.Any("MessageID", v.MessageId), zap.Any("PublishTime", v.PublishTime),
							zap.Any("MessageTag", v.MessageTag), zap.Any("ConsumedTimes", v.ConsumedTimes),
							zap.Any("FirstConsumeTime", v.FirstConsumeTime), zap.Any("NextConsumeTime", v.NextConsumeTime),
							zap.Any("Body", v.MessageBody), zap.Any("Props", v.Properties))

						doFunc(v.MessageBody)
					}

					// NextConsumeTime前若不确认消息消费成功，则消息会被重复消费。
					// 消息句柄有时间戳，同一条消息每次消费拿到的都不一样。
					ackerr := c.con.AckMessage(handles)
					if ackerr != nil {
						// 某些消息的句柄可能超时，会导致消息消费状态确认不成功。
						if errAckItems, ok := ackerr.(errors.ErrCode).Context()["Detail"].([]mq.ErrAckItem); ok {
							for _, errAckItem := range errAckItems {
								logger.Error("Error Handle", zap.Any("ErrorHandle", errAckItem.ErrorHandle),
									zap.Any("ErrorCode", errAckItem.ErrorCode), zap.Any("ErrorMsg", errAckItem.ErrorMsg))
							}
						} else {
							logger.Error("ack err = ", zap.Any("Ack err", ackerr))
						}
						time.Sleep(time.Duration(3) * time.Second)
					} else {
						logger.Info("Ack", zap.Any("Ack -----", handles))
					}

					endChan <- 1
				}
			case err := <-errChan:
				{
					// Topic中没有消息可消费。
					if strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
						if carbon.Now().ToTimestamp()%10 == 0 {
							logger.Info("new msg", zap.Any("new msg continue", "pkg/mq/rocket.go"))
						}
					} else {
						logger.Error("err", zap.Any("err: ", err))
						time.Sleep(time.Duration(3) * time.Second)
					}
					endChan <- 1
				}
			case <-time.After(35 * time.Second):
				{
					logger.Info("timeout", zap.Any("Timeout of consumer message", "pkg/mq/rocket.go"))
					endChan <- 1
				}
			}
		}()

		// 长轮询消费消息。
		// 长轮询表示如果Topic没有消息，则客户端请求会在服务端挂起10s，10s内如果有消息可以消费则立即返回响应。
		c.con.ConsumeMessage(respChan, errChan,
			10, // 一次最多消费3条（最多可设置为16条）。
			10, // 长轮询时间3秒（最多可设置为30秒）。
		)
		<-endChan
	}
}

func (p product) Push(ctxStr, key string, properties map[string]string, logger *zap.Logger) (resp mq.PublishMessageResponse, err error) {
	var msg mq.PublishMessageRequest
	msg = mq.PublishMessageRequest{
		MessageBody: ctxStr,              //消息内容
		MessageTag:  "",                  // 消息标签
		Properties:  map[string]string{}, // 消息属性
	}
	if key != "" {
		msg.MessageKey = key
	}
	if properties != nil && len(properties) > 0 {
		for k, v := range properties {
			msg.Properties[k] = v
		}
	}

	resp, err = p.pro.PublishMessage(msg)
	logger.Info("发送内容", zap.Any("内容：", msg), zap.Any("返回ID：", resp.MessageId),
		zap.Any("err: ", err))

	return
}

// confer default nil
func New(confer mqConnectConfer) rocketClient {
	var connectConf connectConfig
	if confer == nil {
		confer = defaultConnectConfig{}
	}

	connectConf = confer.config()
	mqClient := mq.NewAliyunMQClient(connectConf.endpoint, connectConf.accessKey, connectConf.accessSecret, connectConf.token)

	return rocketClient{client: mqClient}
}

func (r rocketClient) Consumer(icf InstanceConfig2C) consume {
	return consume{con: r.client.GetConsumer(icf.InstanceId, icf.TopicId, icf.GroupId, icf.MsgTag)}
}

func (r rocketClient) Producer(icf InstanceConfig2P) product {
	return product{pro: r.client.GetProducer(icf.InstanceId, icf.TopicId)}
}
