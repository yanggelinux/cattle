package errors

import "testing"

func TestDoService(t *testing.T) {
	err := DoService()
	if err != nil {
		t.Log(err)
	}
}
