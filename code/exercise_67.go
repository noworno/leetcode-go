package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(na string, nb string) string {
	na = reverse(na)
	nb = reverse(nb)
	a := []byte(na)
	b := []byte(nb)
	slen := len(a)
	if slen < len(b) {
		slen = len(b)
		a, b = b, a
	}

	flag := false
	for i := 0; i < slen; i++ {
		if i == len(b) {
			break
		}

		if a[i] == '0' && b[i] == '0' {

			if flag {
				a[i] = '1'
				flag = false
			} else {
				a[i] = '0'
			}
		} else if a[i] == '0' && b[i] == '1' || a[i] == '1' && b[i] == '0' {

			if flag {
				a[i] = '0'
				fmt.Printf("\t %c --- %c\n", a[i], b[i])
			} else {
				a[i] = '1'
			}
		} else if a[i] == '1' && b[i] == '1' {

			if flag {
				a[i] = '1'
			} else {
				a[i] = '0'
				flag = true
			}
		}
	}
	if flag {
		for k := len(b); k < slen; k++ {
			if a[k] == '0' {
				a[k] = '1'
				flag = false
				break
			}
			a[k] = '0'
		}
	}
	if flag {
		a = append(a, '1')
	}
	return reverse(string(a))
}

func reverse(s string) string {
	ns := []byte(s)
	for i := 0; i < len(ns)/2; i++ {
		j := len(ns) - i - 1
		ns[i], ns[j] = ns[j], ns[i]
	}
	return string(ns)
}
