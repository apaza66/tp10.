package main

import (
"encoding/csv"
"fmt"
"os"
"strconv"
)

func main() {
// Abrir el archivo CSV para escritura
file, err := os.Create("productos.csv")
if err != nil {
fmt.Println("Error al crear el archivo:", err)
return
}
defer file.Close()

// Crear el escritor CSV
writer := csv.NewWriter(file)
defer writer.Flush()

fmt.Println("Ingrese los detalles de los productos:")

// Escribir encabezados en el CSV
err = writer.Write([]string{"Producto", "Precio", "Codigo"})
if err != nil {
fmt.Println("Error al escribir encabezados:", err)
return
}

// Variables para calcular la sumatoria de precios
var sumatoriaPrecios float64

// Leer y escribir detalles de productos
for i := 0; i < 10; i++ {
var producto, codigo string
var precio float64

fmt.Printf("\nProducto %d:\n", i+1)

fmt.Print("Nombre del producto: ")
fmt.Scan(&producto)

fmt.Print("Precio del producto: ")
fmt.Scan(&precio)

fmt.Print("Codigo del producto: ")
fmt.Scan(&codigo)

// Convertir el precio a string
precioStr := strconv.FormatFloat(precio, 'f', 2, 64)

// Escribir los detalles en el CSV
err := writer.Write([]string{producto, precioStr, codigo})
if err != nil {
fmt.Println("Error al escribir en el archivo:", err)
return
}

// Agregar el precio a la sumatoria
sumatoriaPrecios += precio
}

fmt.Printf("Los detalles de los productos se han guardado en productos.csv.\n")
fmt.Printf("La sumatoria de los precios es: %.2f\n", sumatoriaPrecios)
}
package main

import (
"encoding/csv"
"fmt"
"os"
"strconv"
)

func main() {
// Abrir el archivo CSV para escritura
file, err := os.Create("productos.csv")
if err != nil {
fmt.Println("Error al crear el archivo:", err)
return
}
defer file.Close()

// Crear el escritor CSV
writer := csv.NewWriter(file)
defer writer.Flush()

fmt.Println("Ingrese los detalles de los productos:")

// Escribir encabezados en el CSV
err = writer.Write([]string{"Producto", "Precio", "Codigo"})
if err != nil {
fmt.Println("Error al escribir encabezados:", err)
return
}

// Variables para calcular la sumatoria de precios
var sumatoriaPrecios float64

// Leer y escribir detalles de productos
for i := 0; i < 10; i++ {
var producto, codigo string
var precio float64

fmt.Printf("\nProducto %d:\n", i+1)

fmt.Print("Nombre del producto: ")
fmt.Scan(&producto)

fmt.Print("Precio del producto: ")
fmt.Scan(&precio)

fmt.Print("Codigo del producto: ")
fmt.Scan(&codigo)

// Convertir el precio a string
precioStr := strconv.FormatFloat(precio, 'f', 2, 64)

// Escribir los detalles en el CSV
err := writer.Write([]string{producto, precioStr, codigo})
if err != nil {
fmt.Println("Error al escribir en el archivo:", err)
return
}

// Agregar el precio a la sumatoria
sumatoriaPrecios += precio
}

fmt.Printf("Los detalles de los productos se han guardado en productos.csv.\n")
fmt.Printf("La sumatoria de los precios es: %.2f\n", sumatoriaPrecios)
}
package main

import (
"encoding/csv"
"fmt"
"os"
"strconv"
)

func main() {
// Abrir el archivo CSV para escritura
file, err := os.Create("productos.csv")
if err != nil {
fmt.Println("Error al crear el archivo:", err)
return
}
defer file.Close()

// Crear el escritor CSV
writer := csv.NewWriter(file)
defer writer.Flush()

fmt.Println("Ingrese los detalles de los productos:")

// Escribir encabezados en el CSV
err = writer.Write([]string{"Producto", "Precio", "Codigo"})
if err != nil {
fmt.Println("Error al escribir encabezados:", err)
return
}

// Variables para calcular la sumatoria de precios
var sumatoriaPrecios float64

// Leer y escribir detalles de productos
for i := 0; i < 10; i++ {
var producto, codigo string
var precio float64

fmt.Printf("\nProducto %d:\n", i+1)

fmt.Print("Nombre del producto: ")
fmt.Scan(&producto)

fmt.Print("Precio del producto: ")
fmt.Scan(&precio)

fmt.Print("Codigo del producto: ")
fmt.Scan(&codigo)

// Convertir el precio a string
precioStr := strconv.FormatFloat(precio, 'f', 2, 64)

// Escribir los detalles en el CSV
err := writer.Write([]string{producto, precioStr, codigo})
if err != nil {
fmt.Println("Error al escribir en el archivo:", err)
return
}

// Agregar el precio a la sumatoria
sumatoriaPrecios += precio
}

fmt.Printf("Los detalles de los productos se han guardado en productos.csv.\n")
fmt.Printf("La sumatoria de los precios es: %.2f\n", sumatoriaPrecios)
}
