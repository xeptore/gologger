package gologger

import (
	yaml "gopkg.in/yaml.v2"
)

const (
	escape             = "\x1b["
	reset              = escape + "0m"
	fgStylePlaceholder = escape + "%s;38;5;%sm%s" + reset
	fgPlaceholder      = escape + "38;5;%sm%s" + reset
)

var (
	config     configType
	styleReset string
	colorReset string
	bgReset    string
)

func main() {
	Build()
}

func Build() error {
	bytes, err := readConfigFile()
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(*bytes, &config)
	if err != nil {
		return err
	}

	styleReset = itoa(config.Styles["reset"])
	colorReset = itoa(config.ColorReset)

	return nil
}

func Log(msg string) {
	color := config.Colors["yellow"]
	log("Log", msg, color)
}

func LogFatal(msg string) {
	endl()
	Log(msg)
	endl()
	exit(0)
}

func Warn(msg string) {
	color := config.Colors["orange"]
	log("Warning", msg, color)
}

func WarnFatal(msg string) {
	endl()
	Warn(msg)
	endl()
	exit(1)
}

func Error(msg string) {
	color := config.Colors["red"]
	log("Error", msg, color)
}

func ErrorFatal(msg string) {
	endl()
	Error(msg)
	endl()
	exit(2)
}

func Success(msg string) {
	color := config.Colors["green"]
	log("Ok", msg, color)
}

func LogFatalv(v interface{}) {
	endl()
	Logv(v)
	endl()
	exit(0)
}

func Logv(v interface{}) {
	logv(v)
}
