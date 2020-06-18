package name

import (
	"strings"

	format "github.com/nurali-techie/module-format"
)

func FullName(firstName, lastName string) string {
	fullName := strings.Join([]string{format.Name(firstName), format.Name(lastName)}, " ")
	return fullName
}
