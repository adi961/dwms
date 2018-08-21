package main

import (
	"time"
)

const (
	Unknown      = "?"
	UpdatePeriod = 5 * time.Second
)

var Items = []statusFunc{
	netStatus("wlan0", "enp0s31f6"),
	batteryStatus("BAT0"),
	audioStatus(),
	timeStatus,
}
