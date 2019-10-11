/*Package log - to create the logging.
Currently just printing in two messages( Info and Error) but this can be categorized for Info, warning and error further
*/
package log

import (
	"log"
	"os"
	"strconv"
	"time"
)

// This can be further enhanced to set in start itself. This file can be created anywhere else too
var fileName = "/tmp/info.log"

//Info - To log Info messages
func Info(msg ...interface{}) {
	message("Info", msg)
}

// Error - To Log error messages
func Error(msg ...interface{}) {
	message("Error", msg)
}

// function which is actually logging
func message(msg ...interface{}) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print(makeTimestamp(), msg)
}

func makeTimestamp() string {
	a := time.Now().UnixNano() / int64(time.Millisecond)
	return strconv.FormatInt(a, 10)
}
