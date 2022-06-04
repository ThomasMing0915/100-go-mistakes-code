package main

import (
	"github.com/ThomasMing0915/100-go-mistakes-code/60/flight"
	"github.com/ThomasMing0915/100-go-mistakes-code/60/publish"
)

func main() {
	position := flight.NewPosition(22.1, 34.2, 10000)

	ph := publish.NewPublishHandler()

	ph.PublishPosition(position)

}
