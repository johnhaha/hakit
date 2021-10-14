package hamsg

import "github.com/fatih/color"

func PrintRed(msg string) {
	color.Red(msg)
}

func PrintGreen(msg string) {
	color.Green(msg)
}

func InRed(msg string) string {
	red := color.New(color.FgRed).SprintFunc()
	block := red(msg)
	return block
}

func InGreen(msg string) string {
	green := color.New(color.FgGreen).SprintFunc()
	block := green(msg)
	return block
}
