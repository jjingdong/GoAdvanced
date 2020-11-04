import "fmt"
import "math"

rand.Intn(10)
math.Sqrt(7)
math.Pi

func add(x int, y int) int{
    return x+y
}

func add(x, y int) int{
    return x+y
}

add(42,13)

func swap(x,y string)(string, string){
    return y,x
}

a,b := swap('hello', 'world')
fmt.PrintLn(a,b)

func split(sum int) (x,y int){
    x = sum * 4 / 9
    y = sum - x
    return
}

func main(){
    fmt.Println(split(17)
}

var c, python, java bool
func main(){
    var i int
    fmt.Println(i, c, python, java)
}

var i,j int = 1,2
func main(){
    var c, python, java = true, false, 'no!'
    fmt.Println(i,j,c,python,java)
}

func main(){
    var i, j int = 1,2
    k := 3
    c, python, java := true, false, 'no!'

    fmt.Println(i,j,k,c,python,java)
}

func main(){
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/*
bool
string
int int8 int16 int32 int64
unit unit8 unit16 unit32 unitptr
byte //alias for unit8
rune //alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128
*/

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe bool = false
    MaxInt unit64 = 1<<64 -1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
var i int = 42
i :=42
var f float64 = float64(i)
f := float64(i)
var u unit = unit(f)
u := unit(f)
*/

func main(){
    var x,y int = 3,4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z unit = unit(f)
    fmt.Println(x,y,z)
}

/*
var i int
j := i // j is an int

i := 42 //int
f := 3.142 //float64
g := 0.867 + 0.5i //complex128
*/

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}










