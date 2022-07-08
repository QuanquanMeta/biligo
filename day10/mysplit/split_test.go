package mysplit

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		// failed
		t.Fatalf("want:%v but got:%v", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split(":a:b:c", ":")
	want := []string{"", "a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		// failed
		t.Fatalf("want:%v but got:%v", want, got)
	}
}

// test group
func TestSplitGroup(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := []testCase{
		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		testCase{",a,b,c", ",", []string{"", "a", "b", "c"}},
	}

	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%#v got:%#v\n", tc.want, got)
		}
	}
}

// test group
func TestSplitSub(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case1": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case2": testCase{",a,b,c", ",", []string{"", "a", "b", "c"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}

// test coverage

// Benchmark test
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

// Fib
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func benchmarFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarFib(b, 10) }
func BenchmarkFib30(b *testing.B) { benchmarFib(b, 30) }
