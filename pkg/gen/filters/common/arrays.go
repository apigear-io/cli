package common

func Contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func GetEmptyStringList() []string {
	return []string {}
}

func AppendList(list []string, s string) []string {
	return append(list, s)
}

func IndexOf(a []any, s string) int {
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1
}
