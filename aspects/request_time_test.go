package ginmon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestNewRequestTimeAspect(values ...float64) *RequestTimeAspect {
	rt := NewRequestTimeAspect()
	for _, n := range values {
		rt.add(n)
	}
	rt.calculate()
	return rt
}

func TestRequestTimer_add(t *testing.T) {
	rt := newTestNewRequestTimeAspect(1.0, 2.0, 3.0) // this will reset len() to 0
	rt.add(5.0)
	expect := 1
	l := len(rt.lastMinuteRequestTimes)
	if assert.Equal(t, expect, l, "Adding a value does not work, expect %d but got %d %s",
		expect, l, ballotX) {
		t.Logf("Adding a value works, expect %d and git %d %s",
			expect, l, checkMark)
	}
}

func TestRequestTimer_GetStats(t *testing.T) {
	rt := newTestNewRequestTimeAspect(0.0)
	for i := 0; i < 100; i++ {
		rt.add(float64(i))
	}
	rt.calculate()

	if assert.NotNil(t, rt.GetStats(), "Return of GetStats should not be nil") {
		t.Logf("Should be an interface %s", checkMark)
	}

	epsilon := 0.01
	stat := rt.GetStats().(*RequestTimeAspect)

	if assert.InEpsilon(t, 99, stat.Max, epsilon, "Return of getstats should have a Max") {
		t.Logf("Should be 99 %s", checkMark)
	}
	if assert.Equal(t, float64(0), stat.Min, "Return of getstats should have a Min") {
		t.Logf("Should be 0 %s", checkMark)
	}
	if assert.InEpsilon(t, 50, stat.Mean, 0.5, "Return of getstats should have a Mean") {
		t.Logf("Should be 50 %s", checkMark)
	}
	if assert.InEpsilon(t, 29.01, stat.Stdev, epsilon, "Return of getstats should have a Stdev") {
		t.Logf("Should be 29.01 %s", checkMark)
	}
	if assert.InEpsilon(t, 90, stat.P90, epsilon, "Return of getstats should have a P90") {
		t.Logf("Should be 90 %s", checkMark)
	}
	if assert.InEpsilon(t, 95, stat.P95, epsilon, "Return of getstats should have a P95") {
		t.Logf("Should be 95 %s", checkMark)
	}
	if assert.InEpsilon(t, 99, stat.P99, epsilon, "Return of getstats should have a P99") {
		t.Logf("Should be 99 %s", checkMark)
	}
}

func (rt *RequestTimeAspect) createValues(n int) {
	for i := 0; i < n; i++ {
		rt.add(float64(i))
	}
}

func (rt *RequestTimeAspect) runBenchRequestTimerCalculate(i int) {
	rt.createValues(i)
	rt.calculate()
	stat := rt.GetStats().(*RequestTimeAspect)
	_ = stat.P95
}

func BenchmarkRequestTimerCalculate100(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runBenchRequestTimerCalculate(100)
	}
}

func BenchmarkRequestTimerCalculate1000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runBenchRequestTimerCalculate(1000)
	}
}

func BenchmarkRequestTimerCalculate10000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runBenchRequestTimerCalculate(10000)
	}
}

func BenchmarkRequestTimerCalculate100000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runBenchRequestTimerCalculate(100000)
	}
}

func BenchmarkRequestTimerCalculate360000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runBenchRequestTimerCalculate(360000)
	}
}

func (rt *RequestTimeAspect) runSliceAppendAndMemcopy(i int) {
	rt.createValues(i)
	a := rt.lastMinuteRequestTimes[:]
	rt.lastMinuteRequestTimes = make([]float64, 0)
	_ = a[5]
}

func BenchmarkSliceAppendAndMemcopy100(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runSliceAppendAndMemcopy(100)
	}
}

func BenchmarkSliceAppendAndMemcopy1000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runSliceAppendAndMemcopy(1000)
	}
}

func BenchmarkSliceAppendAndMemcopy10000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runSliceAppendAndMemcopy(10000)
	}
}

func BenchmarkSliceAppendAndMemcopy100000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runSliceAppendAndMemcopy(100000)
	}
}

func BenchmarkSliceAppendAndMemcopy1000000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runSliceAppendAndMemcopy(1000000)
	}
}

func BenchmarkSliceAppendAndMemcopy360000(b *testing.B) {
	rt := NewRequestTimeAspect()
	for n := 0; n < b.N; n++ {
		rt.runSliceAppendAndMemcopy(360000)
	}
}

func runSliceAppend(n int) {
	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
}

func BenchmarkSliceAppend100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runSliceAppend(100)
	}
}

func BenchmarkSliceAppend1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runSliceAppend(1000)
	}
}

func BenchmarkSliceAppend10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runSliceAppend(10000)
	}
}

func BenchmarkSliceAppend100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runSliceAppend(100000)
	}
}

func BenchmarkSliceAppend1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runSliceAppend(1000000)
	}
}

func TestRequestTimer_Name(t *testing.T) {
	rt := newTestNewRequestTimeAspect(1.0, 2.0, 3.0)
	expect := "RequestTime"
	if assert.Equal(t, expect, rt.Name(), "Return of counter name does not work, expect %s but got %s %s",
		expect, rt.Name(), ballotX) {
		t.Logf("Return of counter name works, expect %s and got %s %s",
			expect, rt.Name(), checkMark)
	}
}

func TestRequestTimer_InRoot(t *testing.T) {
	rt := newTestNewRequestTimeAspect(1.0, 2.0, 3.0)
	expect := false
	if assert.Equal(t, expect, rt.InRoot(), "Expect %v but got %v %s",
		expect, rt.InRoot(), ballotX) {
		t.Logf("Expect %v and got %v %s",
			expect, rt.InRoot(), checkMark)
	}
}
