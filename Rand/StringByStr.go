package Rand

import (
	"math/rand"
	"time"
)

//随机几位字符串
func (*Rand)StringByStr(lenString int) string{
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(time.Now().UnixNano())
	var i int
	for i = 0; i < lenString; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
