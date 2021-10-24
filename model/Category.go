package model

import (
	"github.com/jinzhu/gorm"
	"myGinBlog/utils/errmsg"
)

type Category struct {
	ID  uint 	`gorm:"primary_key;auto_increment" json:"id"` 
	Name string `gorm:"type:varchar(20);not null" json:"name"`

}

//添加分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 检测分类是否存在
func CheckCategory(name string) (code int)  {
	var category Category
	db.Select("id").Where("name = ?" ,name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return  errmsg.SUCCESS
}
//检查除了自己外是否还有同名的分类
func CheckCategoryExceptSelf(id uint ,name string)(code int)  {
	var category Category
	db.Select("id").Where("name = ?",name).First(&category)
	if category.ID > 0 && category.ID != id{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}
// 通过ID 检查是否存在记录
func CheckCategoryById(id int) (code int)  {
	var category Category
	db.Where("id = ?",id).First(&category)
	if category.ID <=0 {
		return errmsg.ERROR_USER_NOT_EXISTS
	}
	return errmsg.SUCCESS
}

// 查询所有分类
func GetCategories(pageSize int ,pageNum int)[]Category  {
	var categories [] Category
	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categories
}

// 删除分类
func DeleteCategory(id int) (code int)  {
	var category Category
	err := db.Where("id = ?",id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return  errmsg.SUCCESS
}

//编辑分类
func EditCategory( id int,data *Category) int  {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&category).Where("id = ?",id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
