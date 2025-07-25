# Types

Generic data structures and utilities for Go.

## Installation

```bash
go get github.com/intezya/types
```

## Iterator

Lazy iterator with method chaining for functional-style data processing.

### Creating Iterators

```go
// From slice
iter := types.IteratorFromSlice([]int{1, 2, 3, 4, 5})

// From iter.Seq
seq := slices.Values([]string{"a", "b", "c"})
iter := types.IteratorFromSeq(seq)
```

### Methods

```go
// Transform elements
iter.Map(func(x int) int { return x * 2 })

// Filter elements  
iter.Filter(func(x int) bool { return x > 3 })

// Reverse order
iter.Reverse()

// Execute for each element
iter.Each(func(x int) { fmt.Println(x) })

// Collect to slice
result := iter.Collect()

// Count elements
count := iter.Count()
count := iter.CountWithPredicate(func(x int) bool { return x > 5 })
```

### Type Conversion

```go
// Map to different type
stringIter := types.MapIterator(intIter, func(x int) string {
return strconv.Itoa(x)
})
```

### Example

```go
types.IteratorFromSlice([]int{1, 2, 3, 4}).
Reverse().
Map(func(x int) int { return x * x }).
Filter(func(x int) bool { return x%2 == 0 }).
Each(func(x int) { fmt.Println(x) })
// Output: 16, 4
```

## Set

Generic set implementation with common set operations.

### Creating Sets

```go
// Empty set
set := types.NewSet[int]()

// From slice
set := types.SetFromSlice([]string{"a", "b", "c"})
```

### Basic Operations

```go
set.Add("value")
set.Remove("value")
exists := set.Contains("value")
size := set.Size()
values := set.Values()
```

### Set Operations

```go
set1 := types.SetFromSlice([]int{1, 2, 3})
set2 := types.SetFromSlice([]int{3, 4, 5})

// Union (modifies set1)
set1.Union(set2) // set1 now contains {1, 2, 3, 4, 5}

// Intersection (modifies set2)
set1.Intersect(set2)

// Clone
clone := set.Clone()

// Comparison
equal := set1.Equal(set2)
isSubset := set1.IsSubsetOf(set2)
isSuperset := set1.IsSupersetOf(set2)
```

### Integration with Iterator

```go
set := types.SetFromSlice([]int{1, 2, 3, 4, 5})
result := set.Iter().
Filter(func(x int) bool { return x > 2 }).
Map(func(x int) int { return x * 2 }).
Collect()
```

## OneOf

Generic union type that can hold one of two possible types.

### Creating OneOf

```go
// Empty OneOf
oneOf := types.NewOneOf[string, int]()

// Set values
oneOf.SetT1("hello")
oneOf.SetT2(42)
```

### Getting Values

```go
// Check which type is present
isT1, isT2 := oneOf.Present()

// Get values with presence check
if str, ok := oneOf.GetT1(); ok {
    fmt.Println("String:", str)
}

if num, ok := oneOf.GetT2(); ok {
    fmt.Println("Number:", num)
}
```

### Example

```go
oneOf := types.NewOneOf[string, error]()

// Success case
oneOf.SetT1("operation successful")
if result, ok := oneOf.GetT1(); ok {
    fmt.Println(result)
}

// Error case
oneOf.SetT2(errors.New("something went wrong"))
if err, ok := oneOf.GetT2(); ok {
    fmt.Println("Error:", err)
}
```

## Requirements

- Go 1.22+
