# guardpanic
recover goroutine panic and support goroutine restart.

# example
package main

import (
	"fmt"
	"github.com/kaimixu/guardpanic"
)

func main() {
	go func() {
		worker := func() {
			var s []int
			s[0] = 1 // panic occur
		}
		procErr := func(err error) {
			fmt.Println(err)
		}

		guardpanic.Run(
			worker, // 业务逻辑处理函数
			3, // 发生panic后重启goroutine，重启次数不超过3次，0：表示goroutine不重启
			procErr, // 处理panic日志
		)
	}()

	select{}
}

