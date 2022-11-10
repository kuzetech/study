package time

import (
	"fmt"
	"time"
)

func test() {
	point := time.Now().Add(-time.Hour)
	fmt.Println(point)
}
