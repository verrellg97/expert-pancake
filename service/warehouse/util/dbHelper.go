package util

func WildCardString(keyword string) string {
	if keyword != "" {
		return "%" + keyword + "%"
	} else {
		return "%"
	}
}
