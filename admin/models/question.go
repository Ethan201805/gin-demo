package models

import "gin-web-demo/Test/gin-demo/admin/global"

type Question struct {
	Name    string `form:"name" json:"name" gorm:"name"`
	Company string `form:"company" json:"company" gorm:"company"`
	Content string `form:"content" json:"content" gorm:"content"`
	Email   string `form:"email" json:"email" gorm:"email"`
}

//insert
func (q *Question) AddQuestion() (res Question, err error) {
	if err := global.ET_DB.Create(&q).Error; err != nil {
		return res, err
	}
	res = *q
	return res, nil
}
