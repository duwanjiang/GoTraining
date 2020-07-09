/**
 * @description:
 * @author Administrator
 * @date 2020/7/10 0010 0:12
 */
package stringutil

func Reverse(s string) string {
	return reverseTwo(s)
}

/*
go build
	go build reverse.go reverseTwo.go
 	won't produce an output file.

go install
 	will place the package inside the pkg directory of the workspace.
*/
