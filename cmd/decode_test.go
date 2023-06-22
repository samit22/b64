package cmd

import "testing"

func Test_decodeFromBase64(t *testing.T) {
	type testCase struct {
		testCase string
		input    string
		output   string
		err      bool
	}

	cases := []testCase{
		{
			testCase: "when input is not base64 encoded it gives error",
			input:    "samit",
			err:      true,
		},
		{
			testCase: "when input is not base64 encoded it gives error",
			input:    "c2FtaXQ=",
			output:   "samit",
		},
	}

	for _, c := range cases {
		t.Run(c.testCase, func(t *testing.T) {
			op, err := decodeFromBase64(c.input)
			if c.err && err == nil {
				t.Errorf("error expected but not received")
			}
			if !c.err && c.output != op {
				t.Errorf("expected %s got %s", c.output, op)
			}
		})
	}
}
