package utils

import (
	"fmt"
)

func PanicOnError(attemptedAction string, err error) {
	if err != nil {
		panic(fmt.Errorf("failed to %s, reason: %s", attemptedAction, err.Error()))
	}
}
