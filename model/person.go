package model

type Person struct {
	ID   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

// used by gorm to determine the table name
func (ci *Person) TableName() string {
	return "person"
}
