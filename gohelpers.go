// Copyright Kueski. All rights reserved.
// Use of this source code is not licensed

// Package provides helper functions.
package gohelpers

import (
	"encoding/json"
	"runtime"
)

// Function IsJSON checks if the supplied string is in a
// valid JSON format. Returns true if the string is valid,
// false if not
func IsJSON(str string) bool {
	var js json.RawMessage

	err := json.Unmarshal([]byte(str), &js)
	return err == nil
}

// Function GetFunctionName returns the name of the current
// running function
func GetFunctionName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}