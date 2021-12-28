package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//Matrix is used as a type of 2d array
type Matrix [SIZE][SIZE]int

var (
	regularTime time.Duration
	//Solution1Time Naive-based execution time
	Solution1Time time.Duration
	//Solution2Time Strassen-based execution time
	Solution2Time time.Duration
	//Solution3Time Cannon-based execution time
	Solution3Time time.Duration
	bestTime      time.Duration
)

const (
	//SIZE of matrix ie SIZE x SIZE
	SIZE = 12
	//RAND used for random matrix generation
	RAND = 100
)

func rowCount(inM Matrix) int {
	rc := len(inM)
	return rc
}

func colCount(inM Matrix) int {
	cc := len(inM[0])
	return cc
}

func printMat(inM Matrix, name string) {
	fmt.Println("Matrix: ", name)
	for _, i := range inM {
		for _, j := range i {
			fmt.Print(" ", j)
		}
		fmt.Println()
	}
	fmt.Println()
	rc := rowCount(inM)
	fmt.Println("Row Count: ", rc)
	cc := colCount(inM)
	fmt.Println("Column Count: ", cc)
	fmt.Println()
}

func printIntro() {
	fmt.Println()
	fmt.Println("A long time ago in a galaxy far, far away....")
	fmt.Println()
	fmt.Println("             ┌───── •✧✧• ─────┐             ")
	fmt.Println("-------------- CONCURRENT WARS ---------------")
	fmt.Println("             └───── •✧✧• ─────┘             ​​​​​")
	fmt.Println("                   _.=+._	")
	fmt.Println("          :.\\`--._/[_/~|;\\_.--'/.:::	")
	fmt.Println("          ::.`.  ` __`\\.-.(  .'.::::	")
	fmt.Println("          ::::.`-:.`'..`-'/\\'.::::::	")
	fmt.Println("          :::::::.\\ `--')/  ) ::::::	")
	fmt.Println("                    `--'	")
	fmt.Println()
	fmt.Println("By:  Ayoub Jdair | Luke Kellett Murray")
	fmt.Println("Student IDs: 18266401 | 18250785")
	fmt.Println("Going for Grade: A1")
	fmt.Println()
}

func printOutro() {
	fmt.Println("             ┌───── •✧✧• ─────┐             ")
	fmt.Println("-------------      THE END     -------------")
	fmt.Println("             └───── •✧✧• ─────┘             ​​​​​")
	fmt.Println()
	fmt.Println("                   _.=+._	")
	fmt.Println("          :.\\`--._/[_/~|;\\_.--'/.:::	")
	fmt.Println("          ::.`.  ` __`\\.-.(  .'.::::	")
	fmt.Println("          ::::.`-:.`'..`-'/\\'.::::::	")
	fmt.Println("          :::::::.\\ `--')/  ) ::::::	")
	fmt.Println("                    `--'	")
	fmt.Println()
	fmt.Println("Written and Directed by George Lucas")
	fmt.Println("Student IDs: 18266401 | 18250785")
	fmt.Println("Going for Grade: A1")
	fmt.Println()
}

func printResults(procs int) {
	println()
	println("The results are in!")
	println()
	println("GOMAXPROCS set to: ", procs)
	println()
	println("Solution 0 time: ", regularTime.Microseconds(), "Microseconds")
	println("Solution 1 time: ", Solution1Time.Microseconds(), "Microseconds")
	println("Solution 2 time: ", Solution2Time.Microseconds(), "Microseconds")
	println("Solution 3 time: ", Solution3Time.Microseconds(), "Microseconds")
	println()
}

func createRandomMatrix() (Matrix, Matrix) {
	var a Matrix
	var b Matrix

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			a[i][j] = rand.Intn(RAND)
			b[i][j] = rand.Intn(RAND)
		}
	}
	return a, b
}

