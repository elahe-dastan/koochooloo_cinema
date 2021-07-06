package model

import (
	"time"
)

type Wallet struct {
	Username  string
	Credit    int
	SpecialAt time.Time
}
