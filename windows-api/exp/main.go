package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func main() {
	text := "hi...."
	caption := "from lynk4"
	_, err := windows.MessageBox(windows.HWND(0), windows.StringToUTF16Ptr(text), windows.StringToUTF16Ptr(caption), windows.MB_OK)
	if err != nil {
		fmt.Println("didn't work........")
	}
	fmt.Println("it worked.....")
}
