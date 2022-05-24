### Go Fundamentals

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

* Link : https://github.com/StephenGrider/GoCasts/tree/master/diagrams
         https://www.draw.io/








