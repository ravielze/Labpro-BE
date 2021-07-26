package dao

import "github.com/ravielze/oculi/common/model/dao"

type (
	Dorayaki struct {
		ID               uint64  `gorm:"primaryKey;"`
		Name             string  `gorm:"not null;type:VARCHAR(256);"`
		Description      *string `gorm:"type:VARCHAR(1024);"`
		TasteName        string  `gorm:"not null;type:VARCHAR(128);"`
		TasteDescription *string `gorm:"type:VARCHAR(1024);"`
		dao.BaseModel    `gorm:"embedded"`
	}
)
