package main

import (
	"log"
	"github.com/tehmoon/errors"
)

func reload(service string) {
	var err error

	switch service {
		case "chronyd":
			err = reloadChronyd()
		case "udhcpc":
			err = reloadUdhcpc()
		default:
			log.Printf("Service %s is not supported\n", service)
			return
	}

	if err != nil {
		log.Println(errors.Wrapf(err, "Error executing service: %s", service).Error())
	}

	return
}

func reloadChronyd() (error) {
	return execString("/sbin/service chronyd restart")
}

func reloadUdhcpc() (error) {
	return execString("/sbin/udhcpc -q -f")
}
