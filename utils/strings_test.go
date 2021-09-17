package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
	"uniq/cli"
)

var testFile = `
aaa
aaa
bbb
ccc`

var cmd = cli.New()

func init() {

	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	cmd.Fprintln = func(w io.Writer, s string) { fmt.Fprintln(w, s) }
}

func ExampleUnique() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout

	Unique(reader, writer, cmd)
	// Output:
	// bbb
	// ccc
}

func ExampleDuplicates() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout
	Duplicates(reader, writer, cmd)
	// Output:
	// aaa
}

func ExampleDeduplicate() {
	var reader = strings.NewReader(testFile)
	var writer = os.Stdout

	Deduplicate(reader, writer, cmd)
	// Output:
	// aaa
	// bbb
	// ccc
} 

func ExampleFindPrefix() {
	var reader = strings.NewReader(testFile)
	
    count := FindPrefix("aa", reader, cmd)
    fmt.Println(count)
	// Output:
	// 2
}
