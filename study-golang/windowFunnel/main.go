package windowFunnel

// Event
// index 1=A事件 ，2=B事件 以此类推
type Event struct {
	time  int64
	index int64
}

type EventsTimestamp struct {
	time  int64
	index int
}

func (e *EventsTimestamp) hasValue() bool {
	if e.time > 0 {
		return true
	} else {
		return false
	}
}

// 仿 clickhouse 实现
func test(data []Event, len int64, window int64) []EventsTimestamp {
	var eventsTimestamp = make([]EventsTimestamp, len)
	var currentArrayIndex int64 = 0

	for index, event := range data {

		timestamp := event.time
		eventIdx := event.index - 1

		if eventIdx == currentArrayIndex {
			eventsTimestamp[eventIdx].time = timestamp
			eventsTimestamp[eventIdx].index = index + 1
		} else if eventIdx == (currentArrayIndex + 1) {
			if eventsTimestamp[eventIdx-1].hasValue() {
				firstTime := eventsTimestamp[0].time
				if timestamp <= firstTime+window {
					eventsTimestamp[eventIdx].time = timestamp
					eventsTimestamp[eventIdx].index = index + 1
					currentArrayIndex += 1
					if eventIdx+1 == len {
						return eventsTimestamp
					}
				}
			}
		}
	}
	return eventsTimestamp
}

func findEventArray(data []Event, eventIndex int64) []int {
	var result = make([]int, 0)
	for i, event := range data {
		if event.index == eventIndex {
			result = append(result, i)
		}
	}
	return result
}

// 自己实现的新算法
// data 是已经排序好的 按时间升序的数组，length 是查找序列的长度，window 是窗口期
func test2(data []Event, length int64, window int64) []EventsTimestamp {
	for i := length; i > 0; i-- {
		eventIndexArray := findEventArray(data, i)
		if len(eventIndexArray) == 0 {
			continue
		}

		for _, eventIndex := range eventIndexArray {
			var eventsTimestamp = make([]EventsTimestamp, i)
			lastEvent := data[eventIndex]
			eventsTimestamp[i-1].time = lastEvent.time
			eventsTimestamp[i-1].index = eventIndex + 1

			if i-1 == 0 {
				return eventsTimestamp
			}

			var currentEventIndex = i - 1
			sliceData := data[:eventIndex]
			for s := eventIndex - 1; s >= 0; s-- {
				event := sliceData[s]
				var timeMatch = window >= eventsTimestamp[i-1].time-event.time
				if event.index == currentEventIndex && timeMatch {
					eventsTimestamp[currentEventIndex-1].time = event.time
					eventsTimestamp[currentEventIndex-1].index = s + 1
					currentEventIndex = currentEventIndex - 1
					if currentEventIndex < 1 {
						return eventsTimestamp
					}
				}
			}
		}
	}
	return make([]EventsTimestamp, 0)
}
