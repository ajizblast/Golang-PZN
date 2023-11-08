package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Aji",
			request: "Aji",
		},
		{
			name:    "Chahyo",
			request: "Chahyo",
		},
		{
			name:    "ChahyoPurnomoAJi",
			request: "Chahyo Purnomo Aji",
		},
		{
			name:    "Azizah",
			request: "Azizah Refa",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				helloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Aji", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helloWorld("Aji")
		}
	})
	b.Run("Chahyo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			helloWorld("Chahyo")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld("Aji")
	}
}

func BenchmarkHelloWorldChahyo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld("Chahyo")
	}
}

func TestTableHelloWorld(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Aji",
			request:  "Aji",
			expected: "Hello Aji",
		},
		{
			name:     "Chahyo",
			request:  "Chahyo",
			expected: "Hello Chahyo",
		},
		{
			name:     "Purnomo",
			request:  "Purnomo",
			expected: "Hello Purnomo",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := helloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Aji", func(t *testing.T) {
		result := helloWorld("Aji")
		require.Equal(t, "Hello Aji", result, "Result must be Hello Aji")
	})
	t.Run("Chahyo", func(t *testing.T) {
		result := helloWorld("Chahyo")
		require.Equal(t, "Hello Chahyo", result, "Result must be Hello Chahyo")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before Unit Test")

	m.Run()

	//after
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows")
	}

	result := helloWorld("Aji")
	assert.Equal(t, "Hello Aji", result, "Result must be Hello Aji")
}

func TestHelloWorldAssert(t *testing.T) {
	result := helloWorld("Aji")
	assert.Equal(t, "Hello Aji", result, "Result must be Hello Aji")
	fmt.Println("TestHelloWorldAsset with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := helloWorld("Aji")
	require.Equal(t, "Hello Aji", result, "Result must be Hello Aji")
	fmt.Println("TestHelloWorldAsset with Require Done")
}

func TestHelloWorldAji(t *testing.T) {
	result := helloWorld("Aji")

	if result != "Hello Aji" {
		//error
		// panic("Result is not Hello Aji")
		// t.Fail()
		t.Error("Result must be 'Hello Aji'")
	}

	fmt.Println("TestHelloWorldAji Done")
}

func TestHelloWorldChahyo(t *testing.T) {
	result := helloWorld("Chahyo")

	if result != "Hello Chahyo" {
		//error
		// panic("Result is not Hello Aji")
		// t.FailNow()
		t.Fatal("Result must be 'Hello Chahyo'")
	}

	fmt.Println("TestHelloWorldChahyo Done")
}
