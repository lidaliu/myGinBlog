package model

import (
	"github.com/jinzhu/gorm"
	"myGinBlog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title string  `gorm:"type:varchar(100);not null" json: "title"`
	Cid int`gorm:"type:int;not null" json: "cid"`
	Desc string `gorm:"type:varchar(200);" json: "desc"`
	Content string `gorm:"type:longtext" json: "content"`
	Img string `gorm:"type:varchar(100);" json: "img"`
}

//添加文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 检测文章是否存在
func CheckArticle(name string) (code int)  {
	var article Article
	db.Select("id").Where("name = ?" ,name).First(&article)
	if article.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return  errmsg.SUCCESS
}
// 通过ID 检查是否存在记录
func CheckArticleById(id int) (code int)  {
	var article Article
	db.Where("id = ?",id).First(&article)
	if article.ID <=0 {
		return errmsg.ERROR_USER_NOT_EXISTS
	}
	return errmsg.SUCCESS
}

// 查询所有文章
func GetArticles(pageSize int ,pageNum int)([]Article, int)  {
	var articleList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,errmsg.ERROR
	}
	return articleList,errmsg.SUCCESS
}
//查询单个文章信息
func GetArticleInfo( id int) (Article,int) {
	var art Article
	err := db.Preload("Category").Where("id = ?",id).First(&art).Error
	if err != nil {
		return art,errmsg.ERROR_ARTICLE_NOT_EXISTS
	}
	return art,errmsg.SUCCESS
}

// 查询某个分类下面的所有文章
func GetArticleByCategory(pageSize int,pageNum int,id int)([]Article,int)  {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?",id).Find(&cateArtList).Error
	if err != nil {
		return cateArtList,errmsg.ERROR_CATEGORY_NOT_EXISTS
	}
	return cateArtList,errmsg.SUCCESS
}


// 删除分类
func DeleteArticle(id int) (code int)  {
	var article Article
	err := db.Where("id = ?",id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return  errmsg.SUCCESS
}

//编辑分类
func EditArticle( id int,data *Article) int  {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&article).Where("id = ?",id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
