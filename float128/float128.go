// This package implements 128-bit ("double double") floating point using
// a pair of 64-bit hardware floating point values and standard hardware
// floating point operations. It is based directly on libqd by Yozo Hida,
// Xiaoye S. Li, David H. Bailey, Yves Renard and E. Jason Riedy. Source:
// http://crd.lbl.gov/~dhbailey/mpdist/qd-2.3.13.tar.gz
package float128

import (
	"errors"
	"fmt"
	"math"
)

// A Float128 represents a double-double floating point number with
// 106 bits of mantissa, or about 32 decimal digits. The zero value
// for a Float128 represents the value 0.
type Float128 [2]float64 // float128 represented by two float64s

//
// SET/GET
//

// Set 128-bit floating point object to 0.0
func Zero() (result Float128) {
	result[0] = 0.0
	result[1] = 0.0
	return
}

// Set 128-bit floating point object to 1.0
func One() (result Float128) {
	result[0] = 1.0
	result[1] = 0.0
	return
}

// Set 128-bit floating point object from float64 value
func SetFloat64(hi float64) (result Float128) {
	result[0] = hi
	result[1] = 0.0
	return
}

// Get float64 value from 128-bit floating point object
func (f Float128) Float64() float64 {
	return f[0]
}

// Set 128-bit floating point object from int64 value
func SetInt64(i int64) (result Float128) {
	result[0] = float64(i) // TODO: need to handle >53 bits more carefully
	result[1] = 0.0
	return
}

// Get int64 value from 128-bit floating point object
func Int64(f Float128) int64 {
	return int64(f[0]) // TODO: need to handle >53 bits more carefully
}

// Set 128-bit floating point object from pair float64 values
func SetFF(hi, lo float64) (result Float128) {
	result[0] = hi
	result[1] = lo
	return
}

// Get pair of float64 values from 128-bit floating point object
func FF(f Float128) (float64, float64) {
	return f[0], f[1]
}

//
// COMPARISON
//

func Compare(a, b Float128) int {
	switch {
	case a[0] < b[0]:
		return -1
	case a[0] > b[0]:
		return 1
	}
	switch {
	case a[1] < b[1]:
		return -1
	case a[1] > b[1]:
		return 1
	}
	return 0
}

// IsLT returns whether a < b.
// The name LT (Less-Than) remembers FORTRAN's ".LT." predicate.
func IsLT(a, b Float128) bool {
	return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
}

// IsLE returns whether a <= b.
// The name LE (Less-than-or-Equal-to) remembers FORTRAN's ".LE." predicate.
func IsLE(a, b Float128) bool {
	return a[0] < b[0] || (a[0] == b[0] && a[1] <= b[1])
}

// IsEQ returns whether a == b.
// The name EQ (Equal-to) remembers FORTRAN's ".EQ." predicate.
func IsEQ(a, b Float128) bool {
	return a[0] == b[0] && a[1] == b[1]
}

// IsGE returns whether a >= b.
// The name GE (Greater-than-or-Equal-to) remembers FORTRAN's ".GE." predicate.
func IsGE(a, b Float128) bool {
	return a[0] > b[0] || (a[0] == b[0] && a[1] >= b[1])
}

// IsGT returns whether a > b.
// The name GT (Greater-Than) remembers FORTRAN's ".GT." predicate.
func IsGT(a, b Float128) bool {
	return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
}

// IsNE returns whether a != b.
// The name NE (Not-Equal-to) remembers FORTRAN's ".NE." predicate.
func IsNE(a, b Float128) bool {
	return a[0] != b[0] || a[1] != b[1]
}

//
// CHARACTERIZATION
//

// IsZero returns whether f is 0.
func IsZero(f Float128) bool {
	return f[0] == 0.0 && f[1] == 0.0
}

// IsPositive returns whether f is strictly greater than zero.
func IsPositive(f Float128) bool {
	return f[0] > 0.0
}

// IsNegative returns whether f is strictly less than zero.
func IsNegative(f Float128) bool {
	return f[0] < 0.0
}

