package bigmath

import (
	"fmt"
	"strings"
)

func IntDivRemain2(s string, div int) (int, error) {

	inp := []byte(strings.TrimSpace(s))

	fmt.Println(inp)

	if err := IsNumber(inp); err != nil {
		return 0, err
	}

	n := 0
	for i := 0; i < len(inp); i++ {
		n = 10*n + (int(inp[i]) - 48)
		cnt := 0
		for j := div; j <= n; j += div {
			cnt += 1
		}
		n -= cnt * div
		fmt.Println("testdoc i:", i, "xcount:", cnt, "testdoc n:", n)
	}
	return n, nil
}

func IsNumber(b []byte) error {

	return nil
}
