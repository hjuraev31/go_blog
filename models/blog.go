package models

type Blog struct {
	Id          uint   `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Image       string `json: "image"`
	UserId      string `json: "userid"`
	User        User   `json: "user";gorm:"foreignKey:UserId`
}
