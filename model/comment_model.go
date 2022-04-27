package model

import (
	"time"
)

type Comment struct {
	Id      int64
	Mid     int64
	Uid     int64
	User    User `gorm:"foreignKey:uid"`
	Static  string
	Comment string
	Time    time.Time
	Help    int64
	Report  int64
	Star    int
}

type Review struct {
	Id      int64
	Mid     int64
	Uid     int64
	User    User `gorm:"foreignKey:uid"`
	Title   string
	Comment string
	Time    time.Time
	People  int
	Likes   int64
	Unlikes int64
	Report  int64
	Star    int
}
