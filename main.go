package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"
)

var cfg = struct {
	now  time.Time
	fmt  string
	span time.Duration

	// for cmdline
	nowint  int64
	spanint int64
}{}

func init() {
	cfg.now = time.Now()
	// consider epochs within ~2 years to be dates by default
	cfg.span = time.Hour * 24 * 365 * 2
}

// input returns the input for etod to read from;  a file to output or
// stdin if it's a pipe and there are no files
func input() (io.Reader, error) {
	args := flag.Args()
	switch len(args) {
	case 0:
		fi, _ := os.Stdin.Stat()
		if (fi.Mode() & os.ModeCharDevice) != 0 {
			break
		}
		return os.Stdin, nil
	case 1:
		return os.Open(args[0])
	}
	return nil, errors.New("usage: etod [opts] <file> or cmd | etod")
}

func main() {
	flag.Int64Var(&cfg.nowint, "now", time.Now().Unix(), "alternative 'current' epoch")
	flag.Int64Var(&cfg.spanint, "span", int64(cfg.span/(time.Hour*24)), "span in days around 'now' to consider a date")
	flag.StringVar(&cfg.fmt, "fmt", "2006-01-02T15:04:05", "time format to display dates in")
	flag.Parse()

	if cfg.spanint > 0 {
		cfg.span = time.Duration(cfg.spanint) * 24 * time.Hour
	}

	if cfg.nowint > 0 {
		cfg.now = time.Unix(cfg.nowint, 0)
	}

	r, err := input()
	if err != nil {
		fmt.Println(err)
		return
	}

	start, end := cfg.now.Add(-cfg.span), cfg.now.Add(cfg.span)
	matcher := regexp.MustCompile(`(\d{10})`)
	s := bufio.NewScanner(r)

	for s.Scan() {
		b := s.Bytes()
		b = matcher.ReplaceAllFunc(b, func(in []byte) []byte {
			i, err := strconv.ParseInt(string(in), 10, 64)
			if err != nil {
				return in
			}
			target := time.Unix(i, 0)
			if target.After(start) && target.Before(end) {
				return []byte(target.Format(cfg.fmt))
			}
			return in
		})
		fmt.Println(string(b))
	}
}
