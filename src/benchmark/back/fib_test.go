package fib_test
import "testing"
//import "fmt"

func fib(n int) int {
    if n < 2 {
        return n
    }

    return fib(n-1) + fib(n-2)
}

func BenchmarkFib10(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fib(10)
    }
}
