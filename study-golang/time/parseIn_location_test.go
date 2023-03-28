package time

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_parseIn_location(t *testing.T) {
	assertions := require.New(t)

	layout := "2006-01-02T15:04:05Z07:00"

	// 东八区 = 1679648520
	// UTC =   1679677320
	specifiedTimeZoneStr := "2023-03-24T17:02:00Z08:00"

	// location, err := time.LoadLocation("UTC")
	location, err := time.LoadLocation("Local")
	assertions.Nil(err)

	// 如果指定了时区，会直接忽略数据中的时区
	result, err := time.ParseInLocation(layout, specifiedTimeZoneStr, location)
	assertions.Nil(err)

	assertions.Equal(int64(1679648520), result.Unix())

}
