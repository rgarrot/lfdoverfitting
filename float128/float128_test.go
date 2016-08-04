package float128

import (
	"fmt"
	"math"
	"testing"
)

//
// INITIALIZATION
//

func BenchmarkLiteral(b *testing.B) {
	var f Float128
	for i := 0; i < b.N; i++ {
		f = Float128{1, 2}
	}
	f[0] = f[0] // force compiler to see f as used
}

var scanTests = []struct {
	s string
	f float64
}{
	{"0", 0.0},
	{"0.0", 0.0},
	{"1", 1.0},
	{"10", 10.0},
	{"100", 100.0},
	{"-1", -1.0},
	{"-10", -10.0},
	{"-100", -100.0},

	{"-0", 0.0},
	{"+0", 0.0},

	{"1e0", 1e0},
	{"1e10", 1e10},
	{"16777216e-30", 16777216e-30},

	/*
		{"3", 3.0},
		{"3.1", 3.1},
		{"3.14", 3.14},
		{"3.141", 3.141},
		{"3.1415", 3.1415},
		{"3.14159", 3.14159},
		{"3.141592", 3.141592},
		{"3.1415926", 3.1415926},
	*/

	{"512", 512},
	{"1024", 1024},
	{"16777216", 16777216},
}

func TestScan(t *testing.T) {
	for i, a := range scanTests {
		var r Float128
		if _, err := fmt.Sscan(a.s, &r); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		if r.Float64() != a.f {
			t.Errorf("#%d got r = %v; want %v (a)", i, r, a.f)
		}
		/*
			if IsNE(r, SetFloat64(a.f)) {
				t.Errorf("#%d got r = %v; want %v (b)", i, r, SetFloat64(a.f))
			}
		*/
	}
}

// Clear

func TestClear(t *testing.T) {
	r := Float128{1, 2}
	z := Float128{0, 0}
	r = Zero()
	if IsNE(r, z) {
		t.Errorf("#%d got r = %v; want %v", 1, r, z)
	}
}

func BenchmarkClear(b *testing.B) {
	var f Float128
	for i := 0; i < b.N; i++ {
		f = Zero()
	}
	_ = f // make compiler happy
}

// SetF

var setFTests = []struct {
	x float64
	r Float128
}{
	{0, Float128{0, 0}},
	{1, Float128{1, 0}},
}

