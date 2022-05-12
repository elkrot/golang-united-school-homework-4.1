package reverse_string
import (
	"strings"
)
func ReverseString(input string) (output string) {
	runes := []rune(input)
	o := []string{}
	for i := len(runes) - 1; i >= 0; i-- {
		o = append(o, string(runes[i]))
	}
	return strings.Join(o, "")
}
