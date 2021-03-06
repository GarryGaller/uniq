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
 

func setBuffer(scanner *bufio.Scanner, bufferSize uint) {
    if bufferSize*1024 > bufio.MaxScanTokenSize {
        buf := make([]byte, bufferSize*1024)
        scanner.Buffer(buf, int(bufferSize*1024))
    }
}

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
    setBuffer(scanner, cmd.BufferSize)

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
    setBuffer(scanner, cmd.BufferSize)
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
    setBuffer(scanner, cmd.BufferSize)
    scanner.Scan()
    prev := cmd.Mapper(scanner.Text())
    cnt := 1

    for scanner.Scan() {
        curr := cmd.Mapper(scanner.Text())
        if cmd.Cutter(prev) != cmd.Cutter(curr) {
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

func CounterLines(
    reader io.Reader,
    writer io.Writer,
    cmd *cli.Cmd) {
    /* Prefix lines by the number of occurrences */

    scanner := bufio.NewScanner(reader)
    setBuffer(scanner, cmd.BufferSize)
    scanner.Scan()
    prev := cmd.Mapper(scanner.Text())
    cnt := 1

    for scanner.Scan() {
        curr := cmd.Cutter(cmd.Mapper(scanner.Text()))
        if cmd.Cutter(prev) != cmd.Cutter(curr) {
            cmd.Fprintln(writer, fmt.Sprintf(cmd.FormatCounter, cnt, prev))
            prev = curr
            cnt = 1
        } else {
            cnt += 1
        }
    }
    cmd.Fprintln(writer, fmt.Sprintf(cmd.FormatCounter, cnt, prev))
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    
}

func CounterLinesByPrefix(
    reader io.Reader,
    writer io.Writer,
    cmd *cli.Cmd) {
    /*The number of rows in which there is a specified substring*/

    scanner := bufio.NewScanner(reader)
    setBuffer(scanner, cmd.BufferSize)
    cnt := 0

    for scanner.Scan() {
        line := cmd.Cutter(cmd.Mapper(scanner.Text()))
        if strings.HasPrefix(line, cmd.Prefix) {
            cnt += 1
        }
    }

    cmd.Fprintln(writer, fmt.Sprintf(cmd.FormatCounter, cnt, cmd.Prefix))
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
