# csvtools
csv解析工具包  拥有更多的容错性

### 安装
``` 
go get github.com/dollarkillerx/csvtools
```

### 使用
``` 
	csv, e := csvtools.ReadFile("test.csv")
	if e != nil {
		panic(e)
	}
	strings := csv.Decode(0,2)  // 如果为空 就打印全部  ,反之打印指定序号的
	if e != nil {
		panic(e)
	}
	for _, v := range strings {
		log.Println(v)
	}
```