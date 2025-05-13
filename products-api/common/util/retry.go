package util

import "time"

func InvokeWithRetry(fn func() error, waitTime time.Duration, retries int) error {
	var err error

	for i := 0; i < retries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(waitTime)
	}

	return err
}
