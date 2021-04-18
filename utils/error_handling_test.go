package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type panicOnFailTestTable struct {
	expectedMessage    string
	errArg             error
	attemptedActionArg string
}

func TestPanicOnFail(t *testing.T) {
	assert := assert.New(t)
	testTable := []panicOnFailTestTable{
		{
			expectedMessage:    "",
			attemptedActionArg: "test action",
			errArg:             nil,
		},
		{
			expectedMessage:    "failed to test action, reason: panic error",
			attemptedActionArg: "test action",
			errArg:             errors.New("panic error"),
		},
	}

	for _, row := range testTable {
		panicFn := func() { PanicOnError(row.attemptedActionArg, row.errArg) }
		if row.errArg != nil {
			assert.PanicsWithError(
				row.expectedMessage,
				panicFn,
			)
		} else {
			assert.NotPanics(panicFn)
		}
	}
}
