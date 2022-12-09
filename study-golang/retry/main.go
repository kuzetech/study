package retry

import (
	"github.com/avast/retry-go/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

func test_retry() {

	var fc = func() error {
		_, err := http.Get("http://asdasd")
		if err != nil {
			return err
		}
		return nil
	}

	var callback = func(n uint, err error) {
		log.Printf("#%d: %s \n", n+1, err)
	}

	// 返回 true 就重试
	var retryIf = func(err error) bool {
		if strings.Index(err.Error(), "unsupported protocol scheme") > -1 {
			return false
		}
		if strings.Index(err.Error(), "no such host") > -1 {
			return false
		}
		return true
	}

	err := retry.Do(
		fc,
		retry.Delay(time.Second*3),
		retry.MaxDelay(time.Second*3),
		retry.Attempts(3),
		retry.OnRetry(callback),
		retry.RetryIf(retryIf),
	)

	if err != nil {
		log.Println(err)
	}
}
