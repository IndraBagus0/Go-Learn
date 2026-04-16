package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Indra")
	if result != "Hello Indra" {
		t.Fatal("Result is not match")
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Indra")
	assert.Equal(t, "Hello Indra", result, "Result must be 'Hello Indra'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Indra")
	require.Equal(t, "Hello Indra", result, "Result must be 'Hello Indra'")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Indra")
	require.Equal(t, "Hello Indra", result, "Result must be 'Hello Indra'")
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("Indra", func(t *testing.T) {
		result := HelloWorld("Indra")
		require.Equal(t, "Hello Indra", result, "Result must be 'Hello Indra'")
	})
	t.Run("Bagus", func(t *testing.T) {
		result := HelloWorld("Bagus")
		require.Equal(t, "Hello Bagus", result, "Result must be 'Hello Bagus'")
	})
	t.Run("Syah", func(t *testing.T) {
		result := HelloWorld("Syah")
		require.Equal(t, "Hello Syah", result, "Result must be 'Hello Syah'")
	})
	t.Run("Putra", func(t *testing.T) {
		result := HelloWorld("Putra")
		require.Equal(t, "Hello Putra", result, "Result must be 'Hello Putra'")
	})

}

func TestHelloWorldIndra(t *testing.T) {
	result := HelloWorld("Indra")

	if result != "Hello Indra" {
		t.Error("Result must be 'Hello Indra'")
	}

	fmt.Println("TestHelloWorldIndra Done")
}

func TestHelloWorldBagus(t *testing.T) {
	result := HelloWorld("Bagus")

	if result != "Hello Bagus" {
		t.Error("Result must be 'Hello Bagus'")
	}

	fmt.Println("TestHelloWorldBagus Done")
}

func TestHelloWorldSyah(t *testing.T) {
	result := HelloWorld("Syah")

	if result != "Hello Syah" {
		t.Error("Result must be 'Hello Syah'")
	}

	fmt.Println("TestHelloWorldSyah Done")
}

func TestHelloWorldPutra(t *testing.T) {
	result := HelloWorld("Putra")

	if result != "Hello Putra" {
		t.Error("Result must be 'Hello Putra'")
	}

	fmt.Println("TestHelloWorldPutra Done")
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Indra",
			request:  "Indra",
			expected: "Hello Indra",
		},
		{
			name:     "Bagus",
			request:  "Bagus",
			expected: "Hello Bagus",
		},
		{
			name:     "Syah",
			request:  "Syah",
			expected: "Hello Syah",
		},
		{
			name:     "Putra",
			request:  "Putra",
			expected: "Hello Putra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Indra")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Indra", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Indra")
		}
	})
	b.Run("Bagus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bagus")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Indra",
			request: "Indra",
		},
		{
			name:    "Bagus",
			request: "Bagus",
		},
		{
			name:    "IndraBagusSyah",
			request: "Indra Bagus Syah",
		},
		{
			name:    "Putra",
			request: "Putra Indra",
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
