# Gofuzz

A simple web fuzzer written in Go.

## Usage
```bash
$ gofuzz -w wordlist.txt -u http://example.com?password={}
```

### Note
I do not recomend using this tool for anything serious. it is functional but not efficient and was written for educational purposes.