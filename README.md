
## Redis Query SlowLog Result Parse

```bash
redis-cli slowlog get
```


### Example

```go
package main

import (
	"fmt"
	"log"
	"slowlog"
)

func main() {
	slowLogs, err := slowlog.Parse(`1) 1) (integer) 387
    2) (integer) 1598574895
    3) (integer) 15104
    4) 1) "DEL"
       2) "x:c:25820156"
       3) "x:u:p:md:25820156"
       4) "x:r:m:25820156"
       5) "x:r:m_i:25820156"
       6) "x:im:l:25820156"
       7) "x:u:25820156:60ccd4bc-1256-4eac-af3e-f4db725976c0"
2) 1) (integer) 386
    2) (integer) 1598486808
    3) (integer) 23541
    4) 1) "DEL"
       2) "x:c:25752412"
       3) "x:u:p:md:25752412"
       4) "x:r:m:25752412"
       5) "x:r:m_i:25752412"
       6) "x:im:l:25752412"
       7) "x:u:25752412:44b9ca8e-4bd8-4c6c-a48c-c945bf153c3e"`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(slowLogs[0].ID)
}
```

