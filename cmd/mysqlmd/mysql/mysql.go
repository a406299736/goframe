package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ Repo = (*dbRepo)(nil)

type Repo interface {
	i()
	GetDb() *gorm.DB
	DbClose() error
}

type dbRepo struct {
	DbConn *gorm.DB
}

func New(dbAddr, dbUser, dbPass, dbName, port string) (Repo, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		dbUser,
		dbPass,
		dbAddr,
		port,
		dbName,
		true,
		"Local")
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, fmt.Errorf("[db connection failed] Database name: %s", dbName)
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	return &dbRepo{
		DbConn: db,
	}, nil
}

func (d *dbRepo) i() {}

func (d *dbRepo) GetDb() *gorm.DB {
	return d.DbConn
}

func (d *dbRepo) DbClose() error {
	sqlDB, err := d.DbConn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
