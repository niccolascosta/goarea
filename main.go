package goarea

import "math"

// Pi é proporcao numerica definida pela relacao entre
// o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ é responsavel por calcular a area da circuferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// React é responsavel por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao e visivel!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
