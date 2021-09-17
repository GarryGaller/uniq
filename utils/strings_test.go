package utils

import (
	"fmt"
	"io"
	"io/ioutil"
    "math/rand"
	"os"
	"strings"
	"time"
    "testing"
    
	"uniq/cli"
)

var testFile = (`AAA
aaa
bbb
bbb
ccc`)

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

	testFile := GenerateRandomStrings(10000)
	var reader = strings.NewReader(testFile)
	var writer = ioutil.Discard
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Unique(reader, writer, cmd)
	}
}

func BenchmarkDuplicates10000(b *testing.B) {

	testFile := GenerateRandomStrings(10000)
	var reader = strings.NewReader(testFile)
	var writer = ioutil.Discard
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Duplicates(reader, writer, cmd)
	}
}

func BenchmarkDeduplicate10000(b *testing.B) {

	testFile := GenerateRandomStrings(10000)
	var reader = strings.NewReader(testFile)
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
BenchmarkUnique10000-4                10            400030 ns/op           36148 B/op       2001 allocs/op
BenchmarkDuplicates10000-4            10            100000 ns/op           20096 B/op       1001 allocs/op
BenchmarkDeduplicate10000-4           10            400020 ns/op           36170 B/op       2001 allocs/op
*/