package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
    
    "uniq/cli"
    
)

var reWord *regexp.Regexp = regexp.MustCompile(`\b(\S+)\b`)

func Substring(
	line string,
	numFields, skipChars, takeChars uint) (idx [2]uint) {
	var (
		start uint
		end   uint = uint(len(line))
	)

	if numFields > 0 {
		fields := reWord.FindAllStringIndex(line, -1)
		lf := uint(len(fields))
		if numFields < lf {
			start = uint(fields[numFields][0])
		} else {
			start = end
		}
	}

	if skipChars > 0 {
		ll := uint(len(line[start:]))
		if skipChars < ll {
			start += skipChars
		} else {
			start = end
		}
	}

	if takeChars > 0 {
		ll := uint(len(line[start:]))
		if takeChars < ll {
			end = start + takeChars
		}
	}

	idx[0] = start
	idx[1] = end

	return
}

func Deduplicate(
	reader io.Reader,
	writer io.Writer,
	cmd *cli.Cmd) {

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	prev := cmd.Mapper(scanner.Text())

	for scanner.Scan() {
		curr := cmd.Mapper(scanner.Text())
		if cmd.Cutter(prev) != cmd.Cutter(curr) {
			cmd.Fprintln(writer, prev)
			prev = curr
		}
	}

	cmd.Fprintln(writer, prev)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func Duplicates(
	reader io.Reader,
	writer io.Writer,
	cmd *cli.Cmd) {

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	prev := cmd.Mapper(scanner.Text())
	cnt := 1

	for scanner.Scan() {
		curr := cmd.Mapper(scanner.Text())
		if cmd.Cutter(prev) != cmd.Cutter(curr) {
			prev = curr
			cnt = 1
		} else {
			if cnt == 1 {
				cmd.Fprintln(writer, prev)
			}
			cnt += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func Unique(
	reader io.Reader,
	writer io.Writer,
	cmd *cli.Cmd) {

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	prev := cmd.Mapper(scanner.Text())
	cnt := 1

	for scanner.Scan() {
		curr := cmd.Mapper(scanner.Text())
		if cmd.Cutter(prev) != cmd.Cutter(curr){
			if cnt == 1 {
				cmd.Fprintln(writer, prev)
			}
			prev = curr
			cnt = 1
		} else {
			cnt += 1
		}
	}
	if cnt == 1 {
		cmd.Fprintln(writer, prev)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func FindPrefix(
	prefix string,
	reader io.Reader,
	cmd *cli.Cmd) (count int) {
	/* Prefix lines by the number of occurrences */

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := cmd.Cutter(cmd.Mapper(scanner.Text()))
		if strings.HasPrefix(line, prefix) {
			count += 1
		}
	}
	return
}
