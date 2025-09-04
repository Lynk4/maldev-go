package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    key := byte(133)

    pShellcodePath := flag.String("file", "", "Path of the shellcode")
    flag.Parse()

    shellcodePath := *pShellcodePath

    clearshellcodeByte, err := os.ReadFile(shellcodePath)
    if err != nil {
        fmt.Println("Error.......opening.....file.....")
        fmt.Println(err.Error())

    }

    var encryptedShellcode []byte

    for i := 0; i < len(clearshellcodeByte); i++ {
        encryptedShellcode = append(encryptedShellcode, clearshellcodeByte[i] ^ key)
    }
    fmt.Print("[]byte{")
    for i := 0; i < len(encryptedShellcode); i++ {
    if i == len(encryptedShellcode)-1 {
        fmt.Printf("%d", encryptedShellcode[i]) // last element, no comma
    } else {
        fmt.Printf("%d, ", encryptedShellcode[i]) // add comma + space
    }
}
    fmt.Println("}")
}