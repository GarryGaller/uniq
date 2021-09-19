package utils

import (
	"io"
	"math/rand"
	"os"
	"sort"
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

func ExampleUnique() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout

	cmd := cli.New()
	Unique(reader, writer, cmd)
	// Output:
	// AAA
	// aaa
	// ccc
}

func ExampleUniqueIgnoreCase() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd := cli.New()
	cmd.Mapper = strings.ToLower

	Unique(reader, writer, cmd)
	// Output:
	// ccc
}

func ExampleDuplicates() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd := cli.New()

	Duplicates(reader, writer, cmd)
	// Output:
	// bbb
}

func ExampleDuplicatesIgnoreCase() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd := cli.New()
	cmd.Mapper = strings.ToLower

	Duplicates(reader, writer, cmd)
	// Output:
	// aaa
	// bbb
}

func ExampleDeduplicate() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout

	cmd := cli.New()
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

	cmd := cli.New()
	cmd.Mapper = strings.ToLower

	Deduplicate(reader, writer, cmd)
	// Output:
	// aaa
	// bbb
	// ccc
}

func ExampleCounterLines() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd := cli.New()
	cmd.Mapper = func(s string) string { return s }

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
	cmd := cli.New()
	cmd.Mapper = strings.ToLower

	CounterLines(reader, writer, cmd)
	// Output:
	// 2 aaa
	// 2 bbb
	// 1 ccc

}

func ExampleCounterLinesByPrefix() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd := cli.New()
	cmd.Prefix = "aa"

	CounterLinesByPrefix(reader, writer, cmd)
	// Output:
	// 1 aa
}

func ExampleCounterLinesByPrefixIgnoreCase() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	cmd := cli.New()
	cmd.Mapper = strings.ToLower
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
			[3]uint{0, 0, 0},
			[2]uint{0, 11}, // [123 456 789]
		},

		{
			"123 456 789",
			[3]uint{1, 2, 0},
			[2]uint{6, 11}, // 123 45[6 789]
		},

		{
			"123 456 789",
			[3]uint{1, 2, 1},
			[2]uint{6, 7}, // 123 45[6] 789
		},

		{
			"123 456 789",
			[3]uint{0, 10, 0},
			[2]uint{10, 11}, // 123 456 78[9]
		},

		{
			"123 456 789",
			[3]uint{0, 0, 11},
			[2]uint{0, 11}, // [123 456 789]
		},

		{
			"123 456 789",
			[3]uint{3, 0, 0},
			[2]uint{11, 11}, // []
		},

		{
			"123 456 789",
			[3]uint{0, 12, 0},
			[2]uint{11, 11}, // []
		},

		/* // fail case
		   {
		       "123 456 789",
		       [3]uint{0,0,0},  //
		       [2]uint{0,0},   //   Substring(123 456 789, 0, 0, 0) = [0 11]; want [0 0]
		   },
		*/
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

	var data = make([]string, 0, count)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		data = append(data, randSeq(10))
	}
	sort.Strings(data)
	outline = strings.Join(data,"\n")
	return
}
  

func GenerateRepeatedStrings(count int) (outline string) {

	var data = make([]string, 0, count)
	var str = "aaaaaaaaaa"
    
    for i := 0; i < count; i++ {
		data = append(data, str)
	}
	
	outline = strings.Join(data,"\n")
	return
}


func BenchmarkUnique10000(b *testing.B) {

	var reader = strings.NewReader(testFile10K)
	var writer = io.Discard
	cmd := cli.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Unique(reader, writer, cmd)
	}
}

func BenchmarkDuplicates10000(b *testing.B) {

	var reader = strings.NewReader(testFile10K)
	var writer = io.Discard
	cmd := cli.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Duplicates(reader, writer, cmd)
	}
}

func BenchmarkDeduplicate10000(b *testing.B) {

	var reader = strings.NewReader(testFile10K)
	var writer = io.Discard
	cmd := cli.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Deduplicate(reader, writer, cmd)
	}
}

func BenchmarkUnique100000(b *testing.B) {

	var reader = strings.NewReader(testFile100K)
	var writer = io.Discard
	cmd := cli.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Unique(reader, writer, cmd)
	}
}

func BenchmarkDuplicates100000(b *testing.B) {

	var reader = strings.NewReader(testFile100K)
	var writer = io.Discard
	cmd := cli.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Duplicates(reader, writer, cmd)
	}
}

func BenchmarkDeduplicate100000(b *testing.B) {

	var reader = strings.NewReader(testFile100K)
	var writer = io.Discard
	cmd := cli.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Deduplicate(reader, writer, cmd)
	}
}

// go test -v ./utils
/*
=== RUN   TestSubstring
--- PASS: TestSubstring (0.00s)
=== RUN   ExampleUnique
--- PASS: ExampleUnique (0.00s)
=== RUN   ExampleUniqueIgnoreCase
--- PASS: ExampleUniqueIgnoreCase (0.00s)
=== RUN   ExampleDuplicates
--- PASS: ExampleDuplicates (0.00s)
=== RUN   ExampleDuplicatesIgnoreCase
--- PASS: ExampleDuplicatesIgnoreCase (0.00s)
=== RUN   ExampleDeduplicate
--- PASS: ExampleDeduplicate (0.00s)
=== RUN   ExampleDeduplicateIgnoreCase
--- PASS: ExampleDeduplicateIgnoreCase (0.00s)
=== RUN   ExampleCounterLines
--- PASS: ExampleCounterLines (0.00s)
=== RUN   ExampleCounterLinesIgnoreCase
--- PASS: ExampleCounterLinesIgnoreCase (0.00s)
=== RUN   ExampleCounterLinesByPrefix
--- PASS: ExampleCounterLinesByPrefix (0.00s)
=== RUN   ExampleCounterLinesByPrefixIgnoreCase
--- PASS: ExampleCounterLinesByPrefixIgnoreCase (0.00s)
PASS
ok      uniq/utils      0.202s
*/

// benchmarks
//go test -bench=. -benchmem ./utils
/*
BenchmarkUnique10000-4            292665              3851 ns/op            4098 B/op          1 allocs/op
BenchmarkDuplicates10000-4        307674              3566 ns/op            4096 B/op          1 allocs/op
BenchmarkDeduplicate10000-4       299982              3810 ns/op            4098 B/op          1 allocs/op
BenchmarkUnique100000-4           264003              3913 ns/op            4109 B/op          1 allocs/op
BenchmarkDuplicates100000-4       329068              3610 ns/op            4100 B/op          1 allocs/op
BenchmarkDeduplicate100000-4      265422              3956 ns/op            4109 B/op          1 allocs/op
PASS
ok      uniq/utils      8.680s
*/

// go test -cover ./utils
/*
PASS
coverage: 96.4% of statements
ok      uniq/utils      0.219s
*/

// analyze code coverage with tests
// go test ./... -coverprofile cover.out
/*
?       uniq    [no test files]
?       uniq/cli        [no test files]
ok      uniq/utils      0.230s  coverage: 96.4% of statements
*/
// open in browser: go tool cover -html=cover.out
// generate html:  go tool cover -html=cover.out -o=cover.html  


// go test -bench=. -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof ./utils
// go tool pprof -svg utils.test.exe mem.prof >mem.svg
// go tool pprof -svg utils.test.exe cpu.prof >cpu.svg  

// go test -bench=BenchmarkUnique10000 -benchmem -cpuprofile=cpu_unique.prof -memprofile=mem_unique.prof ./utils 