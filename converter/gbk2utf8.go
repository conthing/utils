package converter

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ValidGBK GBK编码格式判断，是GB18030的子集，在2字节区间和UTF8有部分重叠
func ValidGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0x7f {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// UTF8toGB18030 UTF8 to GB18030
func UTF8toGB18030(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GB18030.NewEncoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return src, err
	}
	return d, nil
}

// GB18030toUTF8 GB18030 to UTF8
func GB18030toUTF8(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GB18030.NewDecoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return src, err
	}
	return d, nil
}
