package main

import (
	"fmt"
	"io/ioutil"
)

func IntDivRemain(fname string, div int) int {
	inp, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	n := 0
	for i := 0; i < len(inp)-1; i++ {
		n = 10*n + (int(inp[i]) - 48)
		cnt := 0
		for j := div; j < n+1; j += div {
			cnt += 1
		}
		n -= cnt * div
		//fmt.Println("testdoc i:", i, "xcount:", cnt, "testdoc n:", n)
	}
	return n
}

func main() {
	fmt.Println(IntDivRemain("bigmath/butunsel.txt", 20))
}
