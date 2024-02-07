package models

type Car struct {
	Pemilik  string `gorm:"type:varchar(100)" json:"pemilik"`
	Merk     string `gorm:"type:varchar(100)" json:"merk"`
	Harga    int    `gorm:"type:int" json:"harga"`
	TypeCars string `gorm:"type:varchar(100)" json:"type_cars"`
	Id       uint   `gorm:"primaryKey" json:"id"`
}

// func (p *Car) BeforeCreate(tx *gorm.DB) (err error) {
// 	fmt.Println("Car Before Create()")

// 	if len(p.Pemilik) < 4 {
// 		err = errors.New("product name is too short")
// 	}
// 	return
// }
