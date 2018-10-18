package debug

import "log"

func Log(msg interface{}) {
	log.Println(msg)
}

func Logf(msg string, v ...interface{}) {
	log.Printf(msg+"\n", v...)
}

func Fatal(msg interface{}) {
	log.Fatal(msg)
}
