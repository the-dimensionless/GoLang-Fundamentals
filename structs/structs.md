### Structs

* Zero values

1. string => ""
2. int => 0
3. float => 0
4. bool => false


* Structs can be embedded inside one another like
Person can have ContactInfo

* , after multiline structs

### Memory Model
* Go is pass by value: therefore when

Creating a struct => jim := person {}
Stores it in some place in RAM (say OX01)

Updating a struct => using functions with receivers
> func (p person) updateName() {}
p is not the original jim, it's a copy of jim 
changes to p will not be reflected in jim

> func (pointerToP *person) updateName() {}
works as we are passing a pointer to the original jim struct

* Pointers 

var a := 10
p := &a

& -> operator to get the mem. address of variable
printin p gives some mem location

*pointer -> operator to get the value pointed to by this pointer


*person => *<type> => pointer to this (person) type of variable
*pointerVariable => value at that mem pointed by this pointer

* reference types like slices, maps, channels, pointers, functions
are managed internally in terms of value -> pointer and vice versa
conversion
* basic types => int, float, string, bool, struct => we need to take
care of pointers while passing into functions


* Creating a slice
Go automatically creates an array and a structure that recors
the length of slice,
the capacity of the slice,
and a reference to the underlying array