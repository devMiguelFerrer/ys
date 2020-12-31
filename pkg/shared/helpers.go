package shared

import "strings"

// ConvertToVar transform a string (ClassName) to other (className).
func ConvertToVar(s string) string {
	splitted := strings.Split(s, "")
	splitted[0] = strings.ToLower(splitted[0])
	return strings.Join(splitted, "")
}

// ReplaceByRecipe method
func ReplaceByRecipe(base string, recipe []Recipe) string {
	var str string
	for _, r := range recipe {
		if str == "" {
			str = strings.ReplaceAll(base, r.Key, r.Value)
			continue
		}
		str = strings.ReplaceAll(str, r.Key, r.Value)
	}
	return str
}
