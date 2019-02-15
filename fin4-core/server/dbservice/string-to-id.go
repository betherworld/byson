package dbservice

import (
	"strconv"

	"fin4-core/server/datatype"
)

// StringToID converts string to ID
func StringToID(str string) (id datatype.ID, isOK bool) {
	idInt, err := strconv.Atoi(str)
	if err != nil {
		return datatype.ID(0), false
	}
	return datatype.ID(idInt), true
}
