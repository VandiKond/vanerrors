package vanerrors

import (
	"fmt"
	"log"
)

// Creates a log of VanError based on it settings
//
// The method could be used inside some methods (New(), VanError.Error()) and outside
// err := Default(Name, Message, Code, Logger)
// err.Log()
//
// It is a basic logger, using the standard log package
// If you want to have better log, set logs of.
func (e VanError) Log() {
	// Setting the logger output
	options := e.LoggerOptions
	var result string

	// Adding severity
	if options.ShowSeverity {
		if options.IntSeverity {
			result += "level: " + SeverityArray[e.Severity] + ","
		} else {
			result += fmt.Sprintf("level: %d,", e.Severity)
		}
	}

	// Adding code
	if options.ShowCode {
		result += fmt.Sprintf(" %d", e.Code)
	}

	// Adding name
	result += " " + e.Name

	// Adding message
	if options.ShowMessage {
		result += ": " + e.Message
	}

	// Adding , to show the next data
	result += ", "

	// Adding description
	if options.ShowDescription && e.Description != nil {
		description := make([]byte, 4096)
		n, err := e.Description.Read(description)
		if err == nil {
			description = description[:n]
		}

		result += "description: " + string(description) + ", "
	}

	// Adding cause
	if options.ShowCause && e.Cause != nil {
		result += "cause: " + e.Cause.Error()
	}

	// Getting the result string
	logger := log.New(e.logger, "", log.LstdFlags|log.Llongfile|log.Lshortfile)
	if e.Severity == 3 {
		logger.Fatalln(result)
	}
	logger.Println(result)
}
