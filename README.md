# gologger
Simple `fmt.Println` replacement with _<span style="color:#ff0000;">c</span><span style="color:#ff7f00;">o</span><span style="color:#ffff00;">l</span><span style="color:#00ff00;">o</span><span style="color:#00ffff;">r</span><span style="color:#0000ff;">s</span>_ actually!

## Installation
run:
```bash
go get -u github.com/xeptore/gologger
```

## Usage
first of all `import` package:
```go
import (
	"github.com/xeptore/gologger"
)
```
### APIs
*
	```go
	Log(msg string)
	```
	logs `msg` to output  with a _yellow_ title color.  
	  
	example:
	```go
	gologger.Log(time.Now().String())
	```
	output:  
	_an image_

-------
*
	```go
	Logv(v interface{})
	```
	like `.Log`, useful for logging any type of variable.  
	  
	example:
	```go
	v := map[string]int{
		"a key": 1,
		"another key": 2,
	}
	gologger.Logv(v)
	```
	output:  
	_an image_

------
*
	```go
	Error(msg string)
	```
	logs `msg` to output  with a _red_ title color.  
	  
	example:
	```go
	gologger.Error("could not find configuration file.")
	```
	output:  
	_an image_

---
*
	```go
	Warn(msg string)
	```
	logs `msg` to output  with a _orange_ title color.  
	  
	example:
	```go
	gologger.Warn("a warning message")
	```
	output:  
	_an image_

----
*
	```go
	Success(msg string)
	```
	logs `msg` to output  with a _green_ title color.  
	  
	example:
	```go
	gologger.Success("successfully connected to database.")
	```
	output:  
	_an image_

---

### Notes
 There are also _`Fatal`-suffixed_ functions provided useful in cases that you want to finish executing your program after logging (like ```log.Fatalln```).
 For example:
```go
gologger.ErrorFatal("could not find configuration file.")
```
Finishes executing current process after printing message.
following functions can be used to achive this:

*
	```go
	LogFatal(msg string)
	```
*
	```go
	LogvFatal(msg string)
	```
*
	```go
	ErrorFatal(msg string)
	```
*
	```go
	WarnFatal(msg string)
	```

### Config
You can change colors by changing the color ASCII codes in `config.go` file and rebuilding the packge.
ASCII color codes reference _(like almost anything else!)_ can be found in [Wikipedia](https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit).
