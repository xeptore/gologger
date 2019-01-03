package gologger

const (
	escape             = "\x1b["
	reset              = escape + "0m"
	fgStylePlaceholder = escape + "%s;38;5;%sm%s" + reset
	fgPlaceholder      = escape + "38;5;%sm%s" + reset
)

func Log(msg string) {
	color := config.Colors["yellow"]
	log("Log", msg, color)
}

func LogFatal(msg string) {
	Log(msg)
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
	Error(msg)
	endl()
	exit(2)
}

func Success(msg string) {
	color := config.Colors["green"]
	log("Ok", msg, color)
}

func LogvFatal(v interface{}) {
	Logv(v)
	endl()
	exit(0)
}

func Logv(v interface{}) {
	logv(v)
}
