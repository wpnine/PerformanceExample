package main

import (
	"testing"
)


func BenchmarkData1(b *testing.B)  {
	b.ResetTimer()
	var sum *int = new(int)
	*sum = 0

	toSum1(sum)

	b.StopTimer()

}

func BenchmarkData2(b *testing.B)  {
	b.ResetTimer()

	var sum *int = new(int)
	*sum = 0

	toSum2(sum)
	b.StopTimer()
}



func BenchmarkData3(b *testing.B)  {
	b.ResetTimer()

	var sum *int = new(int)
	*sum = 0

	toSum3(sum)
	b.StopTimer()
}

func BenchmarkData4(b *testing.B)  {
	b.ResetTimer()

	var sum *int = new(int)
	*sum = 0

	toSum4(sum)
	b.StopTimer()
}

func BenchmarkData5(b *testing.B)  {
	b.ResetTimer()

	var sum *int = new(int)
	*sum = 0

	toSum5(sum)
	b.StopTimer()
}

func BenchmarkData6(b *testing.B)  {
	b.ResetTimer()

	var sum *int = new(int)
	*sum = 0
	toSum6(sum)
	b.StopTimer()
}

