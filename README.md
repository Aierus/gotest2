Andy Clark, Go Quiz 2

<b> instructions </b>
    `go run GoQuiz2.go {x}`

Where x is your input to the program

<b> expected outcome </b>
`
    - Your input is -> 1231
    - Summing a slice ***without*** concurrency of len 1231 took 542
    - The sum of the slice is 2814405775019430476
    - Summing a slice ***with*** concurrency of len 1231 took 360750
    - The sum of the slice is 2814405775019430476
    - sorting a slice of length 1231 using sort.Sort took 82792
    - sorting a slice using sort.SliceStable took 5750
`

<b> Question 3 </b>

The Big-O time analysis from the go documentation holds

When sorting slices (and therefore arrays after conversion), sort.Stable quadruples when inputs grow by an order of magnitude. These inputs align with the expected `O(n*n*log(n)*n*log(n)*log(n)) == O(3n * 3log(n))`

Sorting using sort.Sort() results in a longer time the larger the slices get, approximately doubling every time x grows an order of magnitude, concurring with the assumption that sort.Sort is O(log(n)) for slices. For arrays, we must first convert the array to a slice to use the sort interface, so we add an additional step n that increases the time needed to sort. After conversion the sort.Sort() time is the same as slices.
