package mp4

import "time"

const (
	NTPEpochOffset = 2208988800 // NTP epoch is 1900, Unix epoch is 1970
)

// NTP64 is NTP timestamp in RFC 5905 64-bit format: uint32 seconds since 1900-01-01, and uint32 fraction.
type NTP64 uint64

// Seconds returns integral seconds part of NTP64
func (n NTP64) Seconds() uint32 { _ = "STUB: not implemented"; return 0 }

// UTCSeconds returns seconds of NTP64 shifted to UNIX epoch.
func (n NTP64) UTCSeconds() uint64 { _ = "STUB: not implemented"; return 0 }

// Fraction returns 32-bit fractional part of NTP64.
func (n NTP64) Fraction() uint32 { _ = "STUB: not implemented"; return 0 }

// UTC returns NTP64 as UTC time in seconds.
func (n NTP64) UTC() float64 { _ = "STUB: not implemented"; return 0 }

// Time returns NTP64 as time.Time in UTC.
func (n NTP64) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NewNTP64 creates NTP64 from UTC time in seconds.
func NewNTP64(utcTime float64) NTP64 { _ = "STUB: not implemented"; return *new(NTP64) }

// String returns NTP64 as UTC time in string format.
func (n NTP64) String() string { _ = "STUB: not implemented"; return "" }
