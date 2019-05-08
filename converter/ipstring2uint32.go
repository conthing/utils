package converter

import (
	"strings"
	"errors"
	"fmt"
	"strconv"
)

//ErrInvalidIPString IP string invalid
var ErrInvalidIPString = errors.New("Invalid IP String")

//Uint32ToIPString Convert uint32 to IP
func Uint32ToIPString(n uint32) string {   
    return fmt.Sprintf("%d.%d.%d.%d",(n>>24)&0xff,(n>>16)&0xff,(n>>8)&0xff,n&0xff)
}

//IPStringToUint32 Convert IP to uint32
func IPStringToUint32(ip string) (uint32,error) {
	var sum uint32
	bytes := strings.Split(ip, ".")
	
	if len(bytes)!=4 {
		return 0, ErrInvalidIPString
	}

	for _,b := range bytes{
		sum <<= 8
		n, err := strconv.Atoi(b)
		if err != nil || n < 0 || n > 255 {
			return 0,ErrInvalidIPString
		}
		sum += uint32(n)
	}
	
    return sum,nil
}