func TestSetF(t *testing.T) {
	for i, a := range setFTests {
		var r Float128
		r = SetFloat64(a.x)
		if IsNE(r, a.r) {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkSetF(b *testing.B) {
	var f Float128
	for i := 0; i < b.N; i++ {
		f = SetFloat64(1)
	}
	_ = f // make compiler happy
}

// SetFF

var setFFTests = []struct {
	x, y float64
	r    Float128
}{
	{0, 0, Float128{0, 0}},
	{1, 2, Float128{1, 2}},
}

func TestSetFF(t *testing.T) {
	for i, a := range setFFTests {
		var r Float128
		r = SetFF(a.x, a.y)
		if IsNE(r, a.r) {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkSetFF(b *testing.B) {
	var f Float128
	for i := 0; i < b.N; i++ {
		f = SetFF(1, 2)
	}
	_ = f // make compiler happy
}

// SetInt64

func TestSetInt64(t *testing.T) {
	for i := uint64(1); i < 54; i++ {
		var r Float128
		x := int64(1<<i - 1)
		r = SetInt64(x)
		if int64(r[0]) != x || r[1] != 0 { // not really, rework when set can split bits between floats
			t.Errorf("#%d got r = %v; want [%v %v]", i, r, x, 0)
		}
	}
}

func BenchmarkSetInt64(b *testing.B) {
	var f Float128
	for i := 0; i < b.N; i++ {
		f = SetInt64(1<<53 - 1)
	}
	_ = f // make compiler happy
}

//
// COMPARISON
//

// Compare

var compareTests = []struct {
	x, y string
	r    int
}{
	{"-1.0000000000000000000000000000001e+200", "-1.0000000000000000000000000000000e+200", -1},
	{"-1.0000000000000000000000000000000e+200", "-1.0000000000000000000000000000001e+200", +1},
	{"-2", "-1", -1},
	{"-1", "-2", +1},
	{"-1.0000000000000000000000000000001e-200", "-1.0000000000000000000000000000000e-200", -1},
	{"-1.0000000000000000000000000000000e-200", "-1.0000000000000000000000000000001e-200", +1},

	{"0", "0", 0},

	{"+1.0000000000000000000000000000001e-200", "+1.0000000000000000000000000000000e-200", +1},
	{"+1.0000000000000000000000000000000e-200", "+1.0000000000000000000000000000001e-200", -1},
	{"+2", "+1", +1},
	{"+1", "+2", -1},
	{"+1.0000000000000000000000000000001e+200", "+1.0000000000000000000000000000000e+200", +1},
	{"+1.0000000000000000000000000000000e+200", "+1.0000000000000000000000000000001e+200", -1},

	{"3.1415926535897932384626433832795", "3.1415926535897932384626433832796", -1},
}

func TestCompare(t *testing.T) {
	for i, a := range compareTests {
		var x, y Float128
		if _, err := fmt.Sscan(a.x, &x); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		if _, err := fmt.Sscan(a.y, &y); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		r := Compare(x, y)
		if r != a.r {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkCompare(b *testing.B) {
	var x, y Float128
	b.StopTimer()
	fmt.Sscan("3.1415926535897932384626433832795", &x) // truncated
	fmt.Sscan("3.1415926535897932384626433832796", &y) // +1 LSD
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Compare(x, y)
	}
}

// LT

var ltTests = []struct {
	x, y string
	r    bool
}{
	{"-1.0000000000000000000000000000001e+200", "-1.0000000000000000000000000000000e+200", true},
	{"-1.0000000000000000000000000000000e+200", "-1.0000000000000000000000000000001e+200", false},
	{"-2", "-1", true},
	{"-1", "-2", false},
	{"-1.0000000000000000000000000000001e-200", "-1.0000000000000000000000000000000e-200", true},
	{"-1.0000000000000000000000000000000e-200", "-1.0000000000000000000000000000001e-200", false},

	{"0", "0", false},

	{"+1.0000000000000000000000000000001e-200", "+1.0000000000000000000000000000000e-200", false},
	{"+1.0000000000000000000000000000000e-200", "+1.0000000000000000000000000000001e-200", true},
	{"+2", "+1", false},
	{"+1", "+2", true},
	{"+1.0000000000000000000000000000001e+200", "+1.0000000000000000000000000000000e+200", false},
	{"+1.0000000000000000000000000000000e+200", "+1.0000000000000000000000000000001e+200", true},

	{"3.1415926535897932384626433832795", "3.1415926535897932384626433832796", true},
}

func TestLT(t *testing.T) {
	for i, a := range ltTests {
		var x, y Float128
		if _, err := fmt.Sscan(a.x, &x); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		if _, err := fmt.Sscan(a.y, &y); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		r := IsLT(x, y)
		if r != a.r {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkLT(b *testing.B) {
	var x, y Float128
	b.StopTimer()
	fmt.Sscan("3.1415926535897932384626433832795", &x) // truncated
	fmt.Sscan("3.1415926535897932384626433832796", &y) // +1 LSD
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsLT(x, y)
	}
}

// LE

var leTests = []struct {
	x, y string
	r    bool
}{
	{"-1.0000000000000000000000000000001e+200", "-1.0000000000000000000000000000000e+200", true},
	{"-1.0000000000000000000000000000000e+200", "-1.0000000000000000000000000000001e+200", false},
	{"-2", "-1", true},
	{"-1", "-2", false},
	{"-1.0000000000000000000000000000001e-200", "-1.0000000000000000000000000000000e-200", true},
	{"-1.0000000000000000000000000000000e-200", "-1.0000000000000000000000000000001e-200", false},

	{"0", "0", true},

	{"+1.0000000000000000000000000000001e-200", "+1.0000000000000000000000000000000e-200", false},
	{"+1.0000000000000000000000000000000e-200", "+1.0000000000000000000000000000001e-200", true},
	{"+2", "+1", false},
	{"+1", "+2", true},
	{"+1.0000000000000000000000000000001e+200", "+1.0000000000000000000000000000000e+200", false},
	{"+1.0000000000000000000000000000000e+200", "+1.0000000000000000000000000000001e+200", true},

	{"3.1415926535897932384626433832795", "3.1415926535897932384626433832796", true},
}

func TestLE(t *testing.T) {
	for i, a := range leTests {
		var x, y Float128
		if _, err := fmt.Sscan(a.x, &x); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		if _, err := fmt.Sscan(a.y, &y); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		r := IsLE(x, y)
		if r != a.r {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkLE(b *testing.B) {
	var x, y Float128
	b.StopTimer()
	fmt.Sscan("3.1415926535897932384626433832795", &x) // truncated
	fmt.Sscan("3.1415926535897932384626433832796", &y) // +1 LSD
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsLE(x, y)
	}
}

// EQ

var eqTests = []struct {
	x, y string
	r    bool
}{
	{"-1.0000000000000000000000000000001e+200", "-1.0000000000000000000000000000000e+200", false},
	{"-1.0000000000000000000000000000000e+200", "-1.0000000000000000000000000000001e+200", false},
	{"-2", "-1", false},
	{"-1", "-2", false},
	{"-1.0000000000000000000000000000001e-200", "-1.0000000000000000000000000000000e-200", false},
	{"-1.0000000000000000000000000000000e-200", "-1.0000000000000000000000000000001e-200", false},

	{"0", "0", true},

	{"+1.0000000000000000000000000000001e-200", "+1.0000000000000000000000000000000e-200", false},
	{"+1.0000000000000000000000000000000e-200", "+1.0000000000000000000000000000001e-200", false},
	{"+2", "+1", false},
	{"+1", "+2", false},
	{"+1.0000000000000000000000000000001e+200", "+1.0000000000000000000000000000000e+200", false},
	{"+1.0000000000000000000000000000000e+200", "+1.0000000000000000000000000000001e+200", false},

	{"3.1415926535897932384626433832795", "3.1415926535897932384626433832796", false},
}

func TestEQ(t *testing.T) {
	for i, a := range eqTests {
		var x, y Float128
		if _, err := fmt.Sscan(a.x, &x); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		if _, err := fmt.Sscan(a.y, &y); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		r := IsEQ(x, y)
		if r != a.r {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkEQ(b *testing.B) {
	var x, y Float128
	b.StopTimer()
	fmt.Sscan("3.1415926535897932384626433832795", &x) // truncated
	fmt.Sscan("3.1415926535897932384626433832796", &y) // +1 LSD
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsEQ(x, y)
	}
}

// GE

var geTests = []struct {
	x, y string
	r    bool
}{
	{"-1.0000000000000000000000000000001e+200", "-1.0000000000000000000000000000000e+200", false},
	{"-1.0000000000000000000000000000000e+200", "-1.0000000000000000000000000000001e+200", true},
	{"-2", "-1", false},
	{"-1", "-2", true},
	{"-1.0000000000000000000000000000001e-200", "-1.0000000000000000000000000000000e-200", false},
	{"-1.0000000000000000000000000000000e-200", "-1.0000000000000000000000000000001e-200", true},

	{"0", "0", true},

	{"+1.0000000000000000000000000000001e-200", "+1.0000000000000000000000000000000e-200", true},
	{"+1.0000000000000000000000000000000e-200", "+1.0000000000000000000000000000001e-200", false},
	{"+2", "+1", true},
	{"+1", "+2", false},
	{"+1.0000000000000000000000000000001e+200", "+1.0000000000000000000000000000000e+200", true},
	{"+1.0000000000000000000000000000000e+200", "+1.0000000000000000000000000000001e+200", false},

	{"3.1415926535897932384626433832795", "3.1415926535897932384626433832796", false},
}

func TestGE(t *testing.T) {
	for i, a := range geTests {
		var x, y Float128
		if _, err := fmt.Sscan(a.x, &x); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		if _, err := fmt.Sscan(a.y, &y); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		r := IsGE(x, y)
		if r != a.r {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkGE(b *testing.B) {
	var x, y Float128
	b.StopTimer()
	fmt.Sscan("3.1415926535897932384626433832795", &x) // truncated
	fmt.Sscan("3.1415926535897932384626433832796", &y) // +1 LSD
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsGE(x, y)
	}
}

// GT

var gtTests = []struct {
	x, y string
	r    bool
}{
	{"-1.0000000000000000000000000000001e+200", "-1.0000000000000000000000000000000e+200", false},
	{"-1.0000000000000000000000000000000e+200", "-1.0000000000000000000000000000001e+200", true},
	{"-2", "-1", false},
	{"-1", "-2", true},
	{"-1.0000000000000000000000000000001e-200", "-1.0000000000000000000000000000000e-200", false},
	{"-1.0000000000000000000000000000000e-200", "-1.0000000000000000000000000000001e-200", true},

	{"0", "0", false},

	{"+1.0000000000000000000000000000001e-200", "+1.0000000000000000000000000000000e-200", true},
	{"+1.0000000000000000000000000000000e-200", "+1.0000000000000000000000000000001e-200", false},
	{"+2", "+1", true},
	{"+1", "+2", false},
	{"+1.0000000000000000000000000000001e+200", "+1.0000000000000000000000000000000e+200", true},
	{"+1.0000000000000000000000000000000e+200", "+1.0000000000000000000000000000001e+200", false},

	{"3.1415926535897932384626433832795", "3.1415926535897932384626433832796", false},
}

func TestGT(t *testing.T) {
	for i, a := range gtTests {
		var x, y Float128
		if _, err := fmt.Sscan(a.x, &x); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		if _, err := fmt.Sscan(a.y, &y); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		r := IsGT(x, y)
		if r != a.r {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkGT(b *testing.B) {
	var x, y Float128
	b.StopTimer()
	fmt.Sscan("3.1415926535897932384626433832795", &x) // truncated
	fmt.Sscan("3.1415926535897932384626433832796", &y) // +1 LSD
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsGT(x, y)
	}
}

// NE

var neTests = []struct {
	x, y string
	r    bool
}{
	{"-1.0000000000000000000000000000001e+200", "-1.0000000000000000000000000000000e+200", true},
	{"-1.0000000000000000000000000000000e+200", "-1.0000000000000000000000000000001e+200", true},
	{"-2", "-1", true},
	{"-1", "-2", true},
	{"-1.0000000000000000000000000000001e-200", "-1.0000000000000000000000000000000e-200", true},
	{"-1.0000000000000000000000000000000e-200", "-1.0000000000000000000000000000001e-200", true},

	{"0", "0", false},

	{"+1.0000000000000000000000000000001e-200", "+1.0000000000000000000000000000000e-200", true},
	{"+1.0000000000000000000000000000000e-200", "+1.0000000000000000000000000000001e-200", true},
	{"+2", "+1", true},
	{"+1", "+2", true},
	{"+1.0000000000000000000000000000001e+200", "+1.0000000000000000000000000000000e+200", true},
	{"+1.0000000000000000000000000000000e+200", "+1.0000000000000000000000000000001e+200", true},

	{"3.1415926535897932384626433832795", "3.1415926535897932384626433832796", true},
}

func TestNE(t *testing.T) {
	for i, a := range neTests {
		var x, y Float128
		if _, err := fmt.Sscan(a.x, &x); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		if _, err := fmt.Sscan(a.y, &y); err != nil {
			t.Errorf("scanf unsuccessful")
		}
		r := IsNE(x, y)
		if r != a.r {
			t.Errorf("#%d got r = %v; want %v", i, r, a.r)
		}
	}
}

func BenchmarkNE(b *testing.B) {
	var x, y Float128
	b.StopTimer()
	fmt.Sscan("3.1415926535897932384626433832795", &x) // truncated
	fmt.Sscan("3.1415926535897932384626433832796", &y) // +1 LSD
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsNE(x, y)
	}
}

// Mul

func TestMul(t *testing.T) {
	//cw := math.GetFPControl()
	//ncw := (cw &^ 0x300) | 0x200
	//math.SetFPControl(ncw)
	//sw := math.GetFPStatus()
	var a, b Float128
	a = SetFloat64(1)
	b = SetFloat64(3)
	for i := 0; i < 8; i++ {
		a.Mul(b)
		//fmt.Printf("%3d: %016x %016x, %+20.16e %+20.16e %v\n", i, math.Float64bits(a[0]), math.Float64bits(a[1]), a[0], a[1], a)
		//fmt.Printf("%3d: %016x %016x, (control=%08x newControl=%08x status=%08x) %+20.16e %+20.16e %v\n", i, math.Float64bits(a[0]), math.Float64bits(a[1]), cw, ncw, sw, a[0], a[1], a)
	}
	//math.SetFPControl(cw)
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < (b.N+99)/100; i++ {
		var a, b Float128
		a = SetFloat64(1)
		b = SetFloat64(3)
		for j := 0; j < 100; j++ {
			a.Mul(b)
		}
	}
}

func TestPower(t *testing.T) {
	//cw := math.GetFPControl()
	////ncw := (cw &^ 0x300) | 0x200
	//math.SetFPControl(ncw)
	var a Float128
	for i := int64(0); i <= 65; i++ {
		a = SetFloat64(3)
		a.PowerI(i)
		//a = PowerDI(SetFloat64(3), i)
		fmt.Printf("%3d: 3**%3d = %016x %016x, %+20.16e %+20.16e %v\n", i, i, math.Float64bits(a[0]), math.Float64bits(a[1]), a[0], a[1], a)
	}
	//math.SetFPControl(cw)
}