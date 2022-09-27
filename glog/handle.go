package glog

import "log"

func HandleDefaultErr(err error) {
	HandleError(err, nil)
}
func HandleError(err error, action func()) {
	if err != nil {
		if action != nil {
			action()
		} else {
			log.Println(err)
		}
	}
}
func HandleErrorF(err error, format string, v ...interface{}) {
	if err != nil {
		log.Printf(format, err, v)
	}
}

func HandleFatalF(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
