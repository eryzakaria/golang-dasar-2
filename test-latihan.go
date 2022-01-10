package main

import "fmt"

// type Blocked func(string) string

// func errorNilai(name string, blocked Blocked) {
// 	if blocked("Anjing") == 2 {
// 		fmt.Println("Ada Kata Kasar", name)
// 	} else {
// 		fmt.Println("Kamu Anjing", name)
// 	}
// }

// func main() {
// 	blocked := func(string) int {
// 		return 1
// 	}

// 	errorNilai("Ery", blocked)
// 	errorNilai("Ery", func(name string) int {
// 		return 2
// 	})
// }

// func facktorialLoop(value int) int {
// 	total := 1
// 	for i := value; i > 0; i-- {
// 		total *= i
// 	}
// 	return total
// }

// func rekursifLoop(value int) int {
// 	if value == 1 {
// 		return 1
// 	} else {
// 		return value * rekursifLoop(value-1)
// 	}
// }

// func main() {

// 	rekursif := rekursifLoop(5)
// 	fmt.Println(facktorialLoop(5))
// 	fmt.Println(rekursif)
// }

//closure
// func sayHello(number ...int) int {
// 	total := 0
// 	for _, numbers := range number {
// 		total += numbers
// 	}
// 	return total
// }

// func main() {
// 	total := sayHello(10, 10, 10, 10, 10)
// 	fmt.Println("Hasil", total)

// 	slice := []int{10, 10, 10, 10, 10}
// 	total = sayHello(slice...)
// 	fmt.Println("Hasil", total)
// }

// func sayHello(name string, filter func(string) string) {
// 	fmt.Println("Hello", filter(name))
// }

// func sayfilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func main() {
// 	sayHello("Anjing", sayfilter)

// 	filter := sayfilter
// 	sayHello("Ery", filter)
// }

// func main() {
// 	closure := 0

// 	increment := func() {
// 		fmt.Println("closure")
// 		closure++
// 	}

// 	increment()

// 	fmt.Println(closure)
// }

// func FilterSpam(name string, blocked Blocked) {
// 	if name == "Domba" {
// 		fmt.Println("----")
// 	} else {
// 		fmt.Println("Welcome ", name)
// 	}
// }

// func main() {
// 	cek := func(string) string {
// 		return "Ery"
// 	}

// 	FilterSpam("Ery", cek)

// }

// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }

// func runApplication() {
// 	defer logging()
// 	fmt.Println("Run Aplication")
// }

// func main() {
// 	runApplication()
// }

// func endApp() {
// 	// panic
// 	// fmt.Println("Aplikasi Selesai")
// 	message := recover()
// 	if message != nil {
// 		fmt.Println("Error dengan Message:", message)
// 	}
// 	fmt.Println("Aplikasi Selesai")
// }

// // func runApp(error bool) {
// func runApp(error bool, number int) {
// 	// func runApp() {
// 	// text defer
// 	defer endApp()
// 	// ini panic
// 	if error {
// 		panic("Panic Error")
// 	}
// 	// jangan ditulis kesini
// 	// message := recover()
// 	// fmt.Println("Pesan Error", message)
// 	fmt.Println("Aplikasi berjalan")
// 	result := 10 / number
// 	fmt.Println("Hasil", result)

// 	// endApp()
// }

// func main() {
// 	runApp(true, 2)
// 	// fmt.Println("zaki")
// }

// type Customer struct {
// 	firstName string
// 	lastName  string
// 	address   string
// 	Age       int
// }

// func main() {
// 	customer := Customer{
// 		firstName: "Ery",
// 		lastName:  "Trimadhani",
// 		address:   "Tegal",
// 		Age:       22,
// 	}

// 	fmt.Println(customer)
// }

// func (customer Customer) sayHello() {
// 	fmt.Println("Hello, my name is ", customer.firstName)
// }

// func main() {
// 	firstName := Customer{firstName: "Ery"}
// 	firstName.sayHello()
// 	fmt.Println(Ups())
// }

// func Ups() interface{} {
// 	return "Ups"
// }

// func Pembagi(nilai int, pembagi int) (int, error) {
// 	if pembagi == 0 {
// 		return 0, errors.New("pembagian dengan nol")
// 	} else {
// 		return nilai / pembagi, nil
// 	}
// }

// func main() {
// 	nilai, err := Pembagi(100, 0)

// 	if err == nil {
// 		fmt.Println("Hasil", nilai)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}
// }

// func sayHello() interface{} {
// 	return true
// }

// func main() {
// 	// fmt.Println(sayHello())
// 	result := sayHello()
// 	// resultString := result.(string)

// 	// fmt.Println(resultString)
// 	switch value := result.(type) {
// 	case int:
// 		fmt.Println("int", value)
// 	case string:
// 		fmt.Println("string", value)
// 	default:
// 		fmt.Println("Unkown Result")
// 	}
// }

// type Address struct {
// 	city, province, country string
// }

// func main() {
// 	address1 := Address{"Tegal", "Jawa Tengah", "Indonesia"}
// 	address2 := &address1

// 	address2.city = "Bandung"
// 	fmt.Println(address1)
// 	fmt.Println(address2.city)

// address2 = &Address{"Bandung", "Jawa Barat", ""}
// fmt.Println(address2)

// 	address := Address{"Tegal", "Jawa Tengah", ""}
// 	ChangeAddress(&address)

// 	fmt.Println(address)

// }

// func ChangeAddress(address *Address) {
// 	address.country = "Indonesia"
// }

// type Customer struct {
// 	Name, Address string
// 	Age           int
// }

// func (customer Customer) sayHi() {
// 	fmt.Println("Hi, Welcome ", customer)
// }

// func main() {
// 	name := Customer{Name: "Ery"}
// 	name.sayHi()
// }

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	name := Man{"Ery"}
	name.Married()

	fmt.Println(name.Name)
}

// func random() interface{} {
// 	return "Ups"
// }

// func main() {
// 	result := random()
// 	switch value := result.(type) {
// 	case string:
// 		fmt.Println("Ini string", value)
// 	case int:
// 		fmt.Println("Ini int", value)
// 	default:
// 		fmt.Println("Unkown Error")
// 	}
// }
