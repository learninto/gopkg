## MapReduce

## 案例：

```
package main

import (
	"fmt"
	"time"

	"github.com/learninto/gopkg/mr"
)

func main() {
	// 抢票结果
	var resp string

	// 抢票：携程
	primary := func(cancel func(error)) {
		time.Sleep(time.Second * 2) //睡 2 秒
		resp += "通过 携程 抢到票了"
		cancel(nil)
	}

	//抢票：12306
	secondary := func(cancel func(error)) {
		time.Sleep(time.Millisecond * 1500) // 睡 1500 毫秒
		resp += "通过 12306 抢到票了"
		cancel(nil)
	}

	// 放入 MapReduce 同时开抢
	_ = mr.MapReduceVoid(func(source chan<- interface{}) {
		source <- primary
		source <- secondary
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		fn := item.(func(func(error)))
		fn(cancel)
	}, func(pipe <-chan interface{}, cancel func(error)) {
		for item := range pipe {
			resp, _ = item.(string)
		}
	})

	// 输出结果：通过 12306 抢到票了 ----
	fmt.Println(resp, "----")
}
```