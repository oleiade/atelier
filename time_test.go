package atelier

import (
	"testing"
	"time"
)

func TestParsePeriod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected time.Duration
		wantErr  bool
	}{
		// Valid second cases
		{input: "1s", expected: time.Second, wantErr: false},
		{input: "2s", expected: 2 * time.Second, wantErr: false},
		{input: "1 sec", expected: time.Second, wantErr: false},
		{input: "2 secs", expected: 2 * time.Second, wantErr: false},
		{input: "1 second", expected: time.Second, wantErr: false},
		{input: "2 seconds", expected: 2 * time.Second, wantErr: false},

		// Valid minute cases
		{input: "1min", expected: time.Minute, wantErr: false},
		{input: "2mins", expected: 2 * time.Minute, wantErr: false},
		{input: "1 min", expected: time.Minute, wantErr: false},
		{input: "2 mins", expected: 2 * time.Minute, wantErr: false},
		{input: "1 minute", expected: time.Minute, wantErr: false},
		{input: "2 minutes", expected: 2 * time.Minute, wantErr: false},

		// Valid hour cases
		{input: "1h", expected: time.Hour, wantErr: false},
		{input: "2h", expected: 2 * time.Hour, wantErr: false},
		{input: "1 hour", expected: time.Hour, wantErr: false},
		{input: "2 hours", expected: 2 * time.Hour, wantErr: false},

		// Valid day cases
		{input: "1d", expected: 24 * time.Hour, wantErr: false},
		{input: "2d", expected: 48 * time.Hour, wantErr: false},
		{input: "1 day", expected: 24 * time.Hour, wantErr: false},
		{input: "2 days", expected: 48 * time.Hour, wantErr: false},

		// Valid week cases
		{input: "1w", expected: 7 * 24 * time.Hour, wantErr: false},
		{input: "2w", expected: 14 * 24 * time.Hour, wantErr: false},
		{input: "1 week", expected: 7 * 24 * time.Hour, wantErr: false},
		{input: "2 weeks", expected: 14 * 24 * time.Hour, wantErr: false},

		// Valid month cases
		{input: "1m", expected: 30 * 24 * time.Hour, wantErr: false},
		{input: "2m", expected: 60 * 24 * time.Hour, wantErr: false},
		{input: "1 month", expected: 30 * 24 * time.Hour, wantErr: false},
		{input: "2 months", expected: 60 * 24 * time.Hour, wantErr: false},

		// Valid year cases
		{input: "1y", expected: 365 * 24 * time.Hour, wantErr: false},
		{input: "2y", expected: 730 * 24 * time.Hour, wantErr: false},
		{input: "1 year", expected: 365 * 24 * time.Hour, wantErr: false},
		{input: "2 years", expected: 730 * 24 * time.Hour, wantErr: false},

		// Valid zero case
		{input: "0d", expected: 0, wantErr: false},

		// Invalid cases
		{input: "", expected: 0, wantErr: true},
		{input: "5x", expected: 0, wantErr: true},
		{input: "abc", expected: 0, wantErr: true},
		{input: "1 days", expected: 0, wantErr: true},
		{input: "2 day", expected: 0, wantErr: true},
		{input: "2 min", expected: 0, wantErr: true},
	}

	for _, test := range tests {
		duration, err := ParseTimePeriod(test.input)
		if duration != test.expected {
			t.Errorf("ParsePeriod(%s) returned unexpected duration: got %v, want %v", test.input, duration, test.expected)
		}

		if (err != nil) != test.wantErr {
			t.Errorf("ParsePeriod(%s) returned unexpected error: got %v, wantErr %v", test.input, err, test.wantErr)
		}
	}
}

func TestParsePeriodMax(t *testing.T) {
	t.Parallel()

	// Calculate the expected duration from the beginning of time to now
	wantDuration := time.Since(BeginningOfTime)

	duration, err := ParseTimePeriod("max")
	if err != nil {
		t.Errorf("ParsePeriod(\"max\") returned unexpected error: %v", err)
	}

	delta := wantDuration - duration
	if delta < 0 {
		delta = -delta
	}

	// Allow a delta of up to 5 seconds
	if delta > 1*time.Second {
		t.Errorf("ParsePeriod(\"max\") returned unexpected duration: got %v, want %v", duration, wantDuration)
	}
}
