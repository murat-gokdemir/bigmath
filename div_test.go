package bigmath

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestIntDivision(t *testing.T) {
	//TEST 1
	fname := "testdoc.txt"
	inp, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Error(err)
	}

	eV := 3
	if r, _ := IntDivRemain2(string(inp), 5); r != eV {
		t.Errorf("remain expected %d, but found: %d", eV, r)
	}

	//TEST 2
	eV = 1
	if r, _ := IntDivRemain2("6", 5); r != eV {
		t.Errorf("remain expected %d, but found: %d", eV, r)
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
