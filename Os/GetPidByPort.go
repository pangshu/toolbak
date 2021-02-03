package Os

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//GetPidByPort 根据端口号获取监听的进程PID.
func (*Os)GetPidByPort(port int) (pid int) {
	files := []string{
		"/proc/net/tcp",
		"/proc/net/udp",
		"/proc/net/tcp6",
		"/proc/net/udp6",
	}

	procDirs, _ := filepath.Glob("/proc/[0-9]*/fd/[0-9]*")
	for _, fpath := range files {
		lines, _ := readInArray(fpath)
		for _, line := range lines[1:] {
			fields := strings.Fields(line)
			if len(fields) < 10 {
				continue
			}

			//非 LISTEN 监听状态
			if fields[3] != "0A" {
				continue
			}

			//本地ip和端口
			ipport := strings.Split(fields[1], ":")
			locPort, _ := hexToDec(ipport[1])

			// 非该端口
			if int(locPort) != port {
				continue
			}

			pid = getPidByInode(fields[9], procDirs)
			if pid > 0 {
				return
			}
		}
	}

	return
}

// getPidByInode 根据套接字的inode获取PID.须root权限.
func getPidByInode(inode string, procDirs []string) (pid int) {
	if len(procDirs) == 0 {
		procDirs, _ = filepath.Glob("/proc/[0-9]*/fd/[0-9]*")
	}

	re := regexp.MustCompile(inode)
	for _, item := range procDirs {
		path, _ := os.Readlink(item)
		out := re.FindString(path)
		if len(out) != 0 {
			pid, _ = strconv.Atoi(strings.Split(item, "/")[2])
			break
		}
	}

	return pid
}

func readInArray(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

func hexToDec(str string) (int64, error) {
	start := 0
	if len(str) > 2 && str[0:2] == "0x" {
		start = 2
	}

	// bitSize 表示结果的位宽（包括符号位），0 表示最大位宽
	return strconv.ParseInt(str[start:], 16, 0)
}