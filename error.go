package main

// Error type from string
type Error string

// Error defines an interface on an error
func (e Error) Error() string { return string(e) }
