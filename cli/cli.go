package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Cmd struct {
	Prefix        string
	Repeated      bool
	Unique        bool
	Count         bool
	IgnoreCase    bool
	NumFields     uint
	SkipChars     uint
	TakeChars     uint
	Range         bool
	Colorize      bool
	FormatCounter string
	BufferSize    uint
    Mapper        func(string) string
	Cutter        func(string) string
	Fprintln      func(io.Writer, string)
}

func New() *Cmd {

	return &Cmd{
		FormatCounter: "%d %s",
		Mapper:        func(s string) string { return s },
		Cutter:        func(s string) string { return s },
		Fprintln:      func(w io.Writer, s string) { fmt.Fprintln(w, s) },
	}
}

func (cmd *Cmd) Usage() {
	fmt.Printf(
		("%s 1.0\n" +
			"Author: Garry G.\n\n" +
			"Usage of %s:\n" +
			"uniq [-c|-d|-u|-p] [-f num_fields] [-s skip_chars] [-w check_chars] [-range] [-color] [input] [output]\n" +
			"if input\\output not specified, then stdin and stdout are used\n" +
			"\n"),
		filepath.Base(os.Args[0]),
		filepath.Base(os.Args[0]),
	)
	flag.PrintDefaults()
}

func (cmd *Cmd) Parse() {
	flag.Usage = cmd.Usage
	flag.BoolVar(&cmd.Count, "c", false, "Количество вхождений каждой строки")
	flag.BoolVar(&cmd.Repeated, "d", false, "Вывести только повторяющиеся строки")
	flag.BoolVar(&cmd.Unique, "u", false, "Вывести только уникальные строки")

	flag.StringVar(&cmd.Prefix, "p", "", "Количество строк в которых есть указанная подстрока")

	flag.BoolVar(&cmd.IgnoreCase, "i", false, "Игнорировать регистр при сравнении строк")
	flag.UintVar(&cmd.NumFields, "f", 0, "Игнорировать n полей разделенных пробелом с начала строки")
	flag.UintVar(&cmd.SkipChars, "s", 0, "Игнорировать n символов с начала строки")
	flag.UintVar(&cmd.TakeChars, "w", 0, "Проверять только n символов строки")

	flag.BoolVar(&cmd.Range, "range", false, "Показать использумый диапазон символов как срез")
	flag.BoolVar(&cmd.Colorize, "color", false, "Выделять использумый диапазон символов цветом")
    
    flag.UintVar(&cmd.BufferSize, "buffer-size", 0, "Установить максимальный размер буфера для сканирования файла (>64kb)")
	flag.Parse()
}
