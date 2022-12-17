package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type Postgres struct {
	PostgresDB *gorm.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connection() error {
	dsn := "host=localhost user=postgres password=1234 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Seoul"
	dsn2 := "host=localhost user=postgres password=1234 dbname=gorm2 port=5432 sslmode=disable TimeZone=Asia/Seoul"
	//dsn3 := "host=localhost user=root password=1234 dbname=gorm3 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	p.PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	p.PostgresDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{
			postgres.Open(dsn2),
		},
		Policy: dbresolver.RandomPolicy{},
	}))
	if err != nil {
		return err
	}

	return nil
}
