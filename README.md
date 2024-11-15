# VanErrors

## What is vanerrors?

A simple package for working with errors
It has:

- Creating custom errors
- Custom display settings
- Optional simple logger
- Custom logger settings
- Methods that allow using VanError as an error and easer compare with them using errors package

## How to install the package

It's simple. 
Just Write this into the console with your project.
```bash
go get github.com/VandiKond/vanerrors
```


## How to use?

So, how to use your vanerrors?
It is simple. you can create a error only with a name 

### Start with name
```go
err := vanerrors.NewName("readme error", nil) // you can ad a io.Writer instead of nil, it will be your logger
fmt.Println(err.Error()) // "readme error"
```

If you want to add more information, for example message, you can do like this:

### Basic parameters  
```go
err := vanerrors.NewBasic("readme error", "here could be the error message", nil) // nil is here also a empty logger
fmt.Println(err.Error()) // "readme error: here could be the error message"
```

You can also use an http format

### HTTP parameters 
```go
err := vanerrors.NewHTTP("readme error", 500, nil) // nil is here also a empty logger
fmt.Println(err.Error()) // "500 readme error"
```

Or you can add a cause error

### Cause parameter 
```go
err := vanerrors.NewHTTP("readme error", errors.New("some cause"), nil) // nil is here also a empty logger
fmt.Println(err.Error()) // "readme error" (the cause will be in the error, but not shown)
```

Do you want a more advanced method? Here:

### ErrorData
```go
err := vanerrors.NewDefault(vanerrors.ErrorData{
    Name: "readme error", // The error name
    Message: "here could be the error message", // The error message
    Code: 500, // The error code
    Cause: errors.New("some cause error"), // The cause error (it could be nil)
    Description:  bytes.NewReader([]byte("you can add more information here. the more - the better")), // The error description (as io.Reader)
    Logger: nil, // The error logger (io.Writer)
    Severity: 2, // The error severity. 1: warn, 2: error, 3: fatal (panic)
})
fmt.Println(err.Error()) // "500 readme error: here could be the error message" (other information wasn't shown because of the show settings)
```

So in this example for some reason not all data was printed.
Why? We've added it to the error settings, so why it is not printed

It is since the display options by default display only code, name, message. Other data is still in the error. You can edit this settings

### Display Options
```go
err := vanerrors.New(vanerrors.ErrorData{
    Name: "readme error",
    Message: "here could be the error message", 
    Code: 500, 
    Cause: errors.New("some cause error"),
    Description:  bytes.NewReader([]byte("you can add more information here. the more - the better")),
    Logger: nil,
    Severity: 2,
    },
    vanerrors.Options{
        ShowMessage: true, // Do you need to show error message
        ShowCode: true, // Do you need to show error code
        ShowSeverity: true, // Do you need to show severity
        IntSeverity: true, // Do you need to show severity as a string or an int (false: string, true: int)
        ShowDescription: true, // Do you need to show description 
        ShowCause: true, // Do you need to show error case
        ShowDate: true, // Do you need to show at which date the error was created
    }, vanerrors.EmptyLoggerOptions // Empty log settings
)
fmt.Println(err.Error()) // "2006-01-02 15:04:05 level: 2, 500 readme error: here could be the error message, description: you can add more information here. the more - the better, cause: some cause error"
```

If some value is false you can skip it, because go automatically sets bool parameters to false. 
If you skip one of the values it would be marked as false

So I've mentioned log settings. How to use it?:

### LoggerOptions
```go
err := vanerrors.New(vanerrors.ErrorData{
    Name: "readme error",
    Message: "here could be the error message", 
    Code: 500, 
    Cause: errors.New("some cause error"),
    Description:  bytes.NewReader([]byte("you can add more information here. the more - the better")),
    Logger: nil,
    Severity: 2,
    },
    vanerrors.Options{
        ShowMessage: true, 
        ShowCode: true, 
    }, vanerrors.LoggerOptions{
        DoLog: true, // Do the program need to log data automatically
        ShowMessage: true, // Do you need to show error message
        ShowCode: true, // Do you need to show error code
        ShowSeverity: true, // Do you need to show severity
        IntSeverity: true, // Do you need to show severity as a string or an int (false: string, true: int)
        ShowDescription: true, // Do you need to show description 
        ShowCause: true, // Do you need to show error case
        LogBy: true, // Do you need to log when the error is created or when it is called (false: created, true: called)
    }
)
err.Error() // "2006-01-02 15:04:05 level: 2, 500 readme error: here could be the error message, description: you can add more information here. the more - the better, cause: some cause error" 
```

