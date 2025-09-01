# Simplex Algorithm Solver

Un solucionador de problemas de programación lineal usando el método Simplex.

## Estado del Proyecto

Este proyecto implementa la **integración con la entrada de datos** (Issue #5) para el algoritmo Simplex. Permite a los usuarios ingresar problemas de programación lineal de manera interactiva con validación completa.

## Características Implementadas

### ✅ Entrada de Datos (Issue #5)
- Entrada interactiva de función objetivo (maximización/minimización)
- Entrada de restricciones con operadores (<=, >=, =)
- Validación de entrada numérica
- Visualización clara de ecuaciones
- Restricciones de no negatividad automáticas

### ✅ Validación de Datos
- Verificación de entrada numérica válida
- Validación de consistencia entre variables
- Mensajes de error informativos en español

### ✅ Visualización de Problemas
- Formato claro de ecuaciones matemáticas
- Resumen del problema ingresado
- Problema de ejemplo del Issue #1

## Instalación y Uso

### Prerrequisitos
- Go 1.16 o superior

### Compilación
```bash
go build -o simplex .
```

### Modos de Uso

#### 1. Modo Interactivo
```bash
./simplex
```
Permite ingresar un problema paso a paso de manera interactiva.

#### 2. Modo Demo
```bash
./simplex --example
```
Muestra el problema de ejemplo del Issue #1:
- Maximizar: 5x + 3y
- Sujeto a: 2x + y ≤ 20, x + y ≤ 12
- Solución esperada: x=8, y=4, valor=52

#### 3. Ayuda
```bash
./simplex --help
```

## Ejemplo de Uso Interactivo

```
Simplex Algorithm Solver
========================

Ingreso de datos para el problema de programación lineal
=========================================================

1. Función Objetivo
-------------------
¿Desea maximizar o minimizar? (max/min): max
Ingrese los nombres de las variables separados por espacios: x y
Ingrese coeficiente de x: 5
Ingrese coeficiente de y: 3

2. Restricciones
----------------
Restricción 1:
Ingrese coeficiente de x: 2
Ingrese coeficiente de y: 1
Ingrese el operador (<=, >=, =): <=
Ingrese valor del lado derecho: 20
¿Desea agregar otra restricción? (s/n): s

Restricción 2:
Ingrese coeficiente de x: 1
Ingrese coeficiente de y: 1
Ingrese el operador (<=, >=, =): <=
Ingrese valor del lado derecho: 12
¿Desea agregar otra restricción? (s/n): n
```

## Estructura del Proyecto

```
simplex/
├── main.go                    # Programa principal
├── pkg/simplex/
│   ├── problem.go            # Estructuras de datos y entrada
│   ├── problem_test.go       # Tests para entrada de datos
│   ├── validation.go         # Validación y utilidades
│   └── validation_test.go    # Tests para validación
├── go.mod                    # Módulo Go
└── README.md                # Esta documentación
```

## Testing

Ejecutar todas las pruebas:
```bash
go test ./pkg/simplex/ -v
```

## Roadmap

### 🚧 Próximas Características (Issues Relacionados)

- **Issue #7**: Diseño del algoritmo Simplex
- **Issue #3**: Visualización del proceso iterativo (tableaux)
- **Issue #4**: Guardar y cargar problemas
- **Issue #2**: Mejoras en carga de datos y visualización

### 🎯 Funcionalidades Futuras
- Resolución completa del algoritmo Simplex
- Exportación/importación de problemas en JSON
- Interfaz web opcional
- Visualización gráfica de soluciones

## Contribución

Este proyecto está en desarrollo activo. Las contribuciones son bienvenidas siguiendo los issues definidos en el repositorio.

## Licencia

MIT License - Ver archivo LICENSE para más detalles.