package db

import "github.com/a406299736/goframe/configs"

// 创建conn1连接
// TODO 未验证能否正确连接
func NewConn1() (Repo, error) {
	cfg := configs.Get().MySQL
	dbr, err := dbConnect(cfg.Conn1read.User, cfg.Conn1read.Pass, cfg.Conn1read.Addr, cfg.Conn1read.Name)
	if err != nil {
		return nil, err
	}

	dbw, err := dbConnect(cfg.Conn1write.User, cfg.Conn1write.Pass, cfg.Conn1write.Addr, cfg.Conn1write.Name)
	if err != nil {
		return nil, err
	}

	return &dbRepo{
		DbR: dbr,
		DbW: dbw,
	}, nil
}
