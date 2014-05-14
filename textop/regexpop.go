package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

//判断ip
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}

	return true
}

func crawlerBaidu() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http get error")
	}

	src := string(body)
	fmt.Print(src)

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	fmt.Print(src)
}

func test1() {
	src := []byte(`
        call hello alice
        hello bob
        call hello eve
    `)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}

func test2() {
	src := `123.125.71.118 - - [11/May/2014:00:00:05 +0800] "GET /company/news/24.html HTTP/1.1" 200 3661 "-" "Mozilla/5.0 (iphone; U; CPU iPhone OS 4_3_5 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5" 120.201.105.162, 127.0.0.1, 10.65.43.62 0.004 0.004`
	pat := regexp.MustCompile(`^(?P<v1>.*?) .*? .*? \[(?P<v2>.*?)\] "(?P<v3>GET|POST) (?P<v4>.*?)" (?P<v5>.*?) (?P<v6>.*?) .*? ".*?" .* (?P<v7>.*?) (?P<v8>.*?)\s*$`)
	fmt.Printf("%v \n", pat.FindStringSubmatch(src))
}

func test3() {
	src := `123.125.71.118 - - [11/May/2014:00:00:05 +0800] "GET /company/news/24.html HTTP/1.1" 200 3661 "-" "Mozilla/5.0 (iphone; U; CPU iPhone OS 4_3_5 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5" 120.201.105.162, 127.0.0.1, 10.65.43.62 0.004 0.004`
	pat := regexp.MustCompile(`^(?P<v1>.*?) .*? .*? \[(?P<v2>.*?)\] "(?P<v3>GET|POST) (?P<v4>.*?) .*?" (?P<v5>.*?) (?P<v6>.*?) .*? .* (?P<v7>.*?) (?P<v8>.*?)\s*$`)
	values := pat.FindStringSubmatch(src)
	values = values[1:]
	fmt.Printf("%v \n", values)
}

func test4() {
	src := `abcdefgh`
	pat := regexp.MustCompile(`(?P<v1>cd)(?P<v2>ef)`)
	fmt.Printf("%v \n", pat.FindStringSubmatch(src))
}

func main() {
	// fmt.Println(IsIP("127.a.0.1"))
	test3()
}
