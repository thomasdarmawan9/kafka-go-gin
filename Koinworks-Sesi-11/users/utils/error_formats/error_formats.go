package error_formats

import (
	"strings"
	"users/utils/error_utils"
)

func ParseError(err error) error_utils.MessageErr {

	if strings.Contains(err.Error(), "no rows in result set") {
		return error_utils.NewNotFoundError("no record found")
	} else if strings.Contains(err.Error(), "violates unique constraint") {
		return error_utils.NewBadRequest("email has been taken, try another one")
	}

	return error_utils.NewInternalServerErrorr("something went wrong")
}
