package serializer

type Task struct {
	UserName  string `json:"user_name" gorm:"user_name"`
	Title     string `json:"title" gorm:"title"`
	Status    string `json:"status" gorm:"status"`
	StartTime int    `json:"start_time" gorm:"start_time"`
	EndTime   int    `json:"end_time" gorm:"end_time"`
	Content   string `json:"content" gorm:"content"`
}

type TaskList struct {
	Status string
	Count  int64 `json:"count"`
	List   []Task
}
