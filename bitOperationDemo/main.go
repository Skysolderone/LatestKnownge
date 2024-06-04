package main

import "fmt"

const (
	Read   = 1 << iota // 1 << 0 -> 1
	Write              // 1 << 1 -> 2
	Delete             // 1 << 2 -> 4
)

const (
	RAW = Read | Write          // 1 | 2 -> 3
	RAD = Read | Delete         // 1 | 4 -> 5
	WAD = Write | Delete        // 2 | 4 -> 6
	RWD = Read | Write | Delete // 1 | 2 | 4 -> 7
)

func HasPermission(userPermissions, permission int) bool {
	return userPermissions&permission != 0
}

func AddPermission(userPermissions, permission int) int {
	return userPermissions | permission
}

func RemovePermission(userPermissions, permission int) int {
	return userPermissions &^ permission
}

func main() {
	fmt.Println(RemovePermission(RWD, Delete))
	fmt.Println(AddPermission(RAW, Delete))
	fmt.Println(HasPermission(RWD, Delete))
}
