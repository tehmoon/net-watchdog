package main

import (
	"strconv"
	"time"
	"regexp"
	"bufio"
	"log"
	"os"
	"github.com/tehmoon/errors"
)

var (
	Reg = regexp.MustCompile(`(\d+-\d+-\d+\ \d+:\d+:\d+)\ \w+\ login\[\d+\]:\ FAILED\ LOGIN\ \(\d+\)\ on\ '/dev/tty(\w+)'\ FOR\ '\w+',\ User\ not\ known\ to\ the\ underlying\ authentication\ module`)
	ErrNotRegularFile = errors.New("Not regular file.")
)

func main() {
	file, err := openRegularFile("/var/log/auth.log")
	if err != nil {
		log.Printf("Err in opening regular file: %v\n", err.Error())
		os.Exit(2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Printf("Err in scanning text: %s\n", err.Error())
			os.Exit(2)
		}

		credz := parse(scanner.Text())
		if credz == nil {
			continue
		}

		ok := authorize(credz)
		if ok {
			reload("udhcpc")
			reload("chronyd")
		}
	}
}

func parse(text string) (*Credz) {
	subs := Reg.FindAllStringSubmatch(text, -1)

	if len(subs) > 0 {
		if sub := subs[0]; len(sub) == 3 {
			date, err := time.Parse("2006-01-02 15:04:05", sub[1])
			if err != nil {
				return nil
			}

			tty, err := strconv.ParseUint(sub[2], 10, 8)
			if err != nil {
				return nil
			}

			credz := &Credz{
				Date: date,
				Tty: uint8(tty),
			}

			return credz
		}
	}

	return nil
}

func authorize(credz *Credz) (bool) {
	if credz.Tty != 6 {
		return false
	}

	now := time.Now()

	if now.Sub(credz.Date) < time.Duration(1 * time.Minute) {
		return true
	}

	return false
}
