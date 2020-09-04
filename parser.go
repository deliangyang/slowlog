package slowlog

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	Flag1 = "1) (integer) "
	Flag2 = "2) (integer) "
	Flag3 = "3) (integer) "
	Flag4 = "4) 1) \""
	Flag5 = ") \""
)

// SlowLog Redis日志结构
type SlowLog struct {
	ID          int64     `json:"id"`
	CreateTime  time.Time `json:"createTime"`
	ExecuteTime int64     `json:"executeTime"`
	Operator    string    `json:"operator"`
	Parameters  []string  `json:"parameters"`
}

// Parse 解析日志
func Parse(log string) ([]SlowLog, error) {
	lines := strings.Split(log, "\n")
	var slowLogs []SlowLog
	var slowLog *SlowLog

	for _, line := range lines {
		line := strings.TrimSpace(line)

		if item := flagIndex(line, Flag1); item != "" {
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

		if item := flagIndex(line, Flag4); item != "" {
			slowLog.Operator = trim(item)
		} else {
			if item := flagIndex(line, Flag5); item != "" {
				slowLog.Parameters = append(slowLog.Parameters, trim(item))
			}
		}

		if item := flagIndex(line, Flag2); item != "" {
			cTime, err := strconv.ParseInt(item, 10, 64)
			if err == nil {
				slowLog.CreateTime = time.Unix(cTime, 0)
			}
		}

		if item := flagIndex(line, Flag3); item != "" {
			executeTime, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				return slowLogs, nil
			}
			slowLog.ExecuteTime = executeTime
		}
	}

	if slowLog != nil {
		slowLogs = append(slowLogs, *slowLog)
	}
	return slowLogs, nil
}

func trim(line string) string {
	return strings.Trim(line, "\"")
}

func flagIndex(line string, substr string) string {
	index := strings.Index(line, substr)
	if index < 0 {
		return ""
	}
	index += len(substr)
	return line[int64(index):]
}
