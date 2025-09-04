package main

import (
    "flag"
    "fmt"
    "os"
)

func main()  {

    key := byte(133)

    // ./prepare -file=/home/kali/clearshellcode.bin
    pShellcodePath := flag.String("file", "", "Path Of the Shellcode")
    flag.Parse()

    shellcodePath := *pShellcodePath
    
    clearShellcodeByte, err := os.ReadFile(shellcodePath)
    if err != nil {
        fmt.Println("Error Opening file")
        fmt.Println(err.Error())
    }

    var encryptedShellcode []byte

    for i := 0; i < len(clearShellcodeByte); i++ {
        encryptedShellcode = append(encryptedShellcode, clearShellcodeByte[i] ^ key )
    }

    fmt.Print("[]byte{")
    for i := 0; i < len(clearShellcodeByte); i++ {
        if i == len(clearShellcodeByte) - 1 {
            fmt.Println(encryptedShellcode[i], "}")
        } else {
            fmt.Print(encryptedShellcode[i], ", ")
        }
    }
}
