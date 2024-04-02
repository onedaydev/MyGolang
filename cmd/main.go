package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes int
	word     string
}

var errInvalidPosArgSpecified = errors.New("more than one positional argument specified")

func getWord(r io.Reader, w io.Writer) (string, error) {
	msg := "Your word?"
	fmt.Fprintln(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	word := scanner.Text()
	if len(word) == 0 {
		return "", errors.New("enter your word")
	}
	return word, nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("repeater", flag.ContinueOnError)
	fs.SetOutput(w) // FlagSet 객체의 진단 메시지 혹은 출력 메시지를 작성하는 데 사용할 writer를 지정
	fs.Usage = func() {
		var usageString = `
A repeater application which prints the name you entered a specified number of times.

Usage of %s: <options> [name]`
		fmt.Fprintf(w, usageString, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}

	fs.IntVar(&c.numTimes, "n", 0, "Number of times to repeat")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() > 1 {
		return c, errInvalidPosArgSpecified
	}
	if fs.NArg() == 1 {
		c.word = fs.Arg(0)
	}
	return c, nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func runCmd(rd io.Reader, w io.Writer, c config) error {
	var err error
	if len(c.word) == 0 {
		c.word, err = getWord(rd, w)
		if err != nil {
			return err
		}
	}
	printWord(c, w)
	return nil
}

func printWord(c config, w io.Writer) {
	msg := fmt.Sprintf(`your word is "%s"`, c.word)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintln(w, msg)
	}
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		if errors.Is(err, errInvalidPosArgSpecified) {
			fmt.Fprint(os.Stdout, err)
		}
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
