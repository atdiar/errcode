// Package errcode defines common error codes used throughout libraries.
// It is supposed to be extended. Creates a Domain Specific Error Universe.
package errcode

// The code names should be self-explanatory enough.
const (
	Unknown   = 0
	BadCookie = iota + 127 // Don't ask why 127
	BadSession
	BadStorage
	Expired
	KeyNotFound
	JSONEncoding
	NoCookie
	NoID
)
