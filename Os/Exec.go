package Os

import (
	"bytes"
	"os/exec"
	"strings"
	"unicode"
)

// Exec 执行一个外部命令.
// retInt为1时失败,为0时成功;outStr为执行命令的输出;errStr为错误输出.
// 命令如
// "ls -a"
// "/bin/bash -c \"ls -a\""
func (*Os)Exec(command string) (retInt int, outStr, errStr []byte) {
	// split command
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})

	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		retInt = 1 //失败
		stderr.WriteString(err.Error())
		errStr = stderr.Bytes()
	} else {
		retInt = 0 //成功
		outStr, errStr = stdout.Bytes(), stderr.Bytes()
	}

	return
}



//// System 与Exec相同,但会同时打印标准输出和标准错误.
//func System(command string) (retInt int, outStr, errStr []byte) {
//	// split command
//	q := rune(0)
//	parts := strings.FieldsFunc(command, func(r rune) bool {
//		switch {
//		case r == q:
//			q = rune(0)
//			return false
//		case q != rune(0):
//			return false
//		case unicode.In(r, unicode.Quotation_Mark):
//			q = r
//			return false
//		default:
//			return unicode.IsSpace(r)
//		}
//	})
//
//	// remove the " and ' on both sides
//	for i, v := range parts {
//		f, l := v[0], len(v)
//		if l >= 2 && (f == '"' || f == '\'') {
//			parts[i] = v[1 : l-1]
//		}
//	}
//
//	var stdout, stderr bytes.Buffer
//	var err error
//
//	cmd := exec.Command(parts[0], parts[1:]...)
//	stdoutIn, _ := cmd.StdoutPipe()
//	stderrIn, _ := cmd.StderrPipe()
//	outWr := io.MultiWriter(os.Stdout, &stdout)
//	errWr := io.MultiWriter(os.Stderr, &stderr)
//
//	err = cmd.Start()
//	if err != nil {
//		retInt = 1 //失败
//		stderr.WriteString(err.Error())
//		fmt.Printf("%s\n", stderr.Bytes())
//		return
//	}
//
//	go func() {
//		_, _ = io.Copy(outWr, stdoutIn)
//	}()
//	go func() {
//		_, _ = io.Copy(errWr, stderrIn)
//	}()
//
//	err = cmd.Wait()
//	if err != nil {
//		stderr.WriteString(err.Error())
//		fmt.Println(stderr.Bytes())
//		retInt = 1 //失败
//	} else {
//		retInt = 0 //成功
//	}
//	outStr, errStr = stdout.Bytes(), stderr.Bytes()
//
//	return
//}
