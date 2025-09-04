## 🖼️ Windows API — MessageBox Example

This example demonstrates how to call a **Windows API function** directly from Go using the `golang.org/x/sys/windows` package.  

Specifically, it calls the `MessageBoxW` function from **user32.dll** to display a pop-up message box.

---

### 📜 Code

```go
package main

import (
	"fmt"
	"golang.org/x/sys/windows"
)

func main() {
	text := "hi...."
	caption := "from lynk4"

	// Call MessageBoxW from user32.dll
	_, err := windows.MessageBox(
		windows.HWND(0),
		windows.StringToUTF16Ptr(text),
		windows.StringToUTF16Ptr(caption),
		windows.MB_OK,
	)

	if err != nil {
		fmt.Println("didn't work........")
	} else {
		fmt.Println("it worked.....")
	}
}

---
```
## cross🛠️ Compiling Go

✅ Make sure you have the golang.org/x/sys/windows package installed: ```kant@APPLEs-MacBook-Pro ~/e/windows-api> go mod tidy```

```bash
kant@APPLEs-MacBook-Pro ~/e/windows-api> GOOS=windows GOARCH=amd64 go build -a -ldflags="-s -w" -trimpath exp/main.go
kant@APPLEs-MacBook-Pro ~/e/windows-api> file main.exe 
main.exe: PE32+ executable (console) x86-64, for MS Windows, 8 sections
kant@APPLEs-MacBook-Pro ~/e/windows-api>
```

---

<img width="1086" height="736" alt="Screenshot 2025-09-04 at 4 56 46 PM" src="https://github.com/user-attachments/assets/f0787924-d7a7-472e-a822-aea86e45a4df" />

---

