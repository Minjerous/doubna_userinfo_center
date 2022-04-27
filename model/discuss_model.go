package model

import "time"

type Discussion struct {
	Id        int64
	User      User `gorm:"foreignKey:uid"`
	Uid       int64
	Mid       int64
	UserName  string
	MovieName string
	Avatar    string
	Title     string
	Value     string
	ReplyNum  int64
	Date      time.Time
	Stars     int64
}
type DiscussionList struct {
	Id        int64
	Uid       int64
	Mid       int64
	User      User `gorm:"foreignKey:uid"`
	UserName  string
	MovieName string
	Title     string
	ReplyNum  int64
	Date      time.Time
}
