package multiWriterLog

import "time"

func printLog() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			Logger.Info().Msg("1111111")
		}
	}
}
