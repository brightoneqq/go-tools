package flow

import "time"

func Retry(maxRetry, intervalSec int, process func() error) error {
	var err error
	for i := 0; i < maxRetry; i++ {
		err = process()
		if err != nil {
			multipleTime := i + 1
			time.Sleep(time.Duration(intervalSec*multipleTime) * time.Second)
		} else {
			//success case
			break
		}
	}
	return err

}
