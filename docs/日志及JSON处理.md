## 标准库
### 日志
#### 日志的基本使用
```go
import (
	"log"
)

func init() {
	log.SetPrefix("TRACE:") // 设定日志的前缀
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile) // 设定日志的格式
}

func main() {
	log.Println("message")
	log.Fatalln("fatal message")
	log.Panicln("panic message")
}
```
#### 日志定义的相关类型
```go
# golang.org/src/log/log.go
const (
// 将下面的位使用或运算符连接在一起，可以控制要输出的信息。没有
// 办法控制这些信息出现的顺序（下面会给出顺序）或者打印的格式
// （格式在注释里描述）。这些项后面会有一个冒号：
// 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
// 日期: 2009/01/23
Ldate = 1 << iota
// 时间: 01:23:23
Ltime
// 毫秒级时间: 01:23:23.123123。该设置会覆盖 Ltime 标志
Lmicroseconds
// 完整路径的文件名和行号: /a/b/c/d.go:23
Llongfile
// 最终的文件名元素和行号: d.go:23
// 覆盖 Llongfile
Lshortfile
// 标准日志记录器的初始值
LstdFlags = Ldate | Ltime
)
```
#### 定制一个日志记录器
```go
import (
  "io"
  "io/ioutil"
  "log"
  "os"
)

var (
  Trace *log.Logger // 记录所有日志
  Info *log.Logger // 重要信息
  Warning *log.Logger // 需要注意的信息
  Error *log.Logger // 非常重要的地方
)

func init() {
  file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    log.Fatalln("Failed to open error log file:", err)
  }

  Trace = Log.New(ioutil.Discard, "TRACE:", log.Ldate|log.Ltime|log.Lshortfile)

  Info = Log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
  
  Warning = log.New(os.Stdout, "WARNING:", log.Ldate|log.Ltime|log.Lshortfile)

  Error = log.New(io.MultiWrite(file, os.Stderr), "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}
```
## JSON
先看一段google的JSON响应
```json
{
  "responseData": {
    "results": [
      {
        "GsearchResultClass": "GwebSearch",
        "unescapedUrl": "https://www.reddit.com/r/golang",
        "url": "https://www.reddit.com/r/golang",
        "visibleUrl": "www.reddit.com",
        "cacheUrl": "http://www.google.com/search?q=cache:W...",
        "title": "r/\u003cb\u003eGolang\u003c/b\u003e - Reddit",
        "titleNoFormatting": "r/Golang - Reddit",
        "content": "First Open Source \u003cb\u003eGolang\u..."
      },
      {
        "GsearchResultClass": "GwebSearch",
        "unescapedUrl": "http://tour.golang.org/",
        "url": "http://tour.golang.org/",
        "visibleUrl": "tour.golang.org",
        "cacheUrl": "http://www.google.com/search?q=cache:O...",
        "title": "A Tour of Go",
        "titleNoFormatting": "A Tour of Go",
        "content": "Welcome to a tour of the Go programming ..."
      }
    ]
  }
}
```
```go
package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

type (
  gResult struct {
    GsearchResultClass string `json:"GsearchResultClass"`
    UnescapedURL string `json:"unescapedUrl"`
    URL string `json:"url"`
    VisibleURL string `json:"visibleUrl"`
    CacheURL string `json:"cacheUrl"`
    Title string `json:"title"`
    TitleNoFormatting string `json:"titleNoFormatting"`
    Content string `json:"content"`
  }

  gResponse struct {
    ResponseData struct {
      Results []gResult `json:"results"`
    } `json:"responseData"`
  }
)

func main() {
  // 这个API挂了
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(gr)
}
```
通过一段反引号标记的字符串，让struct映射JSON中的字段
