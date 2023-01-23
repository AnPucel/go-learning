package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

/*
When the benchmark is executed it runs b.N times
and measures how long it takes.The framework determines
a "good" value for that to let you have some decent results
Sample output:
BenchmarkRepeat-16      10288515               100.2 ns/op
Ran 10,000,000 times w/ average 100.2 nanoseconds per operation
NOTE by default Benchmarks are run sequentially
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
