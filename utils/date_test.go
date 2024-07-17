package utils

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {

	dateString := DateFormat(time.Now())

	t.Logf(dateString)
	t.Logf("SUCCESS")
	t.Errorf("printALL")
}