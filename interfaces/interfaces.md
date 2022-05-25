### Interfaces

Q) If same logic goes for different types, should there
    be multiple different functions for the same logic ?

Ans) NO. Normally we could have Generics or interfaces

* func (eb englishBot) <=> func (englishBot) if eb is not used

* method overloading not done [PROBLEM !!!] Then how to reuse ?


* type bot interface {
	getGreeting() string
}

* children can have their own implementations

* Rules of interfaces

Concrete type => (can create value directly) map, struct, int, string
Interface type => cannot create value directly

Interfaces are not generic types

* Go does not have generic types

Interfaces are implicitly understood, we don't need to explcitly define relation between interface & implementing class

They are a contract to help us manage types.
Garbage in => Garbage out!