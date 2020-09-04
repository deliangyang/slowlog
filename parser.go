package slowlog

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// SlowLog Redis日志结构
type SlowLog struct {
	ID          int64     `json:"id"`
	CreateTime  time.Time `json:"createTime"`
	ExecuteTime int64     `json:"executeTime"`
	Operator    string    `json:"operator"`
	Parameters  []string  `json:"parameters"`
}

func flagIndex(line string, substr string) string {
	index := strings.Index(line, substr)
	if index < 0 {
		return ""
	}
	index += len(substr)
	return line[int64(index):]
}

// Parse 解析日志
func Parse(log string) ([]SlowLog, error) {
	lines := strings.Split(log, "\n")
	var slowLogs []SlowLog
	var slowLog *SlowLog

	for _, line := range lines {
		line := strings.TrimSpace(line)

		item := flagIndex(line, "1) (integer) ")

		if item != "" {
			id, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				return slowLogs, err
			}
			if slowLog != nil {
				slowLogs = append(slowLogs, *slowLog)
			}
			slowLog = &SlowLog{
				ID: id,
			}
		}

		if slowLog == nil {
			return slowLogs, errors.New("slow logs not init")
		}

		item = flagIndex(line, "2) (integer) ")
		if item != "" {
			cTime, err := strconv.ParseInt(item, 10, 64)
			if err == nil {
				slowLog.CreateTime = time.Unix(cTime, 0)
			}
		}

		item = flagIndex(line, "3) (integer) ")
		if item != "" {
			executeTime, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				return slowLogs, nil
			}
			slowLog.ExecuteTime = executeTime
		}

		item = flagIndex(line, "4) 1) \"")
		if item != "" {
			slowLog.Operator = strings.Trim(item, "\"")
		} else {
			item = flagIndex(line, ") \"")
			if item != "" {
				slowLog.Parameters = append(
					slowLog.Parameters,
					strings.Trim(item, "\""))
			}
		}
	}

	if slowLog != nil {
		slowLogs = append(slowLogs, *slowLog)
	}
	return slowLogs, nil
}
