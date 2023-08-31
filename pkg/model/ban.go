package model

import (
	"net/netip"
	"time"
)

type Ban struct {
	IP      netip.Addr
	Reason  string
	EndDate time.Time
}
