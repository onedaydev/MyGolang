package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		args []string
		config
		output string
		err    error
	}{
		{
			args: []string{"-h"},
			output: `
A repeater application which prints the name you entered a specified number of times.

Usage of repeater: <options> [name]

Options: 
  -n int
    	Number of times to repeat
`,
			err:    errors.New("flag: help requested"),
			config: config{numTimes: 0},
		},
		{
			args:   []string{"-n", "10"},
			err:    nil,
			config: config{numTimes: 10},
		},
		{
			args:   []string{"-n", "abc"},
			err:    errors.New("invalid value \"abc\" for flag -n: parse error"),
			config: config{numTimes: 0},
		},
		{
			args:   []string{"-n", "1", "foo bar"},
			err:    nil,
			config: config{numTimes: 1, word: "foo bar"},
		},
		{
			args:   []string{"-n", "1", "foo", "bar"},
			err:    errors.New("more than one positional argument specified"),
			config: config{numTimes: 1},
		},
	}
	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		c, err := parseArgs(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Fatalf("expected nil error, got: %v\n", err)
		}
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("expected error to be: %v, got: %v\n", tc.err, err)
		}
		if c.numTimes != tc.numTimes {
			t.Errorf("expected numTimes to be: %v, got: %v\n", tc.numTimes, c.numTimes)
		}
		gotMsg := byteBuf.String()
		if len(tc.output) != 0 && gotMsg != tc.output {
			t.Errorf("expected stdout message to be: %#v, got%#v\n", tc.output, gotMsg)
		}
		byteBuf.Reset()
	}
}
