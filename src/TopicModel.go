package src

import "time"

type Topic struct {
	TopicID         int       `json:"id" gorm:"primary_key"`
	TopicTitle      string    `json:"title" binding:"min=4,max=20" gorm:"not null"`
	TopicShortTitle string    `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP          string    `json:"ip" binding:"ipv4" gorm:"not null"`
	TopicScore      int       `json:"score" binding:"omitempty,gt=5"`
	TopicUrl        string    `json:"url" binding:"omitempty,topicurl" gorm:"not null"`
	TopicDate       time.Time `json:"url" binding:"required" gorm:"not null"`
}

type Topics struct {
	TopicList     []Topic `json:"topics" binding:"gt=0,lt=3,topics,dive"`
	TopicListSize int     `json:"size" binding:""`
}

func CreateTopic(id int, title string) Topic {
	return Topic{
		TopicID:    id,
		TopicTitle: "title",
	}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}

type TopicClass struct {
	ClassId     int `gorm:"primary_key"`
	ClassName   string
	ClassRemark string
	ClassType   string `gorm:"column:classtype"`
}
