package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readConfigFile() (*[]byte, error) {
	bytes, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return &[]byte{}, err
	}
	return &bytes, nil
}

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
	fmt.Printf(fgPlaceholder, colorReset, msg)
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
