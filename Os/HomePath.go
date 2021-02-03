package Os

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// 获取当前系统的主目录
func (toolOs *Os)HomePath() (string, error) {
	u, err := user.Current()
	if nil == err {
		return u.HomeDir, nil
	}

	if toolOs.IsWindows() {
		return homeWindows()
	} else {
		return homeUnix()
	}
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}


/*获取当前文件执行的路径*/
//func GetCurPath() string {
//	file, _ := exec.LookPath(os.Args[0])
//
//	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
//	path, _ := filepath.Abs(file)
//
//	rst := filepath.Dir(path)
//
//	return rst
//}
