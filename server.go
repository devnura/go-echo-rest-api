package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")

	logrus.Info("choose environment " + ENV)
}

func main() {
	var (
	// APP  = os.Getenv("APP")
	// ENV  = os.Getenv("ENV")
	// PORT = os.Getenv("PORT")
	// NAME = fmt.Sprintf("%s-%s", APP, ENV)
	)

}
