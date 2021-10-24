package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"myGinBlog/utils"
	"time"
)

var (
	db *gorm.DB
	err error
)

func InitDb()  {
	  db,err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	  	utils.DbUser,
	  	utils.DbPassWord,
	  	utils.DbHost,
	  	utils.DbPort,
	  	utils.DbName,
	  	))

	  if err != nil{
	  	fmt.Printf("连接数据库失败，请检查配置!",err)
	  }

	  //禁止自动使用复数来生成表名
	  db.SingularTable(true)
	  //自动迁移
	  db.AutoMigrate(&User{},&Article{},&Category{})

	  db.DB().SetMaxIdleConns(10)
	  db.DB().SetMaxOpenConns(100)
	  db.DB().SetConnMaxLifetime(10 * time.Second)

	  //db.Close()
}
