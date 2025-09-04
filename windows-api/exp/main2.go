package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	text := "hi...."
	caption := "from lynk4"
	textUtf16 := windows.StringToUTF16Ptr(text)
	captionUtf16 := windows.StringToUTF16Ptr(caption)
	// _, err := windows.MessageBox(windows.HWND(0), 
windows.StringToUTF16Ptr(text), windows.StringToUTF16Ptr(caption), windows.MB_OK)
	// if err != nil {
	// 	fmt.Println("didn't work........")
	// }
	USER32DLL := windows.NewLazyDLL("user32.dll")
	procMessageBoxW := USER32DLL.NewProc("MessageBoxW")

	procMessageBoxW.Call(uintptr(0),
		uintptr(unsafe.Pointer(textUtf16)),
		uintptr(unsafe.Pointer(&captionUtf16)),
		uintptr(windows.MB_OK))

	fmt.Println("it worked.....")

}

