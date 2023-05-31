package main

import "fmt"

/*type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}*/

type FiguraGeometrica interface {
	area() float64
}

type Rectangulo struct {
	base, altura float64
}

type Trapecio struct {
	baseMayor, baseMenor, altura float64
}

func (figura Rectangulo) area() float64 {
	return figura.base * figura.altura
}

func (figura Trapecio) area() float64 {
	return (figura.baseMayor + figura.baseMenor) * figura.altura / 2
}

func dameArea(fig FiguraGeometrica) float64 {
	return fig.area()
}

func main() {
	rectangulo := Rectangulo{base: 4, altura: 7.5}
	trapecio := Trapecio{baseMayor: 5, baseMenor: 2, altura: 3}

	fmt.Printf("area del rectangulo %f\n", dameArea(rectangulo))

	fmt.Printf("area del trapecio %f\n", trapecio.area())
	fmt.Printf("area del trapecio %f\n", dameArea(trapecio))
}
