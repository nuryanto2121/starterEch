package util

import "time"

const (
	// WIB :
	WIB string = "Asia/Jakarta"
	// UTC :
	UTC string = "UTC"
)

// GetTimeNow :
func GetTimeNow() time.Time {
	return time.Now().In(GetLocation())
}

// GetLocation - get location wib
func GetLocation() *time.Location {
	return time.FixedZone(WIB, 7*3600)
}
