package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"uniq/cli"
)

var testFile = (`AAA
aaa
bbb
bbb
ccc`)

var testFile10K = GenerateRandomStrings(10000)
var testFile100K = GenerateRandomStrings(100000)

var cmd = cli.New()

func init() {
	cmd.Fprintln = func(w io.Writer, s string) { fmt.Fprintln(w, s) }
}

func ExampleUnique() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }

	Unique(reader, writer, cmd)
	// Output:
	// AAA
	// aaa
	// ccc
}

func ExampleUniqueIgnoreCase() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = strings.ToLower
	cmd.Cutter = func(s string) string { return s }

	Unique(reader, writer, cmd)
	// Output:
	// ccc
}

func ExampleDuplicates() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }

	Duplicates(reader, writer, cmd)
	// Output:
	// bbb
}

func ExampleDuplicatesIgnoreCase() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = strings.ToLower
	cmd.Cutter = func(s string) string { return s }

	Duplicates(reader, writer, cmd)
	// Output:
	// aaa
	// bbb
}

func ExampleDeduplicate() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }

	Deduplicate(reader, writer, cmd)
	// Output:
	// AAA
	// aaa
	// bbb
	// ccc
}

func ExampleDeduplicateIgnoreCase() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = strings.ToLower
	cmd.Cutter = func(s string) string { return s }

	Deduplicate(reader, writer, cmd)
	// Output:
	// aaa
	// bbb
	// ccc
}

func ExampleCounterLines() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }

	CounterLines(reader, writer, cmd)
	// Output:
	// 1 AAA
	// 1 aaa
	// 2 bbb
	// 1 ccc

}

func ExampleCounterLinesIgnoreCase() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = strings.ToLower
	cmd.Cutter = func(s string) string { return s }

	CounterLines(reader, writer, cmd)
	// Output:
	// 2 aaa
	// 2 bbb
	// 1 ccc

}

func ExampleCounterLinesByPrefix() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	cmd.Prefix = "aa"

	CounterLinesByPrefix(reader, writer, cmd)
	// Output:
	// 1 aa
}

func ExampleCounterLinesByPrefixIgnoreCase() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd.Mapper = strings.ToLower
	cmd.Cutter = func(s string) string { return s }
	cmd.Prefix = "aa"

	CounterLinesByPrefix(reader, writer, cmd)
	// Output:
	// 2 aa
}

func TestSubstring(t *testing.T) {

	testCases := []struct {
		line     string
		options  [3]uint // fields,skip,take
		expected [2]uint // range
	}{
		{
            "123 456 789",
			[3]uint{0,0,0},
			[2]uint{0,11},  // [123 456 789]
		},

		{
            "123 456 789",
			[3]uint{1,2,0},
			[2]uint{6,11}, // 123 45[6 789]
		},
        
        {
            "123 456 789",
			[3]uint{1,2,1}, 
			[2]uint{6,7},  // 123 45[6] 789
		},
		
        {
            "123 456 789",
			[3]uint{0,10,0},
			[2]uint{10,11},   // 123 456 78[9]
		},
        
        {
            "123 456 789",
			[3]uint{0,0,11}, 
			[2]uint{0,11},  // [123 456 789]
		},
        
        {
            "123 456 789",
			[3]uint{3,0,0},  // 123 45[6] 789
			[2]uint{11,11},  // []
		},
        // fail case
        {
            "123 456 789",
			[3]uint{0,0,0},  // 
			[2]uint{0,0},   //   Substring(123 456 789, 0, 0, 0) = [0 11]; want [0 0]
		},
	}

	for _, c := range testCases {
		got := Substring(c.line, c.options[0], c.options[1], c.options[2])
		if got != c.expected {
			t.Errorf("Substring(%s, %d, %d, %d) = %v; want %v",
				c.line, c.options[0], c.options[1], c.options[2],
				got, c.expected)
		}
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateRandomStrings(count int) (outline string) {
	var builder strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		builder.WriteString(randSeq(10) + "\n")
	}
	outline = builder.String()
	return
}

func BenchmarkUnique10000(b *testing.B) {

	var reader = strings.NewReader(testFile10K)
	var writer = ioutil.Discard
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Unique(reader, writer, cmd)
	}
}

func BenchmarkDuplicates10000(b *testing.B) {

	var reader = strings.NewReader(testFile10K)
	var writer = ioutil.Discard
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Duplicates(reader, writer, cmd)
	}
}

func BenchmarkDeduplicate10000(b *testing.B) {

	var reader = strings.NewReader(testFile10K)
	var writer = ioutil.Discard
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Deduplicate(reader, writer, cmd)
	}
}

func BenchmarkUnique100000(b *testing.B) {

	var reader = strings.NewReader(testFile100K)
	var writer = ioutil.Discard
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Unique(reader, writer, cmd)
	}
}

func BenchmarkDuplicates100000(b *testing.B) {

	var reader = strings.NewReader(testFile100K)
	var writer = ioutil.Discard
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Duplicates(reader, writer, cmd)
	}
}

func BenchmarkDeduplicate100000(b *testing.B) {

	var reader = strings.NewReader(testFile100K)
	var writer = ioutil.Discard
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Deduplicate(reader, writer, cmd)
	}
}

// go test
// benchmarks
//go test -bench=. -benchtime=10x -benchmem
/*
BenchmarkUnique10000-4                10            400020 ns/op           36170 B/op       2001 allocs/op
BenchmarkDuplicates10000-4            10            100010 ns/op           20096 B/op       1001 allocs/op
BenchmarkDeduplicate10000-4           10            300020 ns/op           36170 B/op       2001 allocs/op
BenchmarkUnique100000-4               10           3200180 ns/op          324245 B/op      20001 allocs/op
BenchmarkDuplicates100000-4           10           1100070 ns/op          164096 B/op      10001 allocs/op
BenchmarkDeduplicate100000-4          10           3300190 ns/op          324222 B/op      20001 allocs/op
PASS
ok      uniq/utils      0.394s
*/

// analyze code coverage with tests
// go test ./... -coverprofile cover.out
/*
?       uniq    [no test files]
?       uniq/cli        [no test files]
--- FAIL: TestSubstring (0.00s)
    strings_test.go:211: Substring(123 456 789, 0, 0, 0) = [0 11]; want [0 0]
FAIL
coverage: 95.2% of statements
FAIL    uniq/utils      0.206s
FAIL
*/
// open in browser: go tool cover -html=cover.out
// generate html:  go tool cover -html=cover.out -o=cover.html