package exercises

import "strings"

func getInitials(fullname string) string {

	initials := ""
	// Split fullname into all substrings separated by empty space
	names := strings.Split(fullname, " ")

	for _, v := range names {
		initials += v[:1]
	}

	return initials
}
