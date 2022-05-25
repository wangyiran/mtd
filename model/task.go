package model

type Task struct {
	User      User
	UserId    string
	Title     string
	Status    string `gorm:"default:'0'"`
	StartTime int64
	EndTime   int64
	Content   string `gorm:"type:longtext"`
}
