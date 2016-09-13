package nssk

import "time"

var logs = []string{}

// AppendLog 扩写日志
func AppendLog(log string) {
	logs = append(logs, "["+time.Now().Format("2006/01-02 15:04:05")+"] "+log)
}

// PullLog 拉取日志
func PullLog() []string {
	logSize := len(logs)
	offset := logSize - 10
	if offset < 0 {
		offset = 0
	}
	return logs[offset:logSize]
}
