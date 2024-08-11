
### About Go

- Compiled
- Technical focus on high performance and concurrency
- ~ to C & Python in syntax
- Type safe

### First the lang basics

- Packages
    - Export / import packages
      - By convention, for creating a locally available package, put it in a folder with the same name as the package
      - `package <name>` - package declaration
      - `import "fmt"` - import a package from the standard library - globally available
      - `import "path/to/package"` - import a package from the local directory
      - only Uppercase variables and methods are exported, rest are private
      - All files in current package are available to each other
      ```
      // package 1
      package abc
      func Add(x, y int) int {
        return x + y
      }
      func sub(x, y int) int {
        return x - y
      }
      // package 2
      package main
      import "abc"
      func main() {
        fmt.Println(abc.Add(1, 2))
        fmt.Println(mul(1, 2)) // works
        fmt.Println(abc.sub(1, 2)) // error
      }
      // package 2 but on a diff file
      package main
      import "abc"
      func mul(x, y int) int {
        return x * y
      }
      ```
    - `main` package - entry point
    - `go run <filename>.go` - run the program on a file
    - `go build <filename>.go` - build the program
    - `go build` - compiles and builds the program in the current directory with module name
    - Third party packages 
      - `go get <package-name>` # install the package and add it to the `go.mod` file
      - `go get` # install all the dependencies
      - `go mod tidy` # remove unused dependencies
      - `go list -m all` # list all the dependencies
      - `go list -m -versions <package-name>` # list all the versions of a package
      - `go list -m -versions all` # list all the versions of all the packages
- Functions
    - `func` keyword
    - `main` function - entry point for a running program
    - `func <name> (<args>) <return type> { <body> }`
    - `func main() { <body> }`
  ```go
  // boilerplate code
  func main() {
      fmt.Println("Hello, World!")
  }
  
  func add(x int, y int) int {
      return x + y
  }
  
  func calc(x, y int) (a int, b int) { // multiple return values
    a = x + y
    b = x - y
    return a,b
    // naked return
    // return
  }
  ```
- Modules
    - `go mod init <module-name>` - create a new module
    - `go mod tidy` - remove unused dependencies
    - `go.mod` - module file
