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