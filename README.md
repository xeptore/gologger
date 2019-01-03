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

    ![screenshot from 2019-01-03 09-47-41](https://user-images.githubusercontent.com/29199390/50625353-dc4d0400-0f3c-11e9-91f1-3b8ade9af2ce.png)


-------
*
    ```go
    Logv(v interface{})
    ```
    logs `msg` to output  with a _blue_ title color. useful for logging any type of variable.  
      
    example:
    ```go
    v := map[string]int{
        "a key": 1,
        "another key": 2,
    }
    gologger.Logv(v)
    ```
    output:  
    
    ![screenshot from 2019-01-03 09-48-02](https://user-images.githubusercontent.com/29199390/50625344-cf301500-0f3c-11e9-9fd5-7498ab5ebce6.png)

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
    
    ![screenshot from 2019-01-03 09-47-16](https://user-images.githubusercontent.com/29199390/50625336-c63f4380-0f3c-11e9-9263-4828ca16e151.png)

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
    
    ![screenshot from 2019-01-03 09-46-55](https://user-images.githubusercontent.com/29199390/50625327-b889be00-0f3c-11e9-9ab1-0cb07c8d85c9.png)

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
    
    ![screenshot from 2019-01-03 09-40-20](https://user-images.githubusercontent.com/29199390/50625273-40bb9380-0f3c-11e9-9aa8-38d04bd7a5b7.png)

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
