/**
 * @Author: DollarKillerX
 * @Description: csv.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:54 2019/11/29
 */
package csvtools

import (
	bytes2 "bytes"
	csv2 "encoding/csv"
	"io"
	"io/ioutil"
)

type csv struct {
	ios *csv2.Reader
}

func ReadFile(file string) (*csv, error) {
	bytes, e := ioutil.ReadFile(file)
	if e != nil {
		return nil, e
	}
	reader := csv2.NewReader(bytes2.NewReader(bytes))
	return &csv{ios: reader}, nil
}

func ReadByte(byt []byte) (*csv, error) {
	reader := csv2.NewReader(bytes2.NewReader(byt))
	return &csv{ios: reader}, nil
}

// 解析你需要的列
func (c *csv) Decode(col ...int) ([][]string) {
	source := c.decode()
	if len(source) == 0 {
		return nil
	}
	if len(col) == 0 {
		return source
	}
	result := [][]string{}
	for _, v := range source {
		item := []string{}
		for _, key := range col {
			if len(v) < key {
				continue
			}
			item = append(item, v[key])
		}
		result = append(result, item)
	}
	return result
}

func (c *csv) decode() ([][]string) {
	result := [][]string{}
	for {
		record, e := c.ios.Read()
		if e == io.EOF {
			break
		}
		if e != nil {
			continue
		}
		result = append(result, record)
	}
	return result
}
