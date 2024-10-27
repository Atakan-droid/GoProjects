package dal

type Todo struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null max:255"`
	Description string `json:"description" gorm:"default:null max:1024"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}
