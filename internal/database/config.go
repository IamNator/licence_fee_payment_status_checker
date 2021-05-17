package database

import "github.com/jinzhu/gorm"

//Config is a record of a configuration
type Config struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"  gorm:"uniqueIndex"`
	Value string `json:"value"`
}

//Add adds a new config record to database
func (c Config) Add(db *gorm.DB) error {
	return db.Model(&c).Create(c).Error
}

//Update updates the record of a configuration in a database
func (c Config) Update(db *gorm.DB) error {
	return db.Model(&c).Update(c).Error
}
