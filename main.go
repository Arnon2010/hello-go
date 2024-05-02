package main
import "fmt"

func main() {
	a := make([] int, 10) //สร้าง array 'a' มีจำนวนสมาชิกเท่ากับ 10 เป็นค่าเริ่มต้น
	a = append(a, 10) //เพิ่ม array ตัวที่ 11 เท่ากับ 10
	a = append(a, 20)
	a[2] = 30 //กำหนดให้ array ตัวที่ 3 = 30
	b := "test"

	for k, v := range a { //range คือ หาจำนวน array
		fmt.Printf("a[%d] = %d\n", k, v) //k=key, v=value
	}

	for i := range a {
		fmt.Printf("a[%d] = %d, %s\n", i, a[i], b) // i = key
	}
}

func main2() {
	// fmt.Printf("hello สยาม .com\n")
	//p1("hello\n")
	//var a1, a2, a3 int
	// a1 := 10
	// a2 := 20
	// a3 := a1 + a2

	// fmt.Printf("a = %d\n" a3)

}
