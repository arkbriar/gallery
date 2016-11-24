package main

import (
	"github.com/Sirupsen/logrus"
)

func main() {
	s, err := NewGalleryServer()
	if err != nil {
		logrus.Panicln(err)
	}
	logrus.Panicln(s.Start())
}
