package utils

func Contains(a []string, v string) bool {
	for _, s := range a {
		if s == v {
			return true
		}
	}
	return false
}
