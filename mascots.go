package main

type Mascots struct {
	ID          uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Language    string `gorm:"column:language;not null"`
	Name        string `gorm:"column:name;not null"`
	Description string `gorm:"column:description"`
}

func (Mascots) TableName() string {
	return "mascots"
}
