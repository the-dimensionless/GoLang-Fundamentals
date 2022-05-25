### Maps

* both keys & values are statically typed


* method 1 => 
    colors:= map[string]string{
		"red":   "#ff0000",
		"black": "#ffffff",
		"white": "#000000",
	}

* method 2 defining => 
    var c map[string]string

* method 3 defining => 
    cols := make(map[string]string)


* manipulation
    c[key] = value => works
    c.key = value => won't work as keys & values are statically typed

* iteration over map
    for key, value := range c {
		fmt.Printf("Key: %v and Value: %v", key, value)
	}

* Map vs Struct

1. All keys & values must be of the same type vs
        values can be of different types (key is not under question)

2. All keys are indexed (can be iterated over) vs
        keys don't support indexing

3. Reference Type vs Value Type

4. Closely related props vs not needed to be related

5. Dynamic keys (can change over time) vs static keys (pre defined)