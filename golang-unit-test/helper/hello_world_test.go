package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldSandi(t *testing.T) {
	// Hello Wordl Memangil package dari helper
	result := HelloWorld("Sandi")

	if result != "Hello Sandi" {
		// error
		// Akan di hentikan dan tetap mengekesuki kode di bawah
		t.Fail()
	}

	// ini akan di eksekusi
	fmt.Println("Test Hello World Sandi Done")
}

func TestHelloWorldSalsa(t *testing.T) {
	result := HelloWorld("Salsa")

	if result != "Hello Salsa" {
		// error
		// Akan di hentikan tanpa mengekesuki kode di bawah
		t.FailNow()
	}
	// ini tidak akan di eksekusi
	fmt.Println("Test Hello World Salsa Done")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Salsa")

	if result != "Hello Salsa" {
		// error
		t.Fatal("Harusnya Hello Salsa")
	}

	fmt.Println("Done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Sandi")

	if result != "Hello Sandi" {
		// error
		t.Error("Harusnya Hello Sandi")
	}

	fmt.Println("Done")
}

// use library from github

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Sansal")
	assert.Equal(t, "Hello Sansal", result, "Result must be 'Hello Sansal'")
	fmt.Println("Test Hello World Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Sansal")
	require.Equal(t, "Hello Sansal", result, "Result must be 'Hello Sansal'")
	fmt.Println("Test Hello World Assert Done")
}

// Skip

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Program Tidak Bisa Berjalan Di Windows")
	}

	// Tidak Akan Di Eksekusi
	result := HelloWorld("sandi")
	require.Equal(t, "Hello sandi", result, "Result must be 'Hello sandi'")

}

// Before & AFter use test main

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Unit Test")

	m.Run()

	// After
	fmt.Println("After Execution")
}

// test sub test

func TestSubTest(t *testing.T) {
	t.Run("sandi", func(t *testing.T) {
		result := HelloWorld("sandi")
		require.Equal(t, "Hello sandi", result, "Result must be 'Hello sandi'")
	})
	t.Run("salsa", func(t *testing.T) {
		result := HelloWorld("sandi")
		require.Equal(t, "Hello salsa", result, "Result must be 'Hello salsa'")
	})
}
