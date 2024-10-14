# L1 Test Tasks

This repository contains solutions to a set of test tasks organized in folders. Each task is implemented in a separate directory `L1.<task_number>` and named `task<task_number>.go`. The tasks cover various aspects of Golang programming, including concurrency, data structures, and algorithms.

<p align="center">
  <img src="/img/gopher.png" alt="Gopher" width="600" />
</p>

## Task List

### L1.1: Struct Inheritance
- **Description**: Implement embedding of methods from a `Human` structure into an `Action` structure to simulate inheritance.
- **Path**: `L1.1/task1.go`

### L1.2: Concurrent Squaring of Numbers
- **Description**: Concurrently compute and output the squares of numbers from the array `[2, 4, 6, 8, 10]`, and calculate the sum of their squares.
- **Path**: `L1.2/task2.go`

### L1.3: Worker Pool
- **Description**: Continuously write data to a channel and set up a pool of workers to read from the channel and output data. The number of workers is configurable at startup.
- **Path**: `L1.3/task3.go`

### L1.4: Graceful Shutdown
- **Description**: Program should terminate gracefully when Ctrl+C is pressed, with all workers shutting down properly.
- **Path**: `L1.4/task4.go`

### L1.5: Timed Channel Communication
- **Description**: Send values to a channel and read them on the other side. The program should terminate after a set number of seconds.
- **Path**: `L1.5/task5.go`

### L1.6: Goroutine Stopping Mechanisms
- **Description**: Implement various ways to stop goroutines.
- **Path**: `L1.6/task6.go`

### L1.7: Concurrent Map Writing
- **Description**: Perform concurrent writing to a map safely.
- **Path**: `L1.7/task7.go`

### L1.8: Bit Manipulation
- **Description**: Modify a given bit (set to 0 or 1) in an `int64` variable.
- **Path**: `L1.8/task8.go`

### L1.9: Number Pipeline
- **Description**: Create a pipeline where numbers from an array are doubled and passed through channels to be printed.
- **Path**: `L1.9/task9.go`

### L1.10: Temperature Grouping
- **Description**: Group temperature fluctuations into subsets with a step of 10 degrees.
- **Path**: `L1.10/task10.go`

### L1.11: Set Intersection
- **Description**: Implement the intersection of two unordered sets.
- **Path**: `L1.11/task11.go`

### L1.12: String Set
- **Description**: Create a set from a sequence of strings.
- **Path**: `L1.12/task12.go`

### L1.13: Swap Numbers
- **Description**: Swap two numbers without using a temporary variable.
- **Path**: `L1.13/task13.go`

### L1.14: Type Detection
- **Description**: Detect the type of a variable (`int`, `string`, `bool`, `channel`) at runtime from an `interface{}`.
- **Path**: `L1.14/task14.go`

### L1.15: String Manipulation Issue
- **Description**: Analyze the given code snippet and correct potential issues with string slicing.
- **Path**: `L1.15/task15.go`

### L1.16: Quicksort Implementation
- **Description**: Implement the quicksort algorithm without using built-in sort functions.
- **Path**: `L1.16/task16.go`

### L1.17: Binary Search
- **Description**: Implement a binary search algorithm.
- **Path**: `L1.17/task17.go`

### L1.18: Concurrent Counter
- **Description**: Create a counter that increments safely in a concurrent environment and outputs the final value upon completion.
- **Path**: `L1.18/task18.go`

### L1.19: String Reversal
- **Description**: Reverse an input string, including Unicode characters.
- **Path**: `L1.19/task19.go`

### L1.20: Word Reversal in String
- **Description**: Reverse the order of words in a string.
- **Path**: `L1.20/task20.go`

### L1.21: Adapter Pattern
- **Description**: Implement the Adapter design pattern.
- **Path**: `L1.21/task21.go`

### L1.22: Arithmetic on Large Numbers
- **Description**: Multiply, divide, add, and subtract two large numbers greater than `2^20`.
- **Path**: `L1.22/task22.go`

### L1.23: Remove Element from Slice
- **Description**: Remove the i-th element from a slice.
- **Path**: `L1.23/task23.go`

### L1.24: Distance Between Points
- **Description**: Compute the distance between two points, represented as a `Point` structure with encapsulated `x` and `y` fields.
- **Path**: `L1.24/task24.go`

### L1.25: Custom Sleep Function
- **Description**: Implement your own version of the `sleep` function.
- **Path**: `L1.25/task25.go`

### L1.26: Unique Characters in String
- **Description**: Check if all characters in a string are unique, case-insensitive.
- **Path**: `L1.26/task26.go`

---

## How to Run

1. Clone the repository.
    ```bash
   git clone git@github.com:mysticalien/L1.git
   ```
2. Navigate to the specific task directory (e.g., `L1.1`).
3. Run the Go file using the command:
   ```bash
   go run task<task_number>.go
   ```