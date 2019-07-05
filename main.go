package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	defdatefmt = "20060102"
	usage      = `adddate adds years, months and days to a base date according to a format.

Usage:

	adddate [-y years] [-m months] [-d days]
	        [-b basedate | -f format | -b basedate -f format]

Arguments are:

	-y years
		adding years
	-m months
		adding months
	-d days
		adding days
	-b basedate
		base date (default: today (ex. %s))
		When don't use with -f option, a format is yyyymmdd
	-f format
		base date format by Go (default: %s)
		When don't use with -b option, a value is %s
`
)

var (
	years    int
	months   int
	days     int
	basedate string
	format   string
)

func init() {
	now := time.Now().Format(defdatefmt)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, now, defdatefmt, defdatefmt)
	}
	flag.IntVar(&years, "y", 0, "adding years")
	flag.IntVar(&months, "m", 0, "adding months")
	flag.IntVar(&days, "d", 0, "adding days")
	msgBasedate := fmt.Sprintf("base date (default: today (ex. %s))", now)
	flag.StringVar(&basedate, "b", now, msgBasedate)
	msgFormat := fmt.Sprintf("base date format by Go (default: %s)", defdatefmt)
	flag.StringVar(&format, "f", defdatefmt, msgFormat)
}

func main() {
	flag.Parse()
	d, err := time.Parse(format, basedate)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	s := d.AddDate(years, months, days).Format(format)
	if len(s) != len(format) {
		fmt.Fprintf(os.Stderr, "%s exceeds a format digit\n", s)
		os.Exit(1)
	}
	fmt.Println(s)
}
