# 🐹 maldev-go  


---

## 📂 hello-world

## init

```bash
go mod init helloworld
```

## main.go

```go

package main
import "fmt"

func main () {
	fmt.Println("hello")
}

```

## 🌍 Cross Compilation

🔹 Build for Windows (x64):

```bash
GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-s -w" exp/main.go
```

---


