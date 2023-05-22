package entity

type Game struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
}