- Variables & Operators(~ to CPP)
  - `var <name> <type> = <value>` 
  - `var <name> = <value>` - type inferred
  - `var <name> <type>` - zero value
  - `:=` - shorthand for declaring and initializing a variable
  - Global and local scopes
  - Type casting - `T(v)`
    ```go
    // other boilerplate code
    
    const i int = 42
    var f, u float64 = float64(i), float64(i)
    // f, u := float64(i), float64(i)
    // f = i, u = i # both allowed
    var u uint = uint(f)
    Println(i, f, u)
    var j = "hello" 
    // j := "hello" # allowed
    j = `hello
    world`
        
    // other boilerplate code
    ```
  - Types:: `bool`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`, `byte`, `rune`, `float32`, `float64`, `complex64`, `complex128`
  - Constants:: `const <name> = <value>`
- Control structures
  - if-else
  ```go
  if a == b {
      // code
  } else if a>b {
      // code
  } else { // a < b
      // code
  }
  ```
  - switch
  ```go
  switch a {
      case 1:
          // code
      case 2:
          // code
      default:
          // code
  }
  ```
  - nesting allowed
  - loops: only `for` loops are allowed in multiple forms
  ```go
  for i := 0; i < 10; i++ {
    // code
  }
  for i < 10 {
    // code
  }
  for { // infinite loop
    // code
  }
  // range based loop - list
  for i, v := range list {
    // code
  }

  // range based loop - map
  for k, v := range map {
    // code
  }
  ```
- Connecting to streams
  - `fmt` package - for i/o stream :: `fmt.Println("Hello, World!")`, `fmt.Printf("Hello, %s!\n", "World")`
  - `os` package - for system files :: `os.Args`, `os.Exit(1)`
  ```go
  os.ReadFile("file.txt")
  os.WriteFile("file.txt", data, 0644)
  ```
- Error handling :: handled as if error is any other value
  - `error` type
  - `errors.New("message")` - create an error
  - `if err != nil { // handle error }`
  - `panic` - stop the program
  - `recover` - recover from a panic
  ```go
  if err != nil {
    panic(err)
  }
  ```
- Pointers :: ~ variable handling as in C/C++ :: Pass by value by default
  - `&` - address of operator
  - `*` - dereference operator
  - `new` - create a pointer
  - `nil` - zero value for pointers :: `var p *int` - means p is nil and not assigned
  - `var p *int` - pointer to an int
  - `i := 42; p = &i` - assign address of i to p
  - `fmt.Println(*p)` - read i through the pointer
```go
i := 42
p := &i

fmt.Println(*p) // read i through the pointer
*p = 21 // set i through the pointer
fmt.Println(i) // see the new value of i
```
- Garbage collection
  - Automatic
  - `runtime.GC()` - manual garbage collection 

### $ Structs

- ~ to as in C/C++
- `type <name> struct { <fields> }` # create a struct type
- `var <name> <struct-name>` # create a struct variable
- `var <name> <struct-name> = <struct-name>{<fields>}` # create a struct variable with values
```go
  type person struct {
      name string
      age int
  }
  var p person
  p.name = "Alice"
  p.age = 21
  fmt.Println(p)
  p1 := person{name: "Bob", age: 22}
  fmt.Println(p1)
  p2 := person{"Charlie", 23}
  fmt.Println(p2)
```
- Pointers to structs
```go
p := person{name: "Alice", age: 21}
pp := &p
fmt.Println(pp)
fmt.Println(*pp)
pp.age = 22 // change the value of age through the pointer without dereferencing - allowed
fmt.Println(p)
```
- Methods with struct
  - `func (<name> <struct-name>) <method-name> (<args>) <return-type> { <body> }`
  ```go
  func (p person) greet() { // pass by value
      fmt.Println("Hello, my name is", p.name)
  }
  ```
  - Pass by value to functions have implications on write operations and hence the performance for large structs
  - `func (<name> *<struct-name>) <method-name> (<args>) <return-type> { <body> }` - pass by reference
  ```go
  func (p *person) birthday() { // pass by reference
      p.age++
  }
  ```
  - Constructors - its a pattern to create a new instance of a struct, not a language feature
  ```go
  func newPerson(name string) (*person, error) {
      if name == "" {
          return nil, errors.New("name can't be empty")
      }
      p := person{name: name}
      p.age = 42
      return &p
  }
  ```
  - Public and private fields rules apply if the struct is imported from another package
    - `Name` - public
    - `age` - private
- Embedding
  - ~ to inheritance in OOP
  - Public / private fields rules apply here as well on the embedded struct
  - `type <name> struct { <struct-name> }` # create a struct type
  - `var <name> <struct-name>` # create a struct variable
  - `var <name> <struct-name> = <struct-name>{<fields>}` # create a struct variable with values
  ```go
  type person struct {
      name string
      age int
  }
  type employee struct {
      person
      salary int
  }
  e := employee{person: person{name: "Alice", age: 21}, salary: 50000}
  fmt.Println(e)
  fmt.Println(e.name)
  e.greet()
  ```
- Custom types
  - `type <name> <type>` # create a custom type
  - `type <name> <underlying-type>` # create a custom type
  - `type <name> func(<args>) <return-type>` # create a custom type
  ```go
  type myInt int
  type myFunc func(int, int) int
  type str string
  
  var i myInt = 42
  fmt.Println(i)
  type myFunc func(int, int) int
  f := func(a, b int) int {
      return a + b
  }
  var mf myFunc = f
  fmt.Println(mf(2, 3))
  ```
- Tags: allow to attach metadata to the fields of a struct which can be picked up by other packages
  - `type <name> struct { <fields> }` # create a struct type
  - `var <name> <struct-name>` # create a struct variable
  - `var <name> <struct-name> = <struct-name>{<fields>}` # create a struct variable with values
  ```go
  type person struct {
      name string `json:"person_name"`
      age int `json:"person_age"`
  }
  p := person{name: "Alice", age: 21}
  b, err := json.Marshal(p) // now key will be person_name and person_age
  if err != nil {
      fmt.Println(err)
      return
  }
  fmt.Println(string(b))
  ```
  
### $ Concurrency - The highlight of using Go

`goroutines` - lightweight threads managed by the Go runtime

- `go <function-name>(<args>)` - start a goroutine
- Channels:
  - can be used as return values and errors
  - `make(chan <type>)` - create a channel
  - `ch <- v` - send v to channel ch
  - `v := <-ch` - receive from ch and assign to v - blocks until a value is received
  - `close(ch)` - close the channel
  - `for i := range ch { // code }` - iterate over the channel
  - `select { case i := <-ch: // code case ch <- i: // code default: // code }` - select over multiple channels
  ```go
  ch := make(chan int) // create a channel
  go func() { // start a goroutine
      ch <- 42 // send value to channel
  }()
  v := <-ch // receive value from channel
  fmt.Println(v) // print the value
  
  func slowpoke(phrase string, c chan done) {
      time.Sleep(2 * time.Second)
      fmt.Println(phrase)
      c <- true
  }
  c := make(chan bool)
  go slowpoke("I'm slow", c)
  <-c
  ```
- Ideally use one channel for one goroutine, race-condition can occur otherwise
```go
c := make(chan int)
func greet(val int, c chan int) {
    fmt.Println("Hello, World!")
    c <- c+1
}
go greet(1, c)
go greet(2, c)
go greet(3, c)

fmt.Println(<-c) // first value received will be printed
// So we can use a wait group
cs := make([]chan int, 3)
go greet(1, cs[0])
go greet(2, cs[1])
go greet(3, cs[2])

for _, c := range cs {
    fmt.Println(<-c)
}
// Alternatively, use a single channel
c := make(chan int)
go greet(1, c)
go greet(2, c)
go greet(3, c)

for c := range c {
    fmt.Println(<-c) // wait for three invocations since we call with 3 goroutines
	// greet() can cause deadlock if the channel is not closed
}

// close the channel - for long-running goroutines which will cause other routines listening to the channel to exit as well
func longgreet(val int, c chan int) {
    fmt.Println("Hello, World!")
    c <- c+1
    close(c)
}
```
- Multiple channels - need to be handled with `select`. Else, the program will block if we wait for 2 or more channels on same goroutine
```go
c1 := make(chan int)
erc := make(chan int)
go func(c1 chan int, erc chan int) {
    c1 <- 42
}()
go func(c1 chan int, erc chan int) {
    c1 <- 43
}()
select {
case v := <-c1:
    fmt.Println("Received from c1:", v)
case v := <-c2:
    fmt.Println("Received from c2:", v)
}
```
- `defer`: Execute a function after the surrounding function returns
  - `defer <function-name>(<args>)`
  - `defer func() { // code }()`
  - `defer fmt.Println("Hello, World!")`
  - `defer func() { fmt.Println("Hello, World!") }()`
  - `defer func() { fmt.Println("Hello, World!") }()` - LIFO
  - `defer` is used to close resources, unlock mutexes, etc.
  - `defer` is evaluated at the time of the call, not at the time of the return
  - `defer` is used to handle panics
  ```go
  func main() {
      defer fmt.Println("Hello, World!")
      fmt.Println("Hello, World!")
  }
  ```
- `sync` package
  - `sync.WaitGroup` - wait for a collection of goroutines to finish
  - `sync.Mutex` - mutual exclusion lock
  - `sync.RWMutex` - reader/writer mutual exclusion lock
  - `sync.Once` - run initialization code exactly once
  - `sync.Cond` - condition variable
  ```go
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
      defer wg.Done()
      fmt.Println("Hello, World!")
  }()
  wg.Wait()
  ```
    

Module list::

- `fmt` - formatted I/O
- `math` - mathematical functions
- `os` - operating system functions
- `strings` - string functions
- `time` - time functions
- `errors` - error functions
- `rand` - random number functions
- `reflect` - reflection functions
- `sync` - synchronization functions
- `json` - JSON functions
- `bufio` - buffered I/O