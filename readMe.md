### Go Fundamentals

* Not an Object Oriented Language

#### Go CLI

* go run -> execute file cmd -> go run <filename> <file2> ...
    compiles & executes immediately

* go build -> compiles a bunch of go source code files
    compiles but does not execute

* go fmt -> formats all code in each file in the current dir

* go install -> compiles & installs a package

* go get -> downloads the raw src code of some other pkg

* go test -> runs any test associate with the current project


#### Go Packages

syntax: package <package-name> 

Package (like a project/workspace) =  collection of Go source files
Each file must declare package name at the top (first line)

* Types of packages

1. Executable Packages : Generates an executable file that we can run
2. Reusable Packages: Code used as 'helpers' (reusable logic)

Type is defined by the package name

package named as main -> used to create executable packages
pkg -> go build -> main.exe (on windows)

package named as any other thing -> resusable 
pkg -> go build -> no executable!

!! Any package main must have a function named main() inside it

#### Imports

* syntax: import "pkg-name"
* Used to access some other pkg code like fmt (format library)
* Standard pkgs : debug, math, encoding, fmt, io, crypto
* Files in the same package do not have to be imported into each other.

* Link : https://github.com/StephenGrider/GoCasts/tree/master/diagrams
         https://www.draw.io/


#### Variables

syntax: var <identifier> <type> = <value>
or <identifier> := <value> (only for new variables not reassignment)
* ex

var card string = "value" (type is defined)
is same as
card := 'value' (type is inferred)

* Basic Go Types -> bool, string, int, float64
* cannot use := outside function body but can define var outside
* every variable defined must be used else compilation error

* function with return type

> func newCard() string {
	return "Five of Diamonds"
}

#### Slices & for loops

1. Array (Fixed length list of things/records)
2. Slice (Array that can grow or shrink)

Both hold homogeneous data  

syntax:
> card := []string { "A", "B" }
slice of type strings

Adding new data
>cards = append(cards, "Six of Spades")
* append returns a new slice, does not modify the original data

Iteration
>   for i, card := range cards {
		fmt.Println(i, card)
	}
iteration over closed set -> index is i, item is card
range loops over the slice


#### Functions with receivers [alternative to classes ?]

Attached functionality to our data

* Custome Type
> type deck []string

* Adding function to our custom type (Receiver)
> func (d deck) print() { ... }

here deck is our data type 
d is the actual copy of the deck type variable we are working with







