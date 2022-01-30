package util

import (
	"math/rand"
	"time"
)

//type GenRandStrEtx struct{
//
//}
//
//}
/**
sType 1 为只含数字
 */
func GenRandStr(sType int32,num int)string{
	raw := "0123456789"
	retStr := ""
	if sType == 0{
		strLen := len(raw)
		for i:= 0;i< num;i++{
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			idx := r.Int31n(int32(strLen))
			retStr = retStr + string(raw[idx])
		}
	}
	return retStr
}