// IsOne returns whether f is 1.
func IsOne(f Float128) bool {
	return f[0] == 1.0 && f[1] == 0.0
}

// IsNaN returns whether f is an IEEE 754 "not-a-number" value.
func IsNaN(f Float128) bool {
	return math.IsNaN(f[0])
}

// IsInf returns whether f is an infinity, according to sign.
// If sign > 0, IsInf returns whether f is positive infinity.
// If sign < 0, IsInf returns whether f is negative infinity.
// If sign == 0, IsInf returns whether f is either infinity.
func IsInf(f Float128, sign int) bool {
	return math.IsInf(f[0], sign)
}

//
// I/O
//

// Error codes returned by failures to scan a floating point number.
var (
	errPoint    = errors.New("float128: multiple '.' symbols")
	errPositive = errors.New("float128: internal '+' sign")
	errNegative = errors.New("float128: internal '-' sign")
	errMantissa = errors.New("float128: no mantissa digits")
)

func (f *Float128) Scan(s fmt.ScanState, ch rune) (err error) {
	(*f) = Zero()

	// skip leading space characters
	s.SkipSpace()

	var done, pointSet bool
	var digits, point, sign, exponent int
	for !done {
		ch, _, err := s.ReadRune()
		if err != nil {
			break
		}

		if ch >= '0' && ch <= '9' {
			f.Mul(Float128{10.0, 0.0})
			f.Add(Float128{float64(ch - '0'), 0.0})
			digits++
		} else {
			switch ch {
			case '.':
				if pointSet {
					return errPoint
				}
				point = digits
				pointSet = true
			case '+':
				if sign != 0 || digits > 0 {
					return errPositive
				}
				sign = 1
			case '-':
				if sign != 0 || digits > 0 {
					return errNegative
				}
				sign = -1
			case 'e', 'E':
				_, err = fmt.Fscanf(s, "%d", &exponent)
				if err != nil {
					return err
				}
				done = true
			default:
				s.UnreadRune()
				done = true
			}

		}
	}

	if digits == 0 {
		return errMantissa
	}

	if pointSet {
		exponent -= digits - point
	}

	if exponent != 0 {
		t := SetFloat64(10.0)
		pot := PowerI(t, int64(exponent))
		f.Mul(pot)
	}

	if sign == -1 {
		f.Neg()
	}

	return nil
}

func (f *Float128) toDigits(precision int) (digits string, expn int) {
	D := precision + 1 // number of digits to compute
	s := make([]byte, D+1)

	r := Abs(*f)

	// handle f == 0.0
	if f[0] == 0.0 {
		expn = 0
		for i := 0; i < precision; i++ {
			s[i] = '0'
		}
		digits = string(s[0:precision])
		return
	}

	// First determine the (approximate) exponent.
	expn = int(math.Floor(math.Log10(math.Abs(f[0]))))

	var t Float128
	switch {
	case expn < -300:
		t = SetFloat64(10.0)
		t.PowerI(300)
		r.Mul(t)
		t = SetFloat64(10.0)
		t.PowerI(int64(expn) + 300)
		r.Div(t)
	case expn > 300:
		r.LdexpI(-53)
		t = SetFloat64(10.0)
		t.PowerI(int64(expn))
		r.Div(t)
		r.LdexpI(53)
	default:
		t = SetFloat64(10.0)
		t.PowerI(int64(expn))
		r.Div(t)
	}

	// adjust exponent if off by one
	ten := SetFloat64(10.0)
	switch {
	case IsGE(r, ten):
		f := SetFloat64(10.0)
		r.Div(f)
	case IsLT(r, One()):
		t := SetFloat64(10.0)
		r.Mul(t)
	}

	// verify exponent
	if IsGE(r, ten) || IsLT(r, One()) {
		// error: can't compute exponent
		return
	}

	// extract the digits
	for i := 0; i < D; i++ {
		d := int64(r[0])
		t := SetInt64(d)
		r.Sub(t)
		t = SetFloat64(10.0)
		r.Mul(t)
		s[i] = byte(d + '0')
	}

	// fix out of range digits
	for i := D - 1; i > 0; i-- {
		if s[i] < '0' {
			s[i-1]--
			s[i] += 10
		} else if s[i] > '9' {
			s[i-1]++
			s[i] -= 10
		}
	}

	// verify digits
	if s[0] <= '0' {
		// error: non-positive leading digit
		return
	}

	// round result
	if s[D-1] >= '5' {
		s[D-2]++

		for i := D - 2; i > 0 && s[i] > '9'; i-- {
			s[i] -= 10
			s[i-1]++
		}
	}

	// if first digit is 10 after rounding, shift everything
	if s[0] > '9' {
		expn++
		for i := precision; i >= 2; i-- {
			s[i] = s[i-1]
		}
		s[0] = '1'
		s[1] = '0'
	}

	digits = string(s[0:precision])
	return
}

