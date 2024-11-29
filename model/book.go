package model

import "gorm.io/gorm"

type Book struct {
    gorm.Model
    Title string
    AuthorID uint
    Author   Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}