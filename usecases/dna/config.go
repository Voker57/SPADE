package dna

import "time"

// gRPC configurations
const (
	MegaByte = 1024 * 1024
	TimeOut  = 5 * time.Second
	// MaxMsgSize in MB
	MaxMsgSize = 10 * MegaByte
)

// Database Configuration
const (
	DbName = "dna_database.sqlite"
	TbName = "users_cipher"
)

// SPADE configuration
const (
	// NumUsers for now let's just assume that each user is
	// going to encrypt a single file, so NumUsers is going
	// to be the number of dna sequence files
	NumUsers = 5012
	// MaxVecSize should be defined to a fixed number, where
	// the system is going to work with it by running spade's
	// setup phase. Actually the length of mpk and msk is set
	// to this value
	MaxVecSize = 78214
	// PaddingItem basically it must be anything larger than
	// the maximum value of biggest element of the vector of
	// dna dinucleotides, in this case for us is 16, so any x > 16 works
	PaddingItem = 22
)