// Convert to string for output
func (f Float128) String() string {
	digits, exponent := f.toDigits(32)
	s := "+"
	if f[0] < 0.0 {
		s = "-"
	}
	return fmt.Sprintf("%s%s.%se%+03d", s, digits[0:1], digits[1:], exponent)
}

//
// ADDITION
//

// Compute fl(a+b) and err(a+b).
func twoSum(a, b float64) (s, err float64) {
	s = a + b
	bb := s - a
	err = (a - (s - bb)) + (b - bb)
	return
}

// Compute fl(a+b) and err(a+b).  Assumes |a| >= |b|.
func quickTwoSum(a, b float64) (s, err float64) {
	s = a + b
	err = b - (s - a)
	return
}

// Compute D = D + D
func Add(a, b Float128) (f Float128) {
	s1, s2 := twoSum(a[0], b[0])
	t1, t2 := twoSum(a[1], b[1])
	s2 += t1
	s1, s2 = quickTwoSum(s1, s2)
	s2 += t2
	f[0], f[1] = quickTwoSum(s1, s2)
	return
}

// Compute D += D
func (f *Float128) Add(a Float128) {
	s1, s2 := twoSum(f[0], a[0])
	t1, t2 := twoSum(f[1], a[1])
	s2 += t1
	s1, s2 = quickTwoSum(s1, s2)
	s2 += t2
	f[0], f[1] = quickTwoSum(s1, s2)
}

//
// SUBTRACTION
//

// Compute fl(a-b) and err(a-b).
func twoDiff(a, b float64) (s, err float64) {
	s = a - b
	bb := s - a
	err = (a - (s - bb)) - (b + bb)
	return
}

// Compute D = D - D
func Sub(a, b Float128) (f Float128) {
	s1, s2 := twoDiff(a[0], b[0])
	t1, t2 := twoDiff(a[1], b[1])
	s2 += t1
	s1, s2 = quickTwoSum(s1, s2)
	s2 += t2
	f[0], f[1] = quickTwoSum(s1, s2)
	return
}

// Compute D -= D
func (f *Float128) Sub(a Float128) {
	s1, s2 := twoDiff(f[0], a[0])
	t1, t2 := twoDiff(f[1], a[1])
	s2 += t1
	s1, s2 = quickTwoSum(s1, s2)
	s2 += t2
	f[0], f[1] = quickTwoSum(s1, s2)
}

//
// SIGNS
//

// Compute D = -D
func (f *Float128) Neg() {
	f[0], f[1] = -f[0], -f[1]
}

// Compute Abs(D)
func (f *Float128) Abs() {
	if f[0] < 0 {
		f[0], f[1] = -f[0], -f[1]
	}
}

// Compute D = Abs(D)
func Abs(a Float128) Float128 {
	if a[0] < 0 {
		return Float128{-a[0], -a[1]}
	}
	return a
}

//
// MULTIPLICATION
//

const (
	splitter       = 134217729.0           // = 2^27 + 1
	splitThreshold = 6.69692879491417e+299 // = 2^996
)

// Compute high and lo words of a float64 value
func split(a float64) (hi, lo float64) {
	if a > splitThreshold || a < -splitThreshold {
		a *= 3.7252902984619140625e-09 // 2^-28
		temp := splitter * a
		hi = temp - (temp - a)
		lo = a - hi
		hi *= 268435456.0 // 2^28
		lo *= 268435456.0 // 2^28
	} else {
		temp := splitter * a
		hi = temp - (temp - a)
		lo = a - hi
	}
	return
}

