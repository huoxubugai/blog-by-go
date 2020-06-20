package db

import (
	"blog/config"
	"blog/models"
	"log"
)

//type即分类，因为type又是go中的关键字，因此不能直接用
func GetAllType() ([]*models.Category, error) {
	var typeList []*models.Category
	rows, err := config.DB.Raw("select id,name,blog_nums from category order by blog_nums desc").Rows()
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var category models.Category
		config.DB.ScanRows(rows, &category)
		typeList = append(typeList, &category)
	}
	return typeList, nil
}

func GetTypeById(id int) (models.Category, error) {
	var cate models.Category
	err := config.DB.Where("id=?", id).Find(&cate).Error
	if err != nil {
		return models.Category{}, err
	}
	return cate, nil
}

func InsertType(category *models.Category) (id int, err error) {
	//category:=models.Category{Id:4,Name:"测试"}
	//DB.NewRecord(category)
	//return 1,nil
	sqlstr := "insert into category(id,name) value(?,?)"
	config.DB.Exec(sqlstr, category.Id, category.Name)
	if err != nil {
		panic(err)
	}
	return
}

func GetMaxBlogNumsInType() (id int, err error) {
	//下面这句结尾加上.Error会报空指针错误，为什么呢？
	err = config.DB.Raw("select id from category order by blog_nums desc limit 1").Row().Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
