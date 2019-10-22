package models

import (
	"fmt"

	"github.com/tidwall/gjson"
	null "gopkg.in/guregu/null.v3"
)

// RunResult keeps track of the outcome of a TaskRun or JobRun. It stores the
// Data and ErrorMessage, and contains a field to track the status.
type RunResult struct {
	ID           uint        `json:"-" gorm:"primary_key;auto_increment"`
	Data         JSON        `json:"data" gorm:"type:text"`
	ErrorMessage null.String `json:"error"`
}

// SetError marks the result as errored and saves the specified error message
func (rr *RunResult) SetError(err error) {
	if err != nil {
		rr.ErrorMessage = null.StringFrom(err.Error())
	}
}

// ResultString returns the string result of the Data JSON field.
func (rr RunResult) ResultString() (string, error) {
	val := rr.Result()
	if val.Type != gjson.String {
		return "", fmt.Errorf("non string result")
	}
	return val.String(), nil
}

// Result returns the result as a gjson object
func (rr RunResult) Result() gjson.Result {
	return rr.Data.Get("result")
}

// ErrorString returns the error as a string if present, otherwise "".
func (rr RunResult) ErrorString() string {
	return rr.ErrorMessage.ValueOrZero()
}
