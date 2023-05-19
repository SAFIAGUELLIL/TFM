package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	MAX_THREADS = 1000
	SIZE        = 10 // 10 2000 500
)

var (
	num_threads int
	matrix_a    [SIZE][SIZE]int
	matrix_b    [SIZE][SIZE]int
	matrix_c    [SIZE][SIZE]int
)

func print_matrix(matrix [SIZE][SIZE]int) {

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func multiply_matrices(thread_id int, wg *sync.WaitGroup) {
	defer wg.Done()

	start := thread_id * SIZE / num_threads     // Índice inicial
	end := (thread_id + 1) * SIZE / num_threads // Índice final

	if thread_id == num_threads-1 {
		end = SIZE
	}

	fmt.Printf("Thread %d calculando de %d a %d\n", thread_id, start, end)

	for i := start; i < end; i++ {
		for j := 0; j < SIZE; j++ {
			sum := 0
			for k := 0; k < SIZE; k++ {
				sum += matrix_a[i][k] * matrix_b[k][j]
			}
			matrix_c[i][j] = sum
		}
	}

}
func main() {

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <num_threads>\n", os.Args[0])
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[1])
	num_threads = num

	if num_threads > MAX_THREADS || err != nil {
		fmt.Printf("Error: Maximum number of threads is %d\n", MAX_THREADS)
		os.Exit(1)
	}

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			matrix_a[i][j] = rand.Intn(10) % 10
			matrix_b[i][j] = rand.Intn(10) % 10
			matrix_c[i][j] = 0
		}
	}

	// fmt.Printf("Matrix a:\n")
	// print_matrix(matrix_a)

	// fmt.Printf("Matrix b:\n")
	// print_matrix(matrix_b)

	if num_threads >= SIZE {
		fmt.Println("Error: The number of threads cannot be greater than or equal to the matrix size.")
		return
	}

	var wg sync.WaitGroup
	wg.Add(num_threads)
	for i := 0; i < num_threads; i++ {
		go multiply_matrices(i, &wg)
	}
	wg.Wait()

	fmt.Printf("Result:\n")
	print_matrix(matrix_c)
	fmt.Printf("Multiplyed %dx%d matrices using %d threads\n", SIZE, SIZE, num_threads)

}