//Solution0 Based on non-concurrent method provided
func Solution0(inA Matrix, inB Matrix, procs int) Matrix {

	fmt.Println("Solution 0: Calculating Non-Concurrently")
	runtime.GOMAXPROCS(procs)
	var i, j int
	k := 0
	total := 0
	var nM Matrix
	start := time.Now()

	for i = 0; i < SIZE; i++ {
		for j = 0; j < SIZE; j++ {
			for k = 0; k < SIZE; k++ {
				total = total + inA[i][k]*inB[k][j]
			}
			nM[i][j] = total
			total = 0
		}
	}
	fmt.Println()
	elapsed := time.Since(start)
	regularTime = elapsed
	fmt.Printf("Solution 0: Time taken to calculate Non-Concurrently %s ", elapsed)
	fmt.Println()
	fmt.Println("Calculation Result = ")
	printMat(nM, "C")
	fmt.Println()
	return nM
}

//Solution1 Based on naive method
func Solution1(inA Matrix, inB Matrix, procs int) Matrix {

	fmt.Println("Solution 1: Calculating Concurrently")
	runtime.GOMAXPROCS(procs)

	var i, j int
	k := 0
	total := 0
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	var nM Matrix
	start := time.Now()

	for i = 0; i < SIZE; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("Go Routine: ", i)
			for j = 0; j < SIZE; j++ {
				for k = 0; k < SIZE; k++ {
					mutex.Lock()
					total = total + inA[i][k]*inB[k][j]
					mutex.Unlock()
				}
				mutex.Lock()
				nM[i][j] = total
				total = 0
				mutex.Unlock()
			}
			wg.Done()
		}(i)
	}
	fmt.Println()

	wg.Wait()
	elapsed := time.Since(start)
	Solution1Time = elapsed
	fmt.Printf("Solution 1: Time taken to calculate concurrently %s ", elapsed)
	fmt.Println("Solution 1: Calculation Result = ")
	printMat(nM, "C")
	fmt.Println()
	return nM
}

//Based on Strassen Method
func multiply(A Matrix, B Matrix) Matrix {
	var n Matrix

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			sum := 0
			for k := 0; k < SIZE; k++ {
				sum += A[i][k] * B[k][j]
			}
			n[i][j] = sum
		}
	}
	return n
}

