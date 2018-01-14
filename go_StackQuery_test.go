package goStackQuery

import (
	"errors"
	"testing"
)

func TestQuery(t *testing.T) {
	err := errors.New("Атстой")

	Query(err, 0)

	// if resultOGRN[0].ID != okOGRN {
	// 	t.Fatalf("Want %v, but got %v", resultOGRN, okOGRN)
	// }
}
