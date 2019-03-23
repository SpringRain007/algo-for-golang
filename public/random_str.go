package public

import (
	"math/rand"
	"time"
)

func GetRandomStr(l int) string {
	str := "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	bytes := []byte(str)
	result := []byte{}
	rand.NewSource(time.Now().UnixNano())
	for i:=0;i<l;i++ {
		result = append(result,bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
