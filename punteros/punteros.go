package main

import (
	"fmt"
	"strings"
)

func mainPunteros() {
	fmt.Println("🧩 EJEMPLOS DE PUNTEROS EN GO")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("\n📌 Ejemplo básico de punteros:")
	ejemploBasico()

	fmt.Println("\n💾 Gestión de memoria:")
	gestionMemoria()

	fmt.Println("\n🔧 Modificación de variables mediante funciones:")
	ejemploFunciones()

	fmt.Println("\n🧑‍🤝‍🧑 Modificación de structs mediante punteros:")
	ejemploStructs()
}

// ==========================================
// EJEMPLO BÁSICO DE PUNTEROS
// ==========================================
func ejemploBasico() {
	var x int = 10
	var ptr *int = &x

	fmt.Println("Valor de x:", x)        // Output: 10
	fmt.Println("Dirección de x:", ptr)  // Output: dirección de memoria
	fmt.Println("Valor apuntado:", *ptr) // Output: 10

	*ptr = 20
	fmt.Println("Nuevo valor de x:", x) // Output: 20
}

// ==========================================
// GESTIÓN DE MEMORIA CON PUNTEROS
// ==========================================
func gestionMemoria() {
	var ptr *int = new(int)
	fmt.Println("Dirección asignada:", ptr)
	fmt.Println("Valor inicial:", *ptr)

	*ptr = 10
	fmt.Println("Nuevo valor:", *ptr)
}

// ==========================================
// PUNTEROS EN FUNCIONES
// ==========================================
func modifyValue(ptr *int) {
	*ptr = 100
}

func ejemploFunciones() {
	var x int = 42
	fmt.Println("Antes de modificar:", x)
	modifyValue(&x)
	fmt.Println("Después de modificar:", x)
}

// ==========================================
// PUNTEROS Y STRUCTS
// ==========================================
type Person struct {
	Name string
	Age  int
}

func modifyPerson(ptr *Person) {
	ptr.Age = 31
	ptr.Name = "John Doe"
}

func ejemploStructs() {
	p := Person{Name: "John", Age: 30}
	fmt.Println("Antes de modificar:", p)
	modifyPerson(&p)
	fmt.Println("Después de modificar:", p)
}

// ==========================================
// FUNCIÓN PRINCIPAL
// ==========================================
func main() {
	mainPunteros()
}
