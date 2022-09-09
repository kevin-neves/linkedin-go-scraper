package main

import (
	"log"

	"github.com/AllenDang/giu"
)

func onClickMe() {
	log.Println("Hello world!")
}

func onImSoCute() {
	log.Println("Im sooooooo cute!!")
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Label("Hello world from giu"),
		giu.Row(
			giu.Button("Click Me").OnClick(onClickMe),
			giu.Button("I'm so cute").OnClick(onImSoCute),
		),
	)
}

func main() {

}
