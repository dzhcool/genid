package snow

import (
	"fmt"
	"testing"
	"time"
)

func TestGenId(t *testing.T) {
	stime := time.Now()
	num := 10000000
	for i := 0; i < num; i++ {
		New().GenId()
	}
	etime := time.Now()
	uptime := etime.Sub(stime)

	t.Log(fmt.Sprintf("num:%d use:%v", num, uptime))
}
