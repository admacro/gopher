// https://golang.org/ref/spec#Integer_literals
// https://golang.org/ref/spec#Floating-point_literals
// https://golang.org/ref/spec#Imaginary_literals
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

		aa = 0i
		bb = 0123i  // == 123i for backward-compatibility
		cc = 0o123i // == 0o123 * 1i == 83i
		dd = 0xabci // == 0xabc * 1i == 2748i
		ee = 0.i
		ff = 2.71828i
		gg = 1.e+0i
		hh = 6.67428e-11i
		ii = 1E6i
		jj = .25i
		kk = .12345E+5i
		ll = 0x1p-2i // == 0x1p-2 * 1i == 0.25i
	)

	fmt.Println(a, b, c, d, e, f, g, h, h1)
	fmt.Println(i, i1, j, j1, k, l, m, m1)
	fmt.Println(n, o, p, q, r, s, t, u, v, w, x, y)
	fmt.Println(aa, bb, cc, dd, ee, ff, gg, hh, ii, jj, kk, ll)
}
