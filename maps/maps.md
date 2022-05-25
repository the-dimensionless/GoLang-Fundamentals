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