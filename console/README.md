#### 使用
1. 初始终端包： cobra init --pkg-name github.com/a406299736/goframe/console
2. 新增执行任务：cobra add MockDemo
3. 执行：go run console/main.go MockDemo 或 go build -o cmd; ./cmd MockDemo;
   执行带参数：go run console/main.go MockDemo -y true 或 ./cmd MockDemo -y true;
   或 定义方式不同,执行也不同: 
   定义1: redisMqConsumerCmd.Flags().StringP("name1", "n", "", "Help message for toggle")
   执行1: go run console/main.go redisMqConsumer -n abc
   定义2: redisMqConsumerCmd.Flags().String("env", "", "Help message for toggle")
   执行2: go run console/main.go redisMqConsumer --env rel
   特别注意: -- 两个短横线