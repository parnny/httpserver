package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)


type TestBody struct {
	TestInt 		int64
	TestString		string
}

type TestData struct {
	Project 		string
	Type 			string
	Timestamp 		int64
	Data			TestBody
}

func GenData() []byte {
	var datalist [1]TestData
	for i := 0; i < 1; i++ {
		datalist[i] = TestData{
			Project:"test project",
			Type:"login",
			Timestamp:time.Now().UTC().UnixNano(),
			Data:TestBody {
				TestInt:123,
				TestString:"test string",
			},
		}
	}
	data, err := json.Marshal(datalist)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func RequestLogic(url string) bool {
	client := &http.Client{}

	testdata := GenData()

	req, err := http.NewRequest("GET", url,
		bytes.NewReader(testdata))
	if err != nil {
		log.Println(err)
		return false
	}
	client.Do(req)

	return true
}

//func Test_loop(t *testing.T)  {
//	url := *flag.String("url", "http://120.131.12.66:80", "http address")
//	tm := *flag.Duration("sleep", 0, "sleep millisecond (ms), 0 is unsleep")
//	flag.Parse()
//
//	for i:=0; i<5; {
//		res := RequestLogic(url)
//		if tm > 0{
//			time.Sleep(tm * time.Millisecond)
//		}
//		if res {
//			t.Log("test pass ")
//		} else {
//			t.Error("test error")
//		}
//	}
//}

// 测试并发效率
func BenchmarkLoopsParallel(b *testing.B) {
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			RequestLogic("http://120.131.12.66:80")
		}
	})
}

func Benchmark_Loop(b *testing.B)  {
	for i := 0; i< b.N;	i++ {
		RequestLogic("http://120.131.12.66:80")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}