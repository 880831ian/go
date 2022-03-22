package greeting

import (
	"fmt"

	"github.com/fatih/color"
)

func Say(s string) {
	fmt.Println(s)
}

func SayWithRed (s string) {
	color.Red(s)
}

func SayWithBlue (s string) {
	color.Blue(s)
}
func SayWithYellow (s string) {
	color.Yellow(s)
}