package socks5

import (
	"github.com/txthinking/x"
)

// Debug enable debug log
var Debug bool = true
var Dial x.Dialer = x.DefaultDial

func init() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
}
