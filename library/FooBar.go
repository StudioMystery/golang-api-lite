package library

import "gorm.io/gorm"

type FooBar struct {
	gorm.Model
	Foo uint
	Bar string
}
