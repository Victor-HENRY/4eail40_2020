package main

import "fmt"

//mon code
type Shape interface {
    String() string
}

type Rectangle struct {
    Longueur  int64
    Largeur int64
}
func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle of width %d and length %d", r.Longueur, r.Largeur)
}

type Circle struct {
    Rayon int64
}
func (c Circle) String() string {
    return fmt.Sprintf("Circle of radius %d", c.Rayon)
}


func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
