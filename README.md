# hasher

This tool was implemented to get different hash sums
of the passed file by using one tool.

## Supported hash algorithms are

- CRC32
- SHA1
- SHA256
- SHA512
- MD5

## Usage

Linux:
> $ ./hasher sha256 my_file.bin
> 5d3a9c7c1f19d4166a1c41786fc678663719b9aa034f597ea234972357e55b90

Windows:
> PS C:\Users\User\hasher> .\hasher.exe sha256 some_file.bin
4f63be23f5338be197eae4fccc078fe90ad5381e2f4cc9f8c78efb9edad321d0

## Build executable

To build this tool from sources you need:
1. Install Go
2. Clone this repository
3. Go to **hasher** directory
4. Run go build command

