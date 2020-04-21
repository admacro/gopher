// https://golang.org/ref/spec#Integer_literals
package main

import "fmt"

func main() {
	var (
		a  = 1
		b  = 1_000
		g  = 1.23i
		h  = 0b1001
		h1 = 0B1001_001
		i  = 0o123
		i1 = 0O1_23
		j  = 0xBadFace
		j1 = 0Xbad_face7581
		k  = 0x_67_7a_2f_cc_40_c6

		l  = 0.
		m  = 72.40
		m1 = 072.40
		c  = 3.1415926
		d  = 2.3e1
		e  = 2.3e+4
		f  = 2.3E-11
		n  = 2e4
		o  = .25
		p  = .12345E+5
		q  = 1_5.
		r  = 0.15e+0_2
		s  = .44e-0_9

		t = 0x1p-2
		u = 0x2.p10
		v = 0x1.Fp+0
		w = 0X.8p-0
		x = 0X_1FFFP-16
		y = 0x15e - 2 // integer subtraction
	)

	fmt.Println(a, b, c, d, e, f, g, h, h1, i, i1, j, j1, k, l, m, m1, n, o, p, q, r, s, t, u, v, w, x, y)
}
