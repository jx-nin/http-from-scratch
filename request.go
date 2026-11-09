package main

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Target  string
	Version string
}

var supportedMethods = map[string]bool{
	"GET":  true,
	"HEAD": true,
	"POST": true,
}

func validTarget(t string) bool {
	if !strings.HasPrefix(t, "/") {
		return false
	}
	for i := 0; i < len(t); i++ {
		if t[i] < 0x21 || t[i] == 0x7F {
			return false
		}
	}
	return true
}

func validVersion(v string) bool {
	return len(v) == 8 &&
		strings.HasPrefix(v, "HTTP/") &&
		isDigit(v[5]) && v[6] == '.' && isDigit(v[7])
}

func parseRequestLine(line string) (Request, error) {
	line = strings.TrimSuffix(line, "\r\n")
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return Request{}, fmt.Errorf("malformed request line: %q", line)
	}

	method, target, version := parts[0], parts[1], parts[2]

	if !isToken(method) {
		return Request{}, fmt.Errorf("invalid method: %q", method)
	}
	if !validTarget(target) {
		return Request{}, fmt.Errorf("invalid target: %q", target)
	}
	if !validVersion(version) {
		return Request{}, fmt.Errorf("invalid version: %q", version)
	}

	return Request{Method: method, Target: target, Version: version}, nil
}
