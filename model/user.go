package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID   int64  `gorm: "column: user_id; type: bigint; primary_key; unique; not null"`
	Role     string `gorm: "column: role; type: nvarchar(50); not null;"`
	Email    string `gorm: "column: email; type: nvarchar(50); unique; not null"`
	Username string `gorm: "column: username; type: nvarchar(50); unique; not null"`
	Password string `gorm: "column: password; type: tinytext; not null"`
}
