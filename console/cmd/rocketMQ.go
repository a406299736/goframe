/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/logger"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/mq"

	"github.com/spf13/cobra"
)

// rocketMQCmd represents the rocketMQ command
var rocketMQCmd = &cobra.Command{
	Use:   "rocketMQ",
	Short: "A brief description of your command",
	Long:  `A longer description that ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rocketMQ called")

		rocket := mq.New(nil)
		conf := configs.Get().Rocket

		log, _ := logger.NewJSONLogger(logger.WithFileP(configs.Get().LogPath()), logger.WithTrace())

		// consumer
		rocket.Consumer(mq.InstanceConfig2C{GroupId: conf.GroupId,
			InstanceConfig2P: mq.InstanceConfig2P{InstanceId: conf.InstanceId,
				TopicId: conf.Topic}}).Pull(func(ctxStr string) {
			fmt.Println("pull msg:" + ctxStr)
		}, log)

		fmt.Println("rocketMQ called end")

	},
}

func init() {
	rootCmd.AddCommand(rocketMQCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rocketMQCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rocketMQCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
