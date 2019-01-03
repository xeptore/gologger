package gologger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func log(title, message string, foreground colorKind) {
	fg := itoa(foreground.Light)
	logTitle(title, fg)
	logMessage(message)
	endl()
}

func logv(v interface{}) {
	color := config.Colors["blue"].Light
	logTitle("variable", itoa(color))
	logVariable(v)
}

func logTitle(title, fg string) {
	bold := itoa(config.Styles["bold"])
	title = "[" + strings.ToUpper(title) + "]  "
	fmt.Printf(fgStylePlaceholder, bold, fg, title)
}

func logMessage(msg string) {
	fmt.Printf(fgPlaceholder, itoa(config.ColorReset), msg)
}

func logVariable(v interface{}) {
	fmt.Printf("%+v", v)
	endl()
}

func endl() {
	fmt.Println()
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func exit(code int) {
	os.Exit(code)
}
