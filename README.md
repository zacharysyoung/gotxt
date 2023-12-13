# GoTXT: like iconv, but way less

Provide common transcodings with Go's [x/text package](https://pkg.go.dev/golang.org/x/text).

Input can read from a named file, or from Stdin if the file is not provided.

**-in** and **-out** control in the input and output encodings (both default to UTF-8):

```none
% echo 'Hello, 世界' | hexdump -C
00000000  48 65 6c 6c 6f 2c 20 e4  b8 96 e7 95 8c 0a        |Hello, .......|
0000000e
```

```none
% echo 'Hello, 世界' | gotxt -out utf16le | hexdump -C
00000000  48 00 65 00 6c 00 6c 00  6f 00 2c 00 20 00 16 4e  |H.e.l.l.o.,. ..N|
00000010  4c 75 0a 00                                       |Lu..|
00000014
```

```none
% echo 'Hello, 世界' | gotxt -out utf16be | hexdump -C
00000000  00 48 00 65 00 6c 00 6c  00 6f 00 2c 00 20 4e 16  |.H.e.l.l.o.,. N.|
00000010  75 4c 00 0a                                       |uL..|
00000014
```

```none
% echo 'Hello, 世界' | gotxt -out shiftjis | hexdump -C
00000000  48 65 6c 6c 6f 2c 20 90  a2 8a 45 0a              |Hello, ...E.|
0000000c
```

Transforming between incompatible encodings, e.g. UTF-8→Latin-1, will generate an error and point to the beginning of the first incompatible byte (index-1). Any portion of the input that was properly transcoded will be written to Stdout before the error message prints:

```none
% echo 'Hello, 世界' | gotxt -out iso8859-1
Hello, 
error: could not transcode, read input up to byte 8: encoding: rune not supported by encoding.
```

**-list** shows valid names; like the comment says, spaces and hyphens will be stripped out and the name made lowercase before the command tries to match the name:

```none
# names are case insensitive; spaces and hyphens will not be used for comparison, i.e., `gotxt -in UTF-8` = `gotxt -in 'Utf 8'` = `gotxt -in utf8`
Big5
IBM Code Page 037
IBM Code Page 437
IBM Code Page 850
IBM Code Page 852
IBM Code Page 855
IBM Code Page 858
IBM Code Page 860
IBM Code Page 862
IBM Code Page 863
IBM Code Page 865
IBM Code Page 866
IBM Code Page 1047
IBM Code Page 1140
EUCJP
EUCKR
GB18030
GBK
HZ-GB2312
ISO 2022-JP
ISO 8859-1
ISO 8859-2
ISO 8859-3
ISO 8859-4
ISO 8859-5
ISO 8859-6
ISO 8859-6E
ISO 8859-6I
ISO 8859-7
ISO 8859-8
ISO 8859-8E
ISO 8859-8I
ISO 8859-9
ISO 8859-10
ISO 8859-13
ISO 8859-14
ISO 8859-15
ISO 8859-16
KOI8-R
KOI8-U
Macintosh
Macintosh Cyrillic
SHIFT JIS
UTF-8
UTF-8 BOM
UTF-16 BE
UTF-16 BE BOM
UTF-16 LE
UTF-32 BE
UTF-32 BE BOM
UTF-32 LE
Windows 874
Windows 1250
Windows 1251
Windows 1252
Windows 1253
Windows 1254
Windows 1255
Windows 1256
Windows 1257
Windows 1258
X-User-Defined
# names are case insensitive; spaces and hyphens will not be used for comparison, i.e., `gotxt -in UTF-8` = `gotxt -in 'Utf 8'` = `gotxt -in utf8`
```
