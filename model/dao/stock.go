package dao

type (
	Stock struct {
		ShopID     uint64 `gorm:"primaryKey;"`
		DorayakiID uint64 `gorm:"primaryKey;"`
		Stock      uint64 `gorm:"not null;"`

		Dorayaki Dorayaki `gorm:"foreignKey:->;DorayakiID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Shop     Shop     `gorm:"foreignKey:->;ShopID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
