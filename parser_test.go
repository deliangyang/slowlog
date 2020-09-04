package slowlog

import (
	"strings"
	"testing"
)

var (
	testData = `1) 1) (integer) 387
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
       7) "x:u:25752412:44b9ca8e-4bd8-4c6c-a48c-c945bf153c3e"`
)

func TestParse(t *testing.T) {
	slowLogs, err := Parse(testData)
	if err != nil {
		t.Fatal(err)
	}

	if len(slowLogs) != 2 {
		t.Failed()
	}

	for _, slowLog := range slowLogs {
		if slowLog.Operator != "DEL" {
			t.Failed()
		}
		if len(slowLog.Parameters) != 7 {
			t.Failed()
		}

		if !strings.Contains(slowLog.Parameters[5], "x:u:25") {
			t.Failed()
		}
	}

	if slowLogs[0].ID != 387 || slowLogs[1].ID != 386 {
		t.Failed()
	}

}
