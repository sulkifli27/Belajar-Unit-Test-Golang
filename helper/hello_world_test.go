package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B){
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name : "Sul",
			request: "Sul",
		},
		{
			name : "Sulkifli",
			request: "Sulkifli",
		},
		{
			name : "Ijul",
			request: "Ijul",
		},
		
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B){
	b.Run("Sulkifli", func (b *testing.B)  {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sulkifli")
		}
	})

	b.Run("Sul", func (b *testing.B)  {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sulkifli")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Sulkifli")
	}
}
func BenchmarkHelloWorldSul(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Sul")
	}
}

func TestTableHelloWorld(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "sul",
			request: "sul",
			expected: "hello sul",
		},
		{
			name: "kifli",
			request: "kifli",
			expected: "hello kifli",
		},
		{
			name: "ijul",
			request: "ijul",
			expected: "hello ijul",
		},
	}

	for _, test := range tests { 
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		}) 
	}
}

func TestSubTest(t *testing.T){
	t.Run("Sul", func(t *testing.T){
		result := HelloWorld("sul")
		require.Equal(t, "hello sul", result, "Result must be 'Hello sul'")
	})
	t.Run("Kifli", func(t *testing.T){
		result := HelloWorld("kifli")
		require.Equal(t, "hello kifli", result, "Result must be 'Hello kifli'")
	})
}


func TestMain(m *testing.M){
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")

}


func TestSkip(t *testing.T){
	if runtime.GOOS == "darwin"{
		t.Skip("Can not run on Mac Os")
	}

	result := HelloWorld("sul")
	require.Equal(t, "hello sul", result, "Result must be 'Hello sul'")

}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("sul")
	require.Equal(t, "hello sul", result, "Result must be 'Hello sul'")
	fmt.Println("Test Hello World Require Done")
}

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("sul")
	assert.Equal(t, "hello sul", result, "Result must be 'Hello sul'")
	fmt.Println("Test Hello World Assert Done")
}

func TestHelloWorldSul(t *testing.T){
	result := HelloWorld("sul")

	if result != "hello sul" {
		// error
		// t.Fail()
		t.Error("Result must be 'Hello sul'")
	}

	fmt.Println("TestingHelloWorld Done")
}


func TestHelloWorldSulkifli(t *testing.T){
	result := HelloWorld("sulkifli")

	if result != "hello sulkifli" {
		// error
		// t.FailNow()
		t.Fatal("Result must be 'Hello sulkifli'")

	}

	fmt.Println("TestingHelloWorldSul Done")

}