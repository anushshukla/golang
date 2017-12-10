// for code reusability of this file
package main


import (
  "fmt"
  "net/http"
  "reflect"
)

func pointerExample() {
  i, j := 42, 2701
  fmt.Println("Original Value of i:", i)
  fmt.Println("Original Value of j:", j)

  p := &i // point to i
  fmt.Println("Value of p which is pointed to relative address of i (&i):", p) // read i through the pointer
  *p = 21 // set i through the pointer
  fmt.Println("Value of i after value of i is overwritten via pointer p (*p):", i)
  fmt.Println("Value of p after value of i is overwritten via pointer p (*p):", p)

  p = &j // point to j
  fmt.Println("Value of p which is pointed to relative address of j (&j):", p)
  *p = *p / 37   // divide j through the pointer
  fmt.Println("Value of j after value of j is overwritten via arithmetic operation on pointer p (*p)):", j)

  // use of pointers
  // var sum int
  // fmt.Println("Value of adding j and i is :", add(i, j, &sum))
}

func add(x int, y int, sum *int) {
  // sum = x + y;
}

type Person struct {  
    firstName string
    lastName  string
}

func changeNameViaPointer(p *Person) {  
  p.firstName = "Bob"
}

func changeNameViaValueRefWithoutReturn(p Person) {  
  p.firstName = "Bob"
}

func changeNameViaValueRefWithReturn(p Person) {  
  p.firstName = "Bob"
  // return p;
}

func helloWorld() {
  var firstName string = "Anush"
  // variable declared and defined (by golang) as string whose datatype cannot be overriden
  lastName := "Shukla"
  // below line will throw an error exception as this variable was already declared
  //lastName := "Shukla"
  // below line will overwrite the value
  //lastName = "Shukla"
  // below line will throw an error due to datatype difference
  //lastName = 123
  fmt.Println("Hello", firstName, lastName);
}

func slicing() {
  primes := [6]int{2, 3, 5, 7, 11, 13}

  var s []int = primes[1:4]
  fmt.Println("Data type of primes is ", reflect.TypeOf(primes), "(though its an array)")
  fmt.Println("0th index of primes has value ", primes[0])
  fmt.Println(s)

}

func startServer() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
  })
  http.HandleFunc("localhost", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, improved world!")
  })
  //enter http://localhost:8080 on the address bar of any Internet Browser
  http.ListenAndServe(":8080", nil)
}

func poninterPracticalUse() {
  person := Person {
        firstName: "Alice",
        lastName: "Dow",
    }
  fmt.Println("Org Person struct before any name change",person)
  changeNameViaValueRefWithoutReturn(person)
  fmt.Println("Person struct after changeNameViaValueRefWithoutReturn",person)
  /*person = changeNameViaValueRefWithReturn(person)
  fmt.Println("Person struct after changeNameViaValueRefWithReturn",person)*/
  changeNameViaPointer(&person)
  fmt.Println("Person struct after changeNameViaPointer",person)
}

func main() {
  helloWorld();
  pointerExample();
  poninterPracticalUse();
  slicing();
  // startServer();
}