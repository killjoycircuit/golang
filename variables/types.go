/*
	int  int8-> 2^8 (256) values like -128 to +127 i.e signed int, int16  int32  int64
	uint uint8-> 2^8 stoes 256 values(unsigned only +ve) uint16 uint32 uint64 uintptr
	float32 float64
	complex64 complex128

	The number represents the number of bits.
	The default type of int is 32 bits and uint is 64 bits.
*/

package variables

import "fmt"

// Capital letter meaning public and small meaning private to current package.
func Types() {
	// declaring. always initiased to certain values 
	var smsSendingLimit int //->0
	var costPerSMS float64 //->0.0
	var hasPermission bool //->false
  
	var username string //->false

  //%v->Prints the default format of a value (generic placeholder).
  //%.2f->Prints a floating-point number with 2 digits after the decimal
  //%q->Prints a string or character in double quotes
	fmt.Printf("Default Values: %v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	fmt.Println()
}

//output
0 0.00 false ""
