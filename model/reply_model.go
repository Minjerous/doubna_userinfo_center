package model

import "time"

type Reply struct {
	Id       int64
	Uid      int64
	Pid      int64
	Like     int
	Ptable   string
	Date     time.Time
	Username string
	Content  string
	Avatar   string
	RepCnt   int
}
