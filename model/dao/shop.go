package dao

import "github.com/ravielze/oculi/common/model/dao"

type (
	Shop struct {
		ID          uint64  `gorm:"primaryKey;"`
		Name        string  `gorm:"not null;type:VARCHAR(256);"`
		Description *string `gorm:"type:VARCHAR(1024);"`

		AddressStreet   string `gorm:"not null;type:VARCHAR(256);"`
		AddressDistrict string `gorm:"not null;type:VARCHAR(128);"`
		AddressCity     string `gorm:"not null;type:VARCHAR(128);"`
		AddressProvince string `gorm:"not null;type:VARCHAR(128);"`
		dao.BaseModel   `gorm:"embedded"`
	}
)
