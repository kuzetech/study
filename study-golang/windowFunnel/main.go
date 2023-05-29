package windowFunnel

// Event
// index 1=A事件 ，2=B事件 以此类推
type Event struct {
	time int64
	step int64
}

type EventsTimestamp struct {
	time        int64
	sourceIndex int
}

func findSpecStepEventIndex(data []Event, step int64) []int {
	var result = make([]int, 0)
	for i, event := range data {
		if event.step == step {
			result = append(result, i)
		}
	}
	return result
}

// 自己实现的新算法
// data 是已经排序好的按时间升序的数组，length 是查找序列的长度，window 是窗口期
func findFunnel(data []Event, maxStep int64, window int64) []EventsTimestamp {
	for lastStep := maxStep; lastStep > 0; lastStep-- {
		lastStepSourceIndexArray := findSpecStepEventIndex(data, lastStep)
		if len(lastStepSourceIndexArray) == 0 {
			continue
		}

		for _, lastStepSourceIndex := range lastStepSourceIndexArray {
			var eventsTimestamp = make([]EventsTimestamp, lastStep+1)
			lastStepEvent := data[lastStepSourceIndex]
			eventsTimestamp[lastStep].time = lastStepEvent.time
			// 为了方便数位置，故意 +1
			eventsTimestamp[lastStep].sourceIndex = lastStepSourceIndex + 1

			if lastStep == 1 {
				return eventsTimestamp[1:]
			}

			var previousStep = lastStep - 1
			sliceData := data[:lastStepSourceIndex]

			for sourceIndex := lastStepSourceIndex - 1; sourceIndex >= 0; sourceIndex-- {
				event := sliceData[sourceIndex]
				var timeMatch = window >= eventsTimestamp[lastStep].time-event.time
				if event.step == previousStep && timeMatch {
					eventsTimestamp[previousStep].time = event.time
					// 为了方便数位置，故意 +1
					eventsTimestamp[previousStep].sourceIndex = sourceIndex + 1
					previousStep = previousStep - 1
					if previousStep < 1 {
						return eventsTimestamp[1:]
					}
				}
			}
		}
	}
	return make([]EventsTimestamp, 0)
}

func (e *EventsTimestamp) hasValue() bool {
	if e.time > 0 {
		return true
	} else {
		return false
	}
}

// 仿 clickhouse 实现
func findFunnelClickHouse(data []Event, len int64, window int64) []EventsTimestamp {
	var eventsTimestamp = make([]EventsTimestamp, len)
	var currentArrayIndex int64 = 0

	for index, event := range data {

		timestamp := event.time
		eventIdx := event.step - 1

		if eventIdx == currentArrayIndex {
			eventsTimestamp[eventIdx].time = timestamp
			eventsTimestamp[eventIdx].sourceIndex = index + 1
		} else if eventIdx == (currentArrayIndex + 1) {
			if eventsTimestamp[eventIdx-1].hasValue() {
				firstTime := eventsTimestamp[0].time
				if timestamp <= firstTime+window {
					eventsTimestamp[eventIdx].time = timestamp
					eventsTimestamp[eventIdx].sourceIndex = index + 1
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
