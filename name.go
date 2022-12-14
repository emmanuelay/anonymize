// Package Anonymize is a small and simple package for anonymizing names, e-mails and domains.
//
// Copyright 2022 Emmanuel Ay. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package anonymize

import (
	"strings"
)

// Name anonymizes the input string based on letter case.
func Name(input string) string {
	return processName(input, '*')
}

// NameWithCustomRune anonymizes the input string based on letter case. The 'anonRune' parameter is to use your custom rune to hide characters.
func NameWithCustomRune(input string, anonRune rune) string {
	return processName(input, anonRune)
}

func processName(input string, anonRune rune) string {

	parts := strings.Split(input, " ")

	for idx, part := range parts {
		parts[idx] = processWord(part, true, anonRune)
	}

	return strings.Join(parts, " ")
}
