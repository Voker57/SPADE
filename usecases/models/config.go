package models

import "time"

type Config struct {
	DbName      string
	TbName      string
	NumUsers    int
	MaxVecSize  int
	PaddingItem int
	TimeOut     time.Duration
	MaxMsgSize  int
}

func NewConfig(dbName, tbName string, numUsers, maxVecSize, padItem int, tot time.Duration, mms int) *Config {
	return &Config{
		DbName:      dbName,
		TbName:      tbName,
		NumUsers:    numUsers,
		MaxVecSize:  maxVecSize,
		PaddingItem: padItem,
		TimeOut:     tot,
		MaxMsgSize:  mms,
	}
}
