package products

type Product struct {
	ID           string `json:"product_id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"column:name"`
	LockedToTime string `json:"locked_to_time" gorm:"column:locked_to_time"`
}
