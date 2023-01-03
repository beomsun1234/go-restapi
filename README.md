# go-restapi
📚 Go RestAPI fiber, gorm, redis, test

## Test


### Test Coverage

![image](https://user-images.githubusercontent.com/68090443/209333062-d1b5650a-bf54-4524-b319-ac817c58b3c2.png)

### Test

![image](https://user-images.githubusercontent.com/68090443/209333170-058a7f81-d3f4-4a2c-9596-86efec52a24d.png)


## Gorm, postgres read, write 설정


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
	
