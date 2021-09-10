#### 使用
1. 初始终端包： cobra init --pkg-name gitlab.weimiaocaishang.com/weimiao/go-basic/console
2. 新增执行任务：cobra add MockDemo
3. 执行：go run console/main.go MockDemo 或 go build -o cmd; ./cmd MockDemo;
   执行带参数：go run console/main.go MockDemo -y true 或 ./cmd MockDemo -y true;