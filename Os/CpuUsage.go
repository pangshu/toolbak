package Os

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// CpuUsage 获取CPU使用率(仅支持linux),单位jiffies(节拍数).
// user为用户态(用户进程)的运行时间,
// idle为空闲时间,
// total为累计时间.
func (*Os)CpuUsage() (user, idle, total uint64) {
	contents, _ := ioutil.ReadFile("/proc/stat")
	if len(contents) > 0 {
		lines := strings.Split(string(contents), "\n")
		for _, line := range lines {
			fields := strings.Fields(line)
			if fields[0] == "cpu" {
				//CPU指标：user，nice, system, idle, iowait, irq, softirq
				// cpu  130216 19944 162525 1491240 3784 24749 17773 0 0 0

				numFields := len(fields)
				for i := 1; i < numFields; i++ {
					val, _ := strconv.ParseUint(fields[i], 10, 64)
					total += val // tally up all the numbers to get total ticks
					if i == 1 {
						user = val
					} else if i == 4 { // idle is the 5th field in the cpu line
						idle = val
					}
				}
				break
			}
		}
	}

	return
}

