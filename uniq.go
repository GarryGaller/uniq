package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"uniq/cli"
	"uniq/utils"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	//"github.com/mattn/go-colorable"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func setReader(reader io.Reader, path string) (r io.Reader, err error) {
	r = reader

	if path != "" {
		r, err = os.OpenFile(path, os.O_RDONLY, 0644)
	}

	return
}

func setWriter(writer io.Writer, path string) (w io.Writer, err error) {
	w = writer

	if path != "" {
		w, err = os.OpenFile(path, os.O_WRONLY, 0644)
	}

	return
}

func main() {
	var reader io.Reader
	var writer io.Writer
	var err error

	cmd := cli.New()
	cmd.Parse()

	var groupCDU int8
	if cmd.Count {
		groupCDU += 1
	}
    
    if cmd.Prefix != "" {
		groupCDU += 1
	}

	if cmd.Unique {
		groupCDU += 1
	}

	if cmd.Repeated {
		groupCDU += 1
	}

	if groupCDU > 1 {
		fmt.Println("Опции группы {-c|-d|-u|-p} взаимоисключающие")
		flag.Usage()
		os.Exit(0)
	}

	inputOutput := [2]string{"", ""}

	for i, arg := range flag.Args() {
		inputOutput[i] = arg
	}
	var builder strings.Builder
	cmd.Mapper = func(s string) string { return s }
	cmd.Cutter = func(s string) string { return s }
	cmd.Fprintln = func(writer io.Writer, line string) {

		if cmd.Colorize || cmd.Range {
			idx := utils.Substring(line,
				cmd.NumFields, cmd.SkipChars, cmd.TakeChars,
			)

			if cmd.Colorize {
				builder.Reset()
				//line = line[:idx[0]] + color.GreenString(line[idx[0]:idx[1]]) + line[idx[1]:]
				builder.WriteString(line[:idx[0]])
				builder.WriteString(color.GreenString(line[idx[0]:idx[1]]))
				builder.WriteString(line[idx[1]:])
				line = builder.String()
			}
			
            if cmd.Range {
				line = fmt.Sprintf("[%d:%d] %s", idx[0], idx[1], line)
			}
		    // for coloring works only fmt.Fprintf
            fmt.Fprintf(writer, "%s\n", line)
        } else {
            // Optimal ?
            io.WriteString(writer, line + "\n")
        }
    }

	if cmd.IgnoreCase {
		cmd.Mapper = strings.ToLower
	}

	if cmd.NumFields != 0 || cmd.SkipChars != 0 || cmd.TakeChars != 0 {
		cmd.Cutter = func(line string) string {
			idx := utils.Substring(line,
				cmd.NumFields, cmd.SkipChars, cmd.TakeChars,
			)
			// to avoid unnecessary attempts to take a slice
			if idx[0] != 0 || idx[1] != uint(len(line)) {
				line = line[idx[0]:idx[1]]
			}
			return line
		}
	}

	reader, err = setReader(os.Stdin, inputOutput[0])
	check(err)

	if file, ok := reader.(*os.File); ok {
		// not stdin
		if !isatty.IsTerminal(file.Fd()) {
			defer file.Close()
		}
	}
	if cmd.Colorize {
		writer, err = setWriter(color.Output, inputOutput[1])
	} else {
		writer, err = setWriter(os.Stdout, inputOutput[1])
	}
	check(err)
	if file, ok := writer.(*os.File); ok {
		// not stdout
		if !isatty.IsTerminal(file.Fd()) {
			defer file.Close()
		}
	}

	//==========================
	if cmd.Count {
		utils.CounterLines(reader, writer, cmd)
	} else if cmd.Prefix != "" {
		utils.CounterLinesByPrefix(reader, writer, cmd)
	} else if cmd.Unique {
		utils.Unique(reader, writer, cmd)
	} else if cmd.Repeated {
		utils.Duplicates(reader, writer, cmd)
	} else {
		utils.Deduplicate(reader, writer, cmd)
	}
}
