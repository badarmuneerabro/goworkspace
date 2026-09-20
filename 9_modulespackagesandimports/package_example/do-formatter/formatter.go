// package format formats the data
package format

import "fmt"

// Number formats any type of int data
func Number(num int) string{
	return fmt.Sprintf("The number is %d", num)
}
