package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== aict Performance Benchmarks ===")
	fmt.Println()

	dir := setupTestData()
	defer os.RemoveAll(dir)

	benchmarkLS(dir)
	benchmarkGrep(dir)
	benchmarkFind(dir)
	benchmarkCat(dir)
	benchmarkDiff(dir)
}

func setupTestData() string {
	dir, _ := os.MkdirTemp("", "aict-bench")

	for i := 0; i < 1000; i++ {
		os.WriteFile(filepath.Join(dir, fmt.Sprintf("file_%d.go", i)), []byte(fmt.Sprintf("package test%d\nfunc Foo() {}\n", i)), 0644)
	}

	subdir := filepath.Join(dir, "src", "pkg", "util")
	os.MkdirAll(subdir, 0755)
	for i := 0; i < 100; i++ {
		os.WriteFile(filepath.Join(subdir, fmt.Sprintf("deep_%d.txt", i)), []byte(strings.Repeat("line\n", 1000)), 0644)
	}

	largeFile := filepath.Join(dir, "large_file.txt")
	f, _ := os.Create(largeFile)
	for i := 0; i < 100000; i++ {
		fmt.Fprintf(f, "line %d: some search text here\n", i)
	}
	f.Close()

	return dir
}

func benchmarkLS(dir string) {
	fmt.Println("--- ls on 1000-file directory ---")

	start := time.Now()
	exec.Command("ls", dir).Run()
	gnuTime := time.Since(start)

	start = time.Now()
	exec.Command("./aict", "ls", dir, "--plain").Run()
	aictTime := time.Since(start)

	ratio := float64(aictTime) / float64(gnuTime)
	fmt.Printf("GNU ls:   %v\n", gnuTime)
	fmt.Printf("aict ls:  %v\n", aictTime)
	fmt.Printf("Ratio:    %.2fx\n", ratio)
	if ratio < 10 {
		fmt.Printf("Status:   ✅ PASS (<10x)\n")
	} else {
		fmt.Printf("Status:   ❌ FAIL (>10x)\n")
	}
	fmt.Println()
}

func benchmarkGrep(dir string) {
	fmt.Println("--- grep on 100k-line file ---")

	largeFile := filepath.Join(dir, "large_file.txt")

	start := time.Now()
	exec.Command("grep", "search", largeFile).Run()
	gnuTime := time.Since(start)

	start = time.Now()
	exec.Command("./aict", "grep", "search", largeFile, "--plain").Run()
	aictTime := time.Since(start)

	ratio := float64(aictTime) / float64(gnuTime)
	fmt.Printf("GNU grep:   %v\n", gnuTime)
	fmt.Printf("aict grep:  %v\n", aictTime)
	fmt.Printf("Ratio:      %.2fx\n", ratio)
	if ratio < 10 {
		fmt.Printf("Status:     ✅ PASS (<10x)\n")
	} else {
		fmt.Printf("Status:     ❌ FAIL (>10x)\n")
	}
	fmt.Println()
}

func benchmarkFind(dir string) {
	fmt.Println("--- find on deep directory tree ---")

	start := time.Now()
	exec.Command("find", dir, "-name", "*.go").Run()
	gnuTime := time.Since(start)

	start = time.Now()
	exec.Command("./aict", "find", dir, "-name", "*.go", "--plain").Run()
	aictTime := time.Since(start)

	ratio := float64(aictTime) / float64(gnuTime)
	fmt.Printf("GNU find:   %v\n", gnuTime)
	fmt.Printf("aict find:  %v\n", aictTime)
	fmt.Printf("Ratio:      %.2fx\n", ratio)
	if ratio < 10 {
		fmt.Printf("Status:     ✅ PASS (<10x)\n")
	} else {
		fmt.Printf("Status:     ❌ FAIL (>10x)\n")
	}
	fmt.Println()
}

func benchmarkCat(dir string) {
	fmt.Println("--- cat on 100k-line file ---")

	largeFile := filepath.Join(dir, "large_file.txt")

	start := time.Now()
	exec.Command("cat", largeFile).Run()
	gnuTime := time.Since(start)

	start = time.Now()
	exec.Command("./aict", "cat", largeFile, "--plain").Run()
	aictTime := time.Since(start)

	ratio := float64(aictTime) / float64(gnuTime)
	fmt.Printf("GNU cat:   %v\n", gnuTime)
	fmt.Printf("aict cat:  %v\n", aictTime)
	fmt.Printf("Ratio:    %.2fx\n", ratio)
	if ratio < 10 {
		fmt.Printf("Status:   ✅ PASS (<10x)\n")
	} else {
		fmt.Printf("Status:   ❌ FAIL (>10x)\n")
	}
	fmt.Println()
}

func benchmarkDiff(dir string) {
	fmt.Println("--- diff on similar files ---")

	f1, _ := os.Create(filepath.Join(dir, "f1.txt"))
	f2, _ := os.Create(filepath.Join(dir, "f2.txt"))

	for i := 0; i < 1000; i++ {
		fmt.Fprintf(f1, "line %d: content\n", i)
		fmt.Fprintf(f2, "line %d: content\n", i)
	}
	fmt.Fprintf(f2, "added line\n")
	f1.Close()
	f2.Close()

	start := time.Now()
	exec.Command("diff", filepath.Join(dir, "f1.txt"), filepath.Join(dir, "f2.txt")).Run()
	gnuTime := time.Since(start)

	start = time.Now()
	exec.Command("./aict", "diff", filepath.Join(dir, "f1.txt"), filepath.Join(dir, "f2.txt"), "--plain").Run()
	aictTime := time.Since(start)

	ratio := float64(aictTime) / float64(gnuTime)
	fmt.Printf("GNU diff:   %v\n", gnuTime)
	fmt.Printf("aict diff:  %v\n", aictTime)
	fmt.Printf("Ratio:      %.2fx\n", ratio)
	if ratio < 10 {
		fmt.Printf("Status:     ✅ PASS (<10x)\n")
	} else {
		fmt.Printf("Status:     ❌ FAIL (>10x)\n")
	}
	fmt.Println()
}
