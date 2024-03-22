package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func getName(r io.Reader, w io.Writer) (string, error) {
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

type config struct {
	numTimes int
	// printUsage bool
}

// var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
// A printer application which prints the word you entered <integer> number of times.`, os.Args[0])

// func printUsage(w io.Writer) {
// 	fmt.Fprintln(w, usageString)
// }

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("repeater", flag.ContinueOnError)
	fs.SetOutput(w) // FlagSet 객체의 진단 메시지 혹은 출력 메시지를 작성하는 데 사용할 writer를 지정
	fs.IntVar(&c.numTimes, "n", 0, "number of times to repeat word")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() != 0 {
		return c, errors.New("positional arguments specified")
	}

	return c, nil
	// var numTimes int
	// var err error
	// c := config{}

	// if len(args) != 1 {
	// 	return c, errors.New("invalid number of arguments")
	// }

	// if args[0] == "-h" || args[0] == "--help" {
	// 	c.printUsage = true
	// 	return c, nil
	// }

	// numTimes, err = strconv.Atoi(args[0])
	// if err != nil {
	// 	return c, err
	// }
	// c.numTimes = numTimes

	// return c, nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	// if c.printUsage {
	//	 printUsage(w)
	// 	return nil
	// }

	word, err := getName(r, w)
	if err != nil {
		return err
	}
	printWord(c, word, w)
	return nil
}

func printWord(c config, word string, w io.Writer) {
	msg := fmt.Sprintf(`your word is "%s"`, word)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintln(w, msg)
	}
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		// printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		// printUsage(os.Stdout)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
