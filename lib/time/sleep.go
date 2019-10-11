package LibTime

import (
	"fmt"
	"math"
	"time"
)

func main() {
	right_now := time.Now()
	timestamp := right_now.UnixNano()

	fmt.Println(right_now)
	fmt.Println(right_now.Unix())
	fmt.Println(timestamp)
	fmt.Println(math.Ceil(float64(timestamp/1000000)) * 1000)
}
