package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//  jdbc:postgresql://192.168.103.113:5432/zhjg?currentSchema=original
func main() {
	//1,打开db实例
	db, err := gorm.Open("postgres",
		"host=192.168.103.113 user=cic_database dbname=zhjg port=5432 sslmode=disable password=cic_database")

	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	//2，设置连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//3，打开日志
	db.LogMode(true)
	// 由于需要返回多条数据所以需要使用切片去接收，如果有且仅有一条数据可以不使用切片
	//list := []Product{}
	//调用原生sql语句
	//db.Raw("SELECT id, name FROM original.test_tang WHERE name=?", "HAT").Find(&list)
	// 打印结果
	//fmt.Println(list)


	var result Product
	//var result int
	////调用原生sql语句
	//db.Raw("SELECT id, name FROM original.test_tang WHERE name=?", "HAT").Scan(&result)
	//// 当在执行inset的时候可以建一个空的结构体
	//db.Raw("delete from user_infos where name=?", "连少").Scan(&result)
	db =db.Raw("insert into original.test_tang (id,name) values (12345,'男')").Scan(&result)



	//db = db.Raw("update original.test_tang set name='鹏少123' where id=123").Scan(&result)
	//db =db.Exec("update original.test_tang set name='鹏少123678' where id=123")
	fmt.Println(db.Error)

}
type Product struct {
	Id int
	Name string
}