/**
 * @Author: DollarKillerX
 * @Description: csv_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午12:04 2019/11/29
 */
package test

import (
	"github.com/dollarkillerx/csvtools"
	"log"
	"testing"
)

func TestCsv(t *testing.T) {
	csv, e := csvtools.ReadFile("test.csv")
	if e != nil {
		panic(e)
	}
	strings := csv.Decode(0,2)
	if e != nil {
		panic(e)
	}
	for _, v := range strings {
		log.Println(v)
	}
}
