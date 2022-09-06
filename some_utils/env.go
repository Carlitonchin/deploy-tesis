package some_utils

import (
	"os"
	"strconv"

	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
)

func GetUintEnv(key string) (uint, error) {
	ui, err := strconv.ParseUint(os.Getenv(key), 10, 64)

	if err != nil {
		type_error := apperrors.Internal
		message := "Error intentando leer del environment"

		err = apperrors.NewError(type_error, message)
	}

	return uint(ui), err
}
