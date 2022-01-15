package utils

import (
	"strings"
)

// FormatQuery 将url的query变成JSON的键值对模式, 例如：query=123&u=456 ==> {"query":"123","u":"456"}
func FormatQuery(query string) string {
	if query != "" {
		queryParts := strings.Split(query, "&")
		var temp strings.Builder
		temp.Grow(20)
		for _, part := range queryParts {
			s := strings.Split(part, "=")
			temp.WriteString("\"")
			temp.WriteString(s[0])
			temp.WriteString("\":\"")
			temp.WriteString(s[1])
			temp.WriteString("\"")
			temp.WriteString(",")
			//temp = temp + fmt.Sprintf(`"%s":"%s",`, s[0], s[1])
		}
		return "{" + strings.TrimRight(temp.String(), ",") + "}"
	}
	return "{}"
}
