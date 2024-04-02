package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		c      config
		input  string
		output string
		err    error
	}{
		// {
		// 	c:      config{printUsage: true},
		// 	output: usageString + "\n",
		// },
		{
			c:      config{numTimes: 5},
			input:  "",
			output: strings.Repeat("Your word?\n", 1),
			err:    errors.New("enter your word"),
		},
		{
			c:     config{numTimes: 5},
			input: "Hello J", output: "Your word?\n" + strings.Repeat("your word is \"Hello J\"\n", 5),
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		rd := strings.NewReader(tc.input)
		err := runCmd(rd, byteBuf, tc.c)
		if err != nil && tc.err == nil {
			t.Fatalf("expected nil error, get: %v\n", err)
		}
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("expected error: %v, Got error: %v\n", tc.err.Error(), err.Error())
		}
		gotMsg := byteBuf.String()
		if gotMsg != tc.output {
			t.Errorf("expected stdout message to be: %v, Got: %v\n", tc.output, gotMsg)
		}
		byteBuf.Reset()
	}
}
