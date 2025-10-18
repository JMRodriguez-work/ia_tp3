# TP3 - Inteligencia Artificial

**Prototipo con modelo de Hopfield para identificación de imágenes**

El modelo de Hopfield es una red neuronal recurrente que funciona como un sistema de memoria asociativa, capaz de recordar patrones a partir de versiones incompletas o distorsionadas. En este prototipo se aplica al reconocimiento de un aro “C” representado en una matriz 10x10, donde cada celda indica si un píxel pertenece o no a la pieza.

Se trabajó con un conjunto reducido de 4 patrones de entrenamiento, incluyendo variaciones del aro (rotado, parcialmente cubierto, con ruido, etc.). Para el aprendizaje se utilizó la regla de Hebb, que actualiza los pesos sinápticos mediante el producto de los valores de los patrones, sin intervención de un maestro.

## Modos de ejecución

Se puede ejecutar el binario o correr el proyecto localmente con GO

### Para el binario dependiendo del SO

En la sección **[Releases](./releases/)** de este repositorio se encuentran los archivos, descargar el correspondiente y hacer doble click en el mismo.

### Requisitos locales

- Tener instalado [Go](https://go.dev/dl/)

#### Pasos

1. Clonar este repositorio o descargar los archivos:
   ```bash
   git clone https://github.com/JMRodriguez-work/ia_tp3
   cd ia_tp3
   ```
2. Ejecutar el programa
   ```go
   go run main.go
   ```