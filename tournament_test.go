package tournament

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

// Define a function Tally(io.Reader, io.Writer) error.
//
// Note that unlike other tracks the Go version of the tally function
// should not ignore errors. It's not idiomatic Go to ignore errors.

var _ func(io.Reader, io.Writer) error = Tally

// These test what testers call the happy path, where there's no error.
var happyTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "good",
		input: `
Allegoric Alaskans;Blithering Badgers;3-0
Devastating Donkeys;Courageous Californians;2-2
Devastating Donkeys;Allegoric Alaskans;4-3
Courageous Californians;Blithering Badgers;0-1
Blithering Badgers;Devastating Donkeys;2-4
Allegoric Alaskans;Courageous Californians;6-1
`,
		expected: `
Team                           | MP |  W |  D |  L | GS | GR | P
Devastating Donkeys            |  3 |  2 |  1 |  0 | 10 |  7 | 7
Allegoric Alaskans             |  3 |  2 |  0 |  1 | 12 |  5 | 6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3 |  7 | 3
Courageous Californians        |  3 |  0 |  1 |  2 |  3 |  9 | 1
`[1:], // [1:] = strip initial readability newline
	},
	{
		description: "ignore comments and newlines",
		input: `

Allegoric Alaskans;Blithering Badgers;3-0
Devastating Donkeys;Courageous Californians;2-2
# Catastrophic Loss of the Californians
Devastating Donkeys;Allegoric Alaskans;4-3

Courageous Californians;Blithering Badgers;0-1
Blithering Badgers;Devastating Donkeys;2-4
Allegoric Alaskans;Courageous Californians;6-1


`,
		expected: `
Team                           | MP |  W |  D |  L | GS | GR | P
Devastating Donkeys            |  3 |  2 |  1 |  0 | 10 |  7 | 7
Allegoric Alaskans             |  3 |  2 |  0 |  1 | 12 |  5 | 6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3 |  7 | 3
Courageous Californians        |  3 |  0 |  1 |  2 |  3 |  9 | 1
`[1:],
	},
	{
		// A complete competition has all teams play eachother once or twice.
		description: "incomplete competition",
		input: `
Allegoric Alaskans;Blithering Badgers;3-0
Devastating Donkeys;Courageous Californians;2-2
Devastating Donkeys;Allegoric Alaskans;4-3
Allegoric Alaskans;Courageous Californians;6-1
`,
		expected: `
Team                           | MP |  W |  D |  L | GS | GR | P
Allegoric Alaskans             |  3 |  2 |  0 |  1 | 12 |  5 | 6
Devastating Donkeys            |  2 |  1 |  1 |  0 |  6 |  5 | 4
Courageous Californians        |  2 |  0 |  1 |  1 |  3 |  8 | 1
Blithering Badgers             |  1 |  0 |  0 |  1 |  0 |  3 | 0
`[1:],
	},
	{
		description: "tie for first and last place",
		input: `
Courageous Californians;Devastating Donkeys;3-2
Allegoric Alaskians;Blithering Badgers;3-1
Devastating Donkeys;Allegoric Alaskians;0-2
Courageous Californians;Blithering Badgers;2-1
Blithering Badgers;Devastating Donkeys;0-0
Allegoric Alaskians;Courageous Californians;1-1
`,
		expected: `
Team                           | MP |  W |  D |  L | GS | GR | P
Allegoric Alaskians            |  3 |  2 |  1 |  0 |  6 |  2 | 7
Courageous Californians        |  3 |  2 |  1 |  0 |  6 |  4 | 7
Blithering Badgers             |  3 |  0 |  1 |  2 |  2 |  5 | 1
Devastating Donkeys            |  3 |  0 |  1 |  2 |  2 |  5 | 1
`[1:],
	},
	{
		description: "teams tied on goal difference",
		input: `
Courageous Californians;Devastating Donkeys;1-1
Allegoric Alaskians;Blithering Badgers;2-1
Devastating Donkeys;Allegoric Alaskians;2-2
Courageous Californians;Blithering Badgers;2-1
Blithering Badgers;Devastating Donkeys;0-3
Allegoric Alaskians;Courageous Californians;1-1
`,
		expected: `
Team                           | MP |  W |  D |  L | GS | GR | P
Devastating Donkeys            |  3 |  1 |  2 |  0 |  6 |  3 | 4
Allegoric Alaskians            |  3 |  1 |  2 |  0 |  5 |  4 | 4
Courageous Californians        |  3 |  1 |  2 |  0 |  4 |  3 | 4
Blithering Badgers             |  3 |  0 |  0 |  3 |  2 |  7 | 0
`[1:],
	},
}

var errorTestCases = []string{
	"Bla;Bla;Bla",
	"Devastating Donkeys_Courageous Californians;1-0",
	"Devastating Donkeys@Courageous Californians;0-1",
	"Devastating Donkeys;Allegoric Alaskians;draw",
	"Devastating Donkeys;Allegoric Alaskians;-1-1",
	"Devastating Donkeys;Allegoric Alaskians;1-1-draw",
}

func TestTallyHappy(t *testing.T) {
	for _, tt := range happyTestCases {
		reader := strings.NewReader(tt.input)
		var buffer bytes.Buffer
		err := Tally(reader, &buffer)
		actual := buffer.String()
		// We don't expect errors for any of the test cases
		if err != nil {
			t.Fatalf("Tally for input named %q returned error %q. Error not expected.",
				tt.description, err)
		}
		if actual != tt.expected {
			t.Fatalf("Tally for input named %q was expected to return...\n%s\n...but returned...\n%s",
				tt.description, tt.expected, actual)
		}
	}
}

func TestTallyError(t *testing.T) {
	for _, s := range errorTestCases {
		reader := strings.NewReader(s)
		var buffer bytes.Buffer
		err := Tally(reader, &buffer)
		if err == nil {
			t.Fatalf("Tally for input %q should have failed but didn't.", s)
		}
		var _ error = err
	}
}

func BenchmarkTally(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range happyTestCases {
			var buffer bytes.Buffer
			Tally(strings.NewReader(tt.input), &buffer)
		}
		for _, s := range errorTestCases {
			var buffer bytes.Buffer
			Tally(strings.NewReader(s), &buffer)
		}
	}
}
