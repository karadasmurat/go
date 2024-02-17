package collections

// No built-in function (like Python's in operator) to directly check if a slice contains a specific value.
// Returning a tuple containing both the index and isFound flag can be a good idea for a Contains function,
// especially when you need to know both the existence and position of the element within the slice.
func Contains(s []int, q int) (int, bool) {

	for k, v := range s {
		if v == q {
			return k, true
		}
	}

	return 0, false
}
