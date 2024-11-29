package model

import "gorm.io/gorm"

type Author struct {
    gorm.Model
    Name  string
    Email *string
    Books []Book  `gorm:"foreignKey:AuthorID"`
}