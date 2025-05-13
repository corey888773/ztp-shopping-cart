package products

type Product struct {
	ID   string `json:"product_id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}

type ProductReservation struct {
	ProductID      string `json:"product_id" gorm:"primaryKey;column:product_id"`
	SequenceNumber int    `json:"sequence_number" gorm:"primaryKey;column:sequence_number;autoIncrement:false"`
	LockedToTime   string `json:"locked_to_time" gorm:"column:locked_to_time"`
	CartID         string `json:"cart_id" gorm:"column:cart_id"`
}
