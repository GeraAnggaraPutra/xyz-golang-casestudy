package util

import "strings"

func ValidateUnique(err, errNew error) error {
	if strings.Contains(err.Error(), "SQLSTATE 23505") {
		err = errNew
	}

	return err
}
