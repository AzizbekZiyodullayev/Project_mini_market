package helper

import (
	"strconv"
	"strings"

	"github.com/xtgo/uuid"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, "@"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, "@"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
