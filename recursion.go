package main



import "fmt"


func sum(num int) int {
    if num == 1 {
        return 1
    }  else {
        return num + sum(num-1)
    }
}

func seperateDigits(num int) int {

    if num == 0 {
        return 0
    }
    digit := num % 10
   return  digit + seperateDigits(num/10)

}
func isDivisibleBy3(num int)  bool {

    if seperateDigits(num)%3 != 0 {
        return false
    }
    return true

}

func main() {
    fmt.Println("hi from new file")
    fmt.Println(sum(5))
    fmt.Println(isDivisibleBy3(63))
}
