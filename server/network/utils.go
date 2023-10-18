package network

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var defaultGormOpts = gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // this removes the "s" at the end of table names automatically added by GORM
		NoLowerCase:   true,
	},
}
