package bodycomposition

import (
	"testing"
	"time"
)

func TestNewBodyCompositionWithDefault(t *testing.T) {

	bc := NewBodyComposition(80, 14.4, 55.2, 37, 45.5, 21, 5, 23, 2250, 12.6, -1)

	now := time.Now()

	if !bc.TimeStamp.After(now.Add(time.Second * -5)) {
		t.Fatal("Timestamp must be within the last 5 seconds")
	}

	if !bc.TimeStamp.Before(now.Add(time.Second * 5)) {
		t.Fatal("Timestamp must be within the next 5 seconds")
	}
}

func TestNewBodyComposition(t *testing.T) {
	// 30.12.2021 17:15:20
	var timeStamp int64 = 1640884520
	// set local time to UTC, no matter where executed
	time.Local = time.UTC

	bc := NewBodyComposition(80, 14.4, 55.2, 37, 45.5, 21, 5, 23, 2250, 12.6, timeStamp)

	year, month, date := bc.TimeStamp.Date()
	hour, min, sec := bc.TimeStamp.Clock()

	if year != 2021 {
		t.Fatalf("Year should be 2021, was %d", year)
	}

	if month != 12 {
		t.Fatalf("Month should be 12, was %d", month)
	}

	if date != 30 {
		t.Fatalf("Date should be 30, was %d", date)
	}

	if hour != 17 {
		t.Fatalf("Hour should be 17, was %d", hour)
	}

	if min != 15 {
		t.Fatalf("Minute should be 15, was %d", min)
	}

	if sec != 20 {
		t.Fatalf("Second should be 20, was %d", sec)
	}
}
