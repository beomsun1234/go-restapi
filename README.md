# go-restapi
📚 Go RestAPI fiber, gorm


## postgres read, write 설정


    //기본적은 write 역활
	gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	

    //read 역활, Register에 Sources로 등록시 write 역활을 한다.
	p.PostgresDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{
			postgres.Open(dsn2),
		},
		Policy: dbresolver.RandomPolicy{},
	}).Register(dbresolver.Config{
		Sources: []gorm.Dialector{
			postgres.Open(dsn3),
		},
	}))
	