func multiply2(A [SIZE / 2][SIZE / 2]int, B [SIZE / 2][SIZE / 2]int) [SIZE / 2][SIZE / 2]int {
	var n [SIZE / 2][SIZE / 2]int
	var wg sync.WaitGroup
	var mutex sync.RWMutex

	for i := 0; i < SIZE/2; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < SIZE/2; j++ {
				sum := 0
				for k := 0; k < SIZE/2; k++ {
					mutex.Lock()
					sum += A[i][k] * B[k][j]
					mutex.Unlock()
				}
				n[i][j] = sum
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return n
}

func add(A [SIZE / 2][SIZE / 2]int, B [SIZE / 2][SIZE / 2]int) [SIZE / 2][SIZE / 2]int {
	var n [SIZE / 2][SIZE / 2]int
	var wg sync.WaitGroup

	for i := 0; i < SIZE/2; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < SIZE/2; j++ {
				n[i][j] += A[i][j] + B[i][j]
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return n
}

func sub(A [SIZE / 2][SIZE / 2]int, B [SIZE / 2][SIZE / 2]int) [SIZE / 2][SIZE / 2]int {
	var n [SIZE / 2][SIZE / 2]int
	var wg sync.WaitGroup

	for i := 0; i < SIZE/2; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < SIZE/2; j++ {
				n[i][j] += A[i][j] - B[i][j]
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return n
}

//Solution2 Based on Strassen divide and conquer
func Solution2(A Matrix, B Matrix, procs int) Matrix {

	fmt.Println("Solution 2: Calculating Concurrently")
	runtime.GOMAXPROCS(procs)

	var n Matrix
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	var A11 [SIZE / 2][SIZE / 2]int
	var A12 [SIZE / 2][SIZE / 2]int
	var A21 [SIZE / 2][SIZE / 2]int
	var A22 [SIZE / 2][SIZE / 2]int
	var B11 [SIZE / 2][SIZE / 2]int
	var B12 [SIZE / 2][SIZE / 2]int
	var B21 [SIZE / 2][SIZE / 2]int
	var B22 [SIZE / 2][SIZE / 2]int

	start := time.Now()
	for i := 0; i < SIZE/2; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < SIZE/2; j++ {
				A11[i][j] = A[i][j]
				A12[i][j] = A[i][j+SIZE/2]
				A21[i][j] = A[i+SIZE/2][j]
				A22[i][j] = A[i+SIZE/2][j+SIZE/2]
				B11[i][j] = B[i][j]
				B12[i][j] = B[i][j+SIZE/2]
				B21[i][j] = B[i+SIZE/2][j]
				B22[i][j] = B[i+SIZE/2][j+SIZE/2]
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	M1 := multiply2(add(A11, A22), add(B11, B22))
	M2 := multiply2(add(A21, A22), B11)
	M3 := multiply2(A11, sub(B12, B22))
	M4 := multiply2(A22, sub(B21, B11))
	M5 := multiply2(add(A11, A12), B22)
	M6 := multiply2(sub(A21, A11), add(B11, B12))
	M7 := multiply2(sub(A12, A22), add(B21, B22))

	for i := 0; i < SIZE/2; i++ {
		wg2.Add(1)
		go func(i int) {
			for j := 0; j < SIZE/2; j++ {
				n[i][j] = M1[i][j] + M4[i][j] - M5[i][j] + M7[i][j]
				n[i][j+SIZE/2] = M3[i][j] + M5[i][j]
				n[i+SIZE/2][j] = M2[i][j] + M4[i][j]
				n[i+SIZE/2][j+SIZE/2] = M1[i][j] - M2[i][j] + M3[i][j] + M6[i][j]
			}
			wg2.Done()
		}(i)
	}
	wg2.Wait()
	elapsed := time.Since(start)
	Solution2Time = elapsed
	fmt.Printf("Solution 2: Time taken to calculate concurrently %s ", elapsed)
	fmt.Println("Solution 2: Calculation Result = ")
	printMat(n, "C")
	fmt.Println()
	return n
}

//ShiftL used on rows to move left with wrap-around
func ShiftL(matrix Matrix, i, count int) {
	ind := 0

	for ind < count {
		temp := matrix[i][0]
		indl := len(matrix[i]) - 1
		for j := 0; j < indl; j++ {
			matrix[i][j] = matrix[i][j+1]
			matrix[i][indl] = temp
			ind++
		}
	}
}

//ShiftU used for columns to shift up with wrap-around
func ShiftU(matrix Matrix, j, count int) {

	ind := 0

	for i := 0; i < count; i++ {
		temp := matrix[0][j]
		indl := len(matrix) - 1
		for i := 0; i < indl; i++ {
			matrix[i][j] = matrix[i+1][j]
			matrix[indl][j] = temp
			ind++
		}
	}
}

//Solution3 Based on Cannon algorithm
func Solution3(A Matrix, B Matrix, procs int) Matrix {

	println("Solution 3: Calculating Concurrently")

	runtime.GOMAXPROCS(procs)
	var nM Matrix
	start := time.Now()

	var wg sync.WaitGroup
	var mutex sync.RWMutex
	for k := 0; k < len(A); k++ {
		wg.Add(1)
		go func(k int) {
			fmt.Println("Go Routine: ", k)
			for i := 0; i < len(A); i++ {
				for j := 0; j < len(B); j++ {
					mutex.Lock()
					m := (i + j + k) % len(B)
					nM[i][j] += A[i][m] * B[m][j]
					mutex.Unlock()
					ShiftL(A, i, 1)
					ShiftU(B, j, 1)
				}
			}
			wg.Done()
		}(k)
	}
	wg.Wait()
	elapsed := time.Since(start)
	Solution3Time = elapsed
	fmt.Printf("Solution 3: Time taken to calculate concurrently %s ", elapsed)
	fmt.Println("Solution 3: Calculation Result = ")
	printMat(nM, "C")
	fmt.Println()
	return nM
}

func main() {

	printIntro()
	a, b := createRandomMatrix()
	procs := 2
	printMat(a, "A")
	printMat(b, "B")
	Solution0(a, b, procs)
	Solution1(a, b, procs)
	Solution2(a, b, procs)
	Solution3(a, b, procs)
	printResults(procs)
	printOutro()
}
