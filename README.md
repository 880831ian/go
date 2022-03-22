# Go 第一支程式

我們依照官網的示範教學，先建立一個資料夾(要放在我們剛剛的 `GOPATH` 目錄下方歐)，來放我們第一個 Go 程式 "印出 Hello world"：

```
mkdir helloworld
cd helloworld
```

<br>

接著下指令來新增 Go module：

```
$ go mod init github.com/880831ian/go/helloworld
go: creating new go.mod: module github.com/880831ian/go/helloworld
```

<br>

如果成功，會產生一個 go.mod 檔案，我們來看看內容有什麼：

```
$ cat go.mod
module github.com/880831ian/go/helloworld

go 1.18
```
go.mod 是用來定義 module 的文件，用來標示此 module 的名稱、所使用的 go 版本以及相依的 Go module。


<br>

我們分別再新增兩個資料夾，以及兩個 `.go` 檔，來建立我們範例所需要的環境：

```sh
mkdir greeting cli
touch greeting/greeting.go cli/say.go
```

<br>

到目前為止結構如下：

```sh
.
├── cli
│   └── say.go
├── go.mod
└── greeting
    └── greeting.go
```

我們來修改一下 `greeting.go` 以及 `say.go` 程式碼吧。

`greeting.go` 是一個簡單的 package，用以顯示所傳入的字串 ; 而 `say.go` 則是以呼叫 `greeting.go` package 所提供的函示來顯示資料。

<br>

`greeting.go` 內容：

```go
package greeting

import "fmt"

func Say(s string) {
	fmt.Println(s)
}
```

<br>

`say.go` 內容：

```go
package main

import (
	"github.com/880831ian/go/helloworld/greeting"
)

func main(){
	greeting.Say("Hello World")
}
```

<br>

順便來介紹一下程式裡面分別是什麼意思吧！

<br>

* Package：package 主要分成兩種，一個是可執行，另一個則是可重複使用的，而 `package main` 就是可執行的檔案，像我們上面這個有包含 package main 的檔案，在編譯時，就會產生一個 `say` 的執行檔，電腦就是依照此檔案執行的。


* Import：當我們寫程式時，一定會引入其他人寫的套件。而 Go 語言的標準函式庫為開發團隊先寫好，提供一些常用的功能，當然也可以使用其他第三方套件，還滿足內建以及標準函式庫的不足。我們在 `greeting.go` 裡面引入的 `fmt` 就是開發團隊寫好的，然而在 `say.go` 裡面引入的就是`greeting.go` ，我們就可以使用其內容的函示來做使用。


* Main Function：每個 Go 語言的專案基本上都會有一個主程式，主程式裡的程式通常都為最核心的部分。

<br>

最後使用 `go run say.go` 來將此程式運行起來：

```sh
$ go run say.go
Hello World
```

就可以看到程式成功將 Hello World 給印出來拉！

<br>

### 常見指令

接下來要簡單介紹一下常用的另外3個指令，分別是 `go build`、`go install`、`go clean`：

<br>

`go get`：來下載套件到當前的模組，並安裝他們

```
$ go get github.com/fatih/color
go: downloading github.com/fatih/color v1.13.0
.... 省略 ....
```

<br>

`go build`：還記得我們前面說 Go 是編譯式程式，所以我們可以將程式用 `go build` 來編譯成電腦看得懂的執行檔歐，檔案會存放在當前目錄或是指定目錄中 ~

```
$ ls
go.mod

$ go build cli/say.go
go.mod   say

$ ./say
Hello World
```
多的這個 say 就是編譯後的執行檔，將他執行會顯示跟我們使用 run 來運行的一樣，顯示 Hello World。

<br>

`go install`：如果編譯沒有錯誤，一樣跟 build 會產生執行檔，不同的是，會將執行檔，產生於 $GOPATH/bin 內。

```
$ ls /Users/ian_zhuang/go/bin
dlv          go-outline   gomodifytags goplay       gopls        gotests      impl         staticcheck

$ go install

$ ls /Users/ian_zhuang/go/bin 
dlv          go-outline   gomodifytags goplay       gopls        gotests      hello        impl         staticcheck
```

<br>

`go clean`：執行後會將 build 產生的檔案都刪除 (install 的不會)

```
$  ls
go.mod   hello    hello.go

$ go clean

$ ls 
go.mod   hello.go
```

<br>

### 套件相依性管理

Go modules 提供的另一個方便的功能則是套件相依性管理，接下來實際透過以下指令來安裝套件：

```sh
$ go get github.com/fatih/color
go: downloading github.com/fatih/color v1.13.0
go: downloading github.com/mattn/go-isatty v0.0.14
go: downloading github.com/mattn/go-colorable v0.1.9
go: downloading golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
go: added github.com/fatih/color v1.13.0
go: added github.com/mattn/go-colorable v0.1.9
go: added github.com/mattn/go-isatty v0.0.14
go: added golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
```

<br>

安裝成功，可以再查看一下 go.mod：
```
module github.com/880831ian/go/helloworld

go 1.18

require github.com/fatih/color v1.13.0

require (
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
)
```
會多了下面這些 `require github.com/fatih/color v1.13.0` 代表目前專案使用 v1.13.0 版本的 `github.com/fatih/color`

下面的 indirect 指的是被相依的套件所使用的 package

<br>

接著我們將 `greeting.go`、`say.go` 兩個檔案修改一下，使用我們剛剛所安裝的 package：

<br>

`greeting.go`

```go
package greeting

import (
	"fmt"

	"github.com/fatih/color"
)

func Say(s string) {
	fmt.Println(s)
}

func SayWithRed (s string) {
	color.Red(s)
}

func SayWithBlue (s string) {
	color.Blue(s)
}

func SayWithYellow (s string) {
	color.Yellow(s)
}
```
我再多 import 了剛剛的 `github.com/fatih.color`，並使用該套件的函式 `color` 來分別顯示 `Red`、`Blud`、`Yellow` 三種顏色。

<br>

`say.go`

```go
package main

import (
	"github.com/880831ian/go/helloworld/greeting"
)

func main(){
	greeting.Say("Hello World")
	greeting.SayWithRed("Hello World")
	greeting.SayWithBlue("Hello World")
	greeting.SayWithYellow("Hello World")
}
```
我們將 `greeting` 三種顯示顏色的函示帶入。

<br>

一樣我們來運行一下程式，來看看結果如何，這次我們直接編譯，使用 `go build` 來編譯，最後直接執行產生的執行檔：

```sh
$ go build cli/say.go
$ ./say
Hello World
Hello World //紅色
Hello World //藍色
Hello World //黃色
```
由於 Makedown 沒辦法於程式碼區域顯示正確顏色，用註解標示一下XD
