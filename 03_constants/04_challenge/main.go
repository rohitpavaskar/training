package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

// NOTE : Always declare constants outside of any function. Constants should ideally have package
// level scope at the least and declaring them inside a function would limit their usage only to
// that particular function which more than often contradicts the purpose of declaring the constants
// in the first place.

func main() {
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d | KB = %d bytes\n", MB/KB, MB)
	fmt.Printf("1 GB = %d | MB = %d | KB = %d bytes\n", GB/MB, GB/KB, GB)
	fmt.Printf("1 TB = %d | GB = %d | MB = %d | KB = %d bytes\n", TB/GB, TB/MB, TB/KB, TB)
}
