package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Metrics struct {
	Lines     int
	Functions int
	Imports   int
	Comments  int
}

func analyzeFile(path string) Metrics {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	metrics := Metrics{}

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		metrics.Lines++

		if strings.HasPrefix(line, "func ") ||
			strings.HasPrefix(line, "def ") {

			metrics.Functions++
		}

		if strings.HasPrefix(line, "import ") {
			metrics.Imports++
		}

		if strings.HasPrefix(line, "//") ||
			strings.HasPrefix(line, "#") {

			metrics.Comments++
		}
	}

	return metrics
}

func compare(goFile string, pyFile string) {

	goMetrics := analyzeFile(goFile)
	pyMetrics := analyzeFile(pyFile)

	fmt.Println("===== GO =====")
	fmt.Printf("%+v\n", goMetrics)

	fmt.Println()

	fmt.Println("===== PYTHON =====")
	fmt.Printf("%+v\n", pyMetrics)

	fmt.Println()

	fmt.Println("===== DIFERENÇA =====")

	fmt.Printf(
		"Linhas: %d\n",
		goMetrics.Lines-pyMetrics.Lines,
	)

	fmt.Printf(
		"Funções: %d\n",
		goMetrics.Functions-pyMetrics.Functions,
	)
}

func main() {

	compare(
		"samples/algo.go",
		"samples/algo.py",
	)
}
