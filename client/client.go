package main

import (
	"net/http"
	"strings"
)

func main() {
	var tripper http.RoundTripper = &RoundTripperCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)
}

// RoundTripperCounter ...
type RoundTripperCounter struct {
	count int
}

// RoundTrip ...
func (rt *RoundTripperCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count++
	return nil, nil
}
