package error_formats

import (
	"orders/utils/error_utils"
	"strings"
)

func ParseError(err error) error_utils.MessageErr {

	if strings.Contains(err.Error(), "no rows in result set") {
		return error_utils.NewNotFoundError("no record found")
	}

	return error_utils.NewInternalServerErrorr("something went wrong")
}
