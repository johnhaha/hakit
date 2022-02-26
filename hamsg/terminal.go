package hamsg

import "github.com/fatih/color"

func PrintRed(msg string) {
	color.Red(msg)
}

func PrintGreen(msg string) {
	color.Green(msg)
}

func PrintBlue(msg string) {
	color.Blue(msg)
}

func PrintYellow(msg string) {
	color.Yellow(msg)
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

func InBlue(msg string) string {
	blue := color.New(color.FgBlue).SprintFunc()
	block := blue(msg)
	return block
}

func InYellow(msg string) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	block := yellow(msg)
	return block
}
