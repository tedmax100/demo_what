package main

type hero struct {
	name       string
	superpower string
}

//go:noinline
func createSuperMan() hero {
	h := hero{
		name:       "Superman",
		superpower: "X-ray vision",
	}
	println("Superman ", &h)
	return h
}

//go:noinline
func createTheFlash() *hero {
	h := hero{
		name:       "The Flash",
		superpower: "Super speed",
	}
	println("The Flash ", &h)
	return &h
}

func main() {
	h1 := createSuperMan()

	// h will escape to heap, after createTheFlash returns its value, no matter if h2 is declared or not
	h2 := createTheFlash()

	println("Superman ", &h1, "The Flash", &h2)
}