// Compute fl(a*b) and err(a*b).
func twoProd(a, b float64) (p, err float64) {
	p = a * b
	aHi, aLo := split(a)
	bHi, bLo := split(b)
	err = ((aHi*bHi - p) + aHi*bLo + aLo*bHi) + aLo*bLo
	return
}

// Compute D * 2**exp
func (f *Float128) LdexpI(exp int) {
	f[0] = math.Ldexp(f[0], exp)
	f[1] = math.Ldexp(f[1], exp)
}

// Compute D * 2**exp
func Ldexp(a Float128, exp int) (ret Float128) {
	ret[0] = math.Ldexp(a[0], exp)
	ret[1] = math.Ldexp(a[1], exp)
	return
}

// Compute D = D * F, where F = 2**k
func (f *Float128) MulPwr2(a Float128, b float64) {
	f[0] *= b
	f[1] *= b
}

// Compute D = D * D
func Mul(a, b Float128) (f Float128) {
	p1, p2 := twoProd(a[0], b[0])
	p2 += a[0]*b[1] + a[1]*b[0]
	f[0], f[1] = quickTwoSum(p1, p2)
	return
}

// Compute D *= D
func (f *Float128) Mul(a Float128) {
	p1, p2 := twoProd(f[0], a[0])
	p2 += f[0]*a[1] + f[1]*a[0]
	f[0], f[1] = quickTwoSum(p1, p2)
}

//
// POWERS
//

// Compute fl(a*a) and err(a*a).  Faster than twoProd.
func twoSqr(a float64) (q, err float64) {
	q = a * a
	hi, lo := split(a)
	err = ((hi*hi - q) + 2.0*hi*lo) + lo*lo
	return
}

// Compute D = D^2
func Sqr(a Float128) (f Float128) {
	p1, p2 := twoSqr(a[0])
	p2 += 2.0 * a[0] * a[1]
	p2 += a[1] * a[1]
	f[0], f[1] = quickTwoSum(p1, p2)
	return
}

// Compute D^2
func (f *Float128) Sqr() {
	*f = Sqr(*f)
}

func absInt64(n int64) int64 {
	if n < 0 {
		n = -n
	}
	return n
}

// Compute D = D^n
func PowerI(a Float128, n int64) (f Float128) {
	if n == 0 {
		if IsZero(a) { // 0**0 is invalid
			f = SetFloat64(math.NaN())
		} else { // x**0 is 1.0
			f = SetFloat64(1.0)
		}
		return
	}

	r := a
	s := Float128{1.0, 0.0}
	N := absInt64(n)

	if N > 1 {
		for N > 0 {
			if N&1 == 1 {
				s.Mul(r)
			}
			N >>= 1
			if N > 0 {
				r.Sqr()
			}
		}
	} else {
		s = r
	}

	if n < 0 {
		f = SetFloat64(1.0)
		f.Div(s)
	} else {
		f = s
	}

	return
}

// Compute D^n
func (f *Float128) PowerI(n int64) {
	*f = PowerI(*f, n)
}

//
// DIVISION
//

// Compute D = D / D
func Div(a, b Float128) (f Float128) {
	q1 := a[0] / b[0]
	f1 := SetFloat64(q1)

	t1 := Mul(f1, b)
	r := Sub(a, t1) // r = a - q1*b

	q2 := r[0] / b[0]
	f2 := SetFloat64(q2)
	t2 := Mul(f2, b)
	r.Sub(t2) // r -= q2*b

	q3 := r[0] / b[0]
	f3 := SetFloat64(q3)

	f[0], f[1] = quickTwoSum(q1, q2)
	f.Add(f3)

	return
}

// Compute D /= D
func (f *Float128) Div(a Float128) {
	*f = Div(*f, a)
}

// Compute D = Floor(D)
func Floor(a Float128) Float128 {
	return Float128{math.Floor(a[0]), 0.0}
}