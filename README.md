# Data Structures and Algorithms

This repository contains implementations of basic **searching and sorting algorithms**
using **Golang (Go)**.

The goal is to maintain **clear, readable implementations** of fundamental algorithms
commonly asked in technical interviews.


## Searching Algorithms

### Linear Search

Linear search checks each element in the list sequentially until the target value
is found or the list ends.

**Characteristics:**
- Works on unsorted data
- Simple implementation

**Time Complexity:** O(n)

### Binary Search

Binary search operates on a **sorted array** by repeatedly dividing the search
space into half.

**Characteristics:**
- Requires sorted input
- Faster than linear search for large datasets

**Time Complexity:** O(log n)

## Sorting Algorithms

### Bubble Sort

Bubble sort repeatedly compares adjacent elements and swaps them if they are in
the wrong order. After each pass, the largest element moves to the end of the array.

**Time Complexity:** O(n²)


### Selection Sort

Selection sort selects the minimum element from the unsorted portion of the array
and places it at the correct position.

**Time Complexity:** O(n²)

## Algorithm Comparison

| Algorithm        | Category | Sorted Input | Time Complexity |
|------------------|----------|--------------|-----------------|
| Linear Search    | Search   | No           | O(n)            |
| Binary Search    | Search   | Yes          | O(log n)        |
| Bubble Sort      | Sort     | No           | O(n²)           |
| Selection Sort   | Sort     | No           | O(n²)           |



## Language

- Go (Golang)


## Notes

- All implementations focus on correctness and clarity.
- No external libraries are used.
