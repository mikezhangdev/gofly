package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func GenMd5Str(str string)(string,error){
	h := md5.New()
	_,err := io.WriteString(h,str)
	if err!= nil{
		return "",err
	}
	return fmt.Sprintf("%X",h.Sum(nil)),nil
}
