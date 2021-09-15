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
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/api/service/demo"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/cache"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/core"
	db2 "gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/db"
	"go.uber.org/zap"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/logger"

	"github.com/spf13/cobra"
)

// MockDemoCmd represents the MockDemo command
var MockDemoCmd = &cobra.Command{
	Use:   "MockDemo",
	Short: "A brief description of your command",
	Long:  `A longer description ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MockDemo called")
		// TODO something ...

		// log应作为参数传递，多次NewJSONLogger会造成trace-id变化
		log, _ := logger.NewJSONLogger(logger.WithFileP(configs.Get().LogPath()), logger.WithTrace())
		log.Info("MockDemo called")

		is, _ := cmd.Flags().GetBool("isYou")
		fmt.Println(is)
		if is {
			fmt.Println("Is you")
			initMock(log, args)
		} else {
			log.Info("Not you")
			fmt.Println("Not you")
		}
	},
}

func initMock(logger *zap.Logger, args []string) {
	// TODO something
	fmt.Println("initMock called")

	db, err := db2.New()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	redis, err := cache.New()
	if err != nil {
		fmt.Println("%v", err)
		return
	}

	srv := demo.NewDemoService2(db, redis, core.NewCmdContext(logger))
	info, e := srv.Create()
	if e != nil {
		fmt.Printf("%v", e)
		return
	}
	fmt.Printf("%v", info)
}

func init() {
	rootCmd.AddCommand(MockDemoCmd)

	MockDemoCmd.Flags().BoolP("isYou", "y", false, "test rootCmd.Flags()...")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MockDemoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MockDemoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
