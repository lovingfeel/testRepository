package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Name string
	Id int

}

func main() {
	////创建记录
	//users := []User{{Name: "Jinzhu", Id: 18},{Name: "Tanglw", Id: 19},{Name: "denglijun", Id: 20}}
	////1,打开db实例
	//db, err := gorm.Open("postgres",
	//	"host=192.168.103.113 user=cic_database dbname=zhjg port=5432 sslmode=disable password=cic_database")
	//
	//if err != nil {
	//	panic("连接数据库失败")
	//}
	//defer db.Close()
	//result := db.Table("original.test_tang").Create(&users) // 通过数据的指针来创建
	//
	////user.ID             // 返回插入数据的主键
	////result.Error        // 返回 error
	////result.RowsAffected // 返回插入记录的条数
	//fmt.Println("返回的记录数:",result.RowsAffected)

	dsn := "host=192.168.103.113 user=cic_database password=cic_database dbname=zhjg port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败")
	}


	users := []User{{Name: "Jinzhu", Id: 17},{Name: "Tanglw", Id: 19},{Name: "denglijun", Id: 20}}

	//result := db.Table("original.test_tang").Create(&users) // 通过数据的指针来创建
	result := db.Table("original.test_tang").CreateInBatches(&users,2) // 通过数据的指针来创建

	fmt.Println("返回的记录数:",result.Error)
	// https://github.com/go-gorm/postgres
	//db, err := gorm.Open(postgres.New(postgres.Config{
	//	DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	//	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	//}), &gorm.Config{})

}


func batchInsert()  {
	
}

func singeInsert()  {
	dsn := "host=192.168.103.113 user=cic_database password=cic_database dbname=zhjg port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败")
	}

	user := User{Name: "Jinzhu", Id: 18}
	result := db.Table("original.test_tang").Create(&user) // 通过数据的指针来创建
	fmt.Println("返回的记录数:",result.RowsAffected)
	// https://github.com/go-gorm/postgres
	//db, err := gorm.Open(postgres.New(postgres.Config{
	//	DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	//	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	//}), &gorm.Config{})
}