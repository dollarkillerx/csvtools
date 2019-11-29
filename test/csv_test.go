/**
 * @Author: DollarKillerX
 * @Description: csv_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午12:04 2019/11/29
 */
package test

import (
	"errors"
	"github.com/dollarkillerx/csvtools"
	"log"
	"reflect"
	"testing"
)

func TestCsv(t *testing.T) {
	csv, e := csvtools.ReadFile("test.csv")
	if e != nil {
		panic(e)
	}
	strings := csv.Decode(0, 2)
	if e != nil {
		panic(e)
	}
	for _, v := range strings {
		log.Println(v)
	}
}

type he struct {
	Ip string `csv:"ip"`
}

// 反射测试
func TestRef(t *testing.T) {
	csvList := &[]he{}
	e := mu(csvList)
	if e != nil {
		log.Fatalln(e)
	}
}

func mu(ptr interface{}) error {
	kind := reflect.TypeOf(ptr).Kind()
	if kind != reflect.Ptr {
		return errors.New("not prt")
	}
	ptrc := reflect.TypeOf(ptr).Elem()
	if ptrc.Kind() != reflect.Slice {
		return errors.New("not slice")
	}
	i := ptrc.NumField()
	log.Println(i)

	return nil
}

func TestSlice(t *testing.T) {
	slc := []string{}
	tpc(&slc)
	log.Println(slc)
}
func tpc(it *[]string) {
	//it = append(it, "sadasd")
}