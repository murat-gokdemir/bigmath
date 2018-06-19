package bigmath

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test05(t *testing.T) {

	inp, err := ioutil.ReadFile("testdoc.txt")
	if err != nil {
		panic(err)
	}

	lastn := uint(inp[len(inp)-2]) - 48 //int() that covered uint() was removed
	fmt.Println("testdoc lastn :", lastn)

	div := 5
	n := 0
	for i := 0; i < len(inp)-1; i++ {
		n = 10*n + (int(inp[i]) - 48)
		cnt := 0
		for j := div; j < n+1; j += div {
			cnt += 1
		}
		n -= cnt * div
		fmt.Println("testdoc i:", i, "xcount:", cnt, "testdoc n:", n)
	}
}

func Test01(t *testing.T) {
	var r string = strconv.Itoa(123456789)
	leftch := r[1:9]
	nbr, mistake := strconv.Atoi(r[1:5])
	fmt.Println("r:", r, "type:", reflect.TypeOf(r))
	fmt.Println("number:", nbr, "type:", reflect.TypeOf(nbr), mistake)

	fmt.Println(strings.Count("1234567890123456789", "")) // before & after each rune
	fmt.Println(leftch)
}

func Test19(t *testing.T) {
	var divider int = 19
	bignbr, err := ioutil.ReadFile("butunsel.txt")
	fmt.Println("length of bignumber:", len(bignbr), "capacity of bignumber:", cap(bignbr))
	lastn := int(uint(bignbr[len(bignbr)-2])) - 48
	fmt.Println("lastn :", lastn)
	n := 0
	for i := 0; i < len(bignbr)-12690; i++ {
		fmt.Println("i:", i)
	}
	var i int = 0
	n = 10*n + (int(uint(bignbr[i])) - 48)
	n = 10*n + (int(uint(bignbr[i+1])) - 48) - 3*divider

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("first two number:", n)
	fmt.Println(bignbr[12691:12692])
}
