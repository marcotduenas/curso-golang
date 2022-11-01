<<<<<<< HEAD
package main

import "fmt"

func main(){
    array1 := [3]int{1, 2, 3}
    var slice1 []int

    slice1 = append(slice1, 4, 5, 6)
    fmt.Println(array1, slice1)

    slice2 := make([]int, 2, 3)
    copy(slice2, slice1)
    slice2 = append(slice2, 6, 7, 8)
    fmt.Println(slice2)
}
=======
package main

import "fmt"

func main(){
    array1 := [3]int{1, 2, 3}
    var slice1 []int

    slice1 = append(slice1, 4, 5, 6)
    fmt.Println(array1, slice1)

    slice2 := make([]int, 2, 3)
    copy(slice2, slice1)
    slice2 = append(slice2, 6, 7, 8)
    fmt.Println(slice2)
}
>>>>>>> be4a4b7a26d6b38a10d7f0fad28ec7684e092621
