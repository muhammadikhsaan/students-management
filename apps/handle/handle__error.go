package handle

import (
	"io"
	"log"
	"os"
	"time"
)

var f *os.File

func init() {
	var err error
	f, err = os.Create("logger/logger_system-error.log")
	if err != nil {
		log.Fatalf("[SYSTEM-ERROR] %v | %v", time.Now(), err)
	}
}

//ErrorHandle used for handle error on application
func ErrorHandle(err error) {
	if err != nil {
		log.Fatalf("[SYSTEM-ERROR] %v | %v", time.Now(), err)
		wrt := io.MultiWriter(f)
		log.SetOutput(wrt)
	}
}
