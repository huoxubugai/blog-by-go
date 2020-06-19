package db

import (
	"blog/config"
	"blog/models"
)

func GetAllTag() ([]models.Tag, error) {
	var tags []models.Tag
	rows, err := config.DB.Raw("select id,name,blog_nums from tag").Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tag models.Tag
		config.DB.ScanRows(rows, &tag)
		tags = append(tags, tag)
	}
	return tags, nil
}
