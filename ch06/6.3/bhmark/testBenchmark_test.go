package main

import "testing"

var final int
func benchmarkfb1(b *testing.B,n int)  {
	var end int
	for i := 0; i < b.N; i++ {
		end = fb1(n)
	}
	final = end
}

func Benchmarkfb2(b *testing.B,n int)  {
	var end int
	for i := 0; i < b.N; i++ {
		end = fb2(n)
	}
	final = end
}

func Benchmarkfb3(b *testing.B,n int)  {
	var end int
	for i := 0; i < b.N; i++ {
		end = fb3(n)
	}
	final = end
}

func Benchmark50fb1(b *testing.B)  {
	benchmarkfb1(b,30)
}

func Benchmark50fb2(b *testing.B)  {
	Benchmarkfb2(b,30)
}

func Benchmark50fb3(b *testing.B)  {
	Benchmarkfb3(b,30)
}