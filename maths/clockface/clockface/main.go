package main

import (
	"os"
	"time"

	"github.com/biggsean/learn-go-with-tests2/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
