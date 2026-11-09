package main

import "strings"

// RFC 9110, 5.6.2
func isTchar(b byte) bool {
	switch {
	case 'a' <= b && b <= 'z', 'A' <= b && b <= 'Z', '0' <= b && b <= '9':
		return true
	}
	return strings.IndexByte("!#$%&'*+-.^_`|~", b) >= 0
}

func isToken(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !isTchar(s[i]) {
			return false
		}
	}
	return true
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