So now you can say that it is to complicated? Yes it is!

This is why i've added 3 variables for default recommended values. 
One i've already shown you: it is empty logger options

### EmptyLoggerOptions
```go
opt := vanerrors.EmptyLoggerOptions
opt.ShowMessage = true // Enable the message
```
Yea it's just 
```go
var EmptyLoggerOptions LoggerOptions = LoggerOptions{}
```
However who cares?

The next one is also for logger options, however it has a lot of enabled values, that i recommend to leave enabled

### DefaultLoggerOptions 
```go
opt := vanerrors.DefaultLoggerOptions
opt.ShowCause = false // Disable the cause
```
So it is more complicated:
```go
var DefaultLoggerOptions LoggerOptions = LoggerOptions{
	DoLog:           true,
	ShowMessage:     true,
	ShowCode:        true,
	ShowSeverity:    true,
	ShowDescription: true,
	ShowCause:       true,
}
```

And the last one. It's the same but for options

### DefaultOptions
```go
opt := vanerrors.DefaultOptions
opt.ShowSeverity = true // Enable the severity
```
It has only two default values
```go
var DefaultOptions Options = Options{
	ShowMessage: true,
	ShowCode:    true,
}
```

Okay. So we know how to create an error, edit settings.
Lets get  the information about it!

### Error Information 
```go
err := vanerrors.NewBasic("readme error", "here could be the error message", nil)
errorText := err.Error()
```
So actually it just shows the error

And what about logs? We can make the program show the log by creating and calling the error.
Bu what about Logging the error when you want?
### Log error
```go
err := vanerrors.NewBasic("readme error", "here could be the error message", nil)
err.Log()  // "2006-01-02 15:04:05 level: 2, 500 readme error: here could be the error message" (other info not shown because it is empty)
```
So now you can log the error when you want

Now about other methods:

### Methods for other interfaces (for instance errors package)
```go
err := vanerrors.NewBasic("readme error", "here could be the error message", nil)

// As
var targetErr error
err.As(&targetErr) // Target error will be the same as err

// Unwrap
err.Cause = errors.New("some cause")
fmt.Println(err.Unwrap) // "some cause"

// UnwrapAll
err2 := vanerrors.NewBasic("readme error", "here could be the error message", nil)
err2.Cause = err
fmt.Println(err2.UnwrapAll()) // {"readme error: here could be the error message", "some cause"}

// Is
err3 := vanerrors.NewBasic("readme error", "here could be the error message", nil)
err4 := vanerrors.NewBasic("readme error", "here could be the error message", nil)
fmt.Println(err3.Is(err4)) // true
err5 := vanerrors.NewBasic("other readme error", "here could be the error message", nil)
fmt.Println(err5.Is(err4)) // false

// Get
func GetError() error {
    return  vanerrors.NewBasic("readme error", "here could be the error message", nil)
} 

func GetOtherError() error {
    return errors.New("some error")
}

fmt.Println(Get(GetError())) // "readme error: here could be the error message"
fmt.Println(Get(GetOtherError)) // nil
```
Now you can do more operations with the error

Now, how to get a value from error if you don't know is it a van error and don't wan't to yse Get

### Getters
```go
func GetError() error {
    return vanerrors.NewDefault(vanerrors.ErrorData{
        Name: "readme error", 
        Message: "here could be the error message", 
        Code: 500, 
        Cause: errors.New("some cause error"),
        Description:  bytes.NewReader([]byte("you can add more information here. the more - the better")),
        Logger: nil, 
        Severity: 2,
    })

    err := GetError()
    err2 := errors.New("not vandi error")

    // GetName
    GetName(err) // "readme error"
    GetName(err2) // ""

    // GetMessage
    GetMessage(err) // "here could be the error message"
    GetMessage(err2) // ""

    // GetCode
    GetCode(err) // 500
    GetCode(err2) // 0

    // GetSeverityStr
    GetSeverityStr(err) // "error"
    GetSeverityStr(err2) // ""

    // GetSeverityInt
    GetSeverityInt(err) // 2
    GetSeverityInt(err2) // 0

    // GetDescription
    GetDescription(err) // bytes.Reader
    GetDescription(err2) // nil

    // GetDate
    GetDate(err) // time.Now()
    GetDate(err2) // nil
}
```
Use it to get special data of the error


## Other information 

### License

[MIT](LICENSE)