# adddate

adddate adds years, months and days to a base date according to a format.

## Installation

When you have installed the Go, Please execute following `go get` command:

```sh
go get -u github.com/qt-luigi/adddate
```

## Usage

```sh
$ adddate -h
adddate adds years, months and days to a base date according to a format.

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
		base date (default: today (ex. 20190701))
		When don't use with -f option, a format is yyyymmdd
	-f format
		base date format by Go (default: 20060102)
		When don't use with -b option, a value is 20060102
```

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)

