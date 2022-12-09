package util

import (
	"errors"
	"fmt"
	"time"
)

func TimeNow() string {
	timeNow := time.Now()
	return fmt.Sprintf("%v", timeNow)
}
func SetError(str string) error {
	return errors.New(str)
}
func ObjectIdToString(Id interface{}) string {
	var ID string
	idpass := fmt.Sprintf("%v", Id)
	if len(idpass) != 0 {
		lengthStr := len(idpass)
		ID = idpass[10 : lengthStr-2]
	}
	return ID
}
