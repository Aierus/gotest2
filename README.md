Andy Clark, Go Quiz 2

The Big-O time analysis from the go documentation holds

When sorting slices (and therefore arrays after conversion), sort.Stable quadruples when inputs grow by an order of magnitude. These inputs align with the expected `O(n*n*log(n)*n*log(n)*log(n)) == O(3n * 3log(n))`

Sorting using sort.Sort() results in a longer time the larger the slices get, approximately doubling every time x grows an order of magnitude, concurring with the assumption that sort.Sort is O(log(n)) for slices. For arrays, we must first convert the array to a slice to use the sort interface, so we add an additional step n that increases the time needed to sort. After conversion the sort.Sort() time is the same as slices.
