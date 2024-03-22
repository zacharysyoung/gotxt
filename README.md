# GoTXT: like iconv, but way less

Provide common transcodings with Go's [x/text package](https://pkg.go.dev/golang.org/x/text).

Install with `go install github.com/zacharysyoung/gotxt@latest`.

Input can read from a named file, or from Stdin if the file is not provided.

**-in** and **-out** control in the input and output encodings (both default to UTF-8):

**-list** shows all valid names, **-list-u** shows names for just the UTF variants.  All names given to -in and -out must match exactly as listed.

```none
% echo 'Hello, 世界' | hexdump -C
00000000  48 65 6c 6c 6f 2c 20 e4  b8 96 e7 95 8c 0a        |Hello, .......|
0000000e
```

```none
% echo 'Hello, 世界' | gotxt -out utf-16-le | hexdump -C
00000000  48 00 65 00 6c 00 6c 00  6f 00 2c 00 20 00 16 4e  |H.e.l.l.o.,. ..N|
00000010  4c 75 0a 00                                       |Lu..|
00000014
```

```none
% echo 'Hello, 世界' | gotxt -out utf-16-be | hexdump -C
00000000  00 48 00 65 00 6c 00 6c  00 6f 00 2c 00 20 4e 16  |.H.e.l.l.o.,. N.|
00000010  75 4c 00 0a                                       |uL..|
00000014
```

```none
% echo 'Hello, 世界' | gotxt -out shift-jis | hexdump -C
00000000  48 65 6c 6c 6f 2c 20 90  a2 8a 45 0a              |Hello, ...E.|
0000000c
```

Trying to transcode a rune with an output encoding that doesn't support the rune, e.g., "界" (U+754C) to ISO-8859-1, will generate an error and point to the beginning of the first incompatible byte (index-1). Any portion of the input that was properly transcoded will be printed to Stdout before the error message prints to Stderr.

```none
% echo 'Hello, 世界' | gotxt -out iso-8859-1>done.txt 2>error.txt
% cat -en done.txt error.txt 
     1  Hello, $
     1  error: could not transcode, read input up to byte 8: encoding: rune not supported by encoding.$
```

The seventh byte (space) was transcoded, the eighth byte (starting the 3-byte sequence for 界) was out of range for ISO-8859-1 and the program errored-out.
