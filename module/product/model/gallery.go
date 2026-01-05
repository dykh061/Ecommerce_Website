package productmodel

import "time"

type Gallery struct {
	Id        int        `json:"-" gorm:"column:id;"`
	ProductId int        `json:"-" gorm:"not null;index"`
	ImageURL  string     `json:"image_url" gorm:"column:image_url;type:varchar(500);not null"`
	IsMain    bool       `json:"is_main" gorm:"not null;default:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
}

func (Gallery) TableName() string { return "galleries" }

type GalleryCreate struct {
	ProductId int    `json:"-" gorm:"not null;index"`
	ImageURL  string `json:"image_url" gorm:"column:image_url;type:varchar(500);not null"`
	IsMain    bool   `json:"is_main" gorm:"not null;default:false"`
}

func (GalleryCreate) TableName() string { return Gallery{}.TableName() }
