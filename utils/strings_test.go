package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
	"uniq/cli"
)

var testFile = (
`AAA
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
