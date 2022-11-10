package fileCreateTime

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func timeSpecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(ts.Sec, ts.Nsec)
}

func test() {
	path := "/Users/huangsw/code/funny/turbine/sources/ingest-client/input/a/1.log"

	state, _ := os.Stat(path)
	fileAttr := state.Sys().(*syscall.Stat_t)

	fmt.Println(timeSpecToTime(fileAttr.Atimespec))
	fmt.Println(timeSpecToTime(fileAttr.Ctimespec))
	fmt.Println(timeSpecToTime(fileAttr.Mtimespec))
}
