# Proyecto de Descenso de Gradiente en Go

Este proyecto implementa un algoritmo de descenso de gradiente en Go para minimizar una función cuadrática simple \( f(x) = (x-3)^2 \).

## Funciones Implementadas

1. **objectiveFunction**: Calcula el valor de la función cuadrática \( f(x) = (x-3)^2 \).
2. **derivative**: Calcula la derivada de la función \( f'(x) = 2(x-3) \).
3. **gradientDescent**: Implementa el algoritmo de descenso de gradiente para encontrar el valor de \( x \) que minimiza la función.
4. **main**: Configura los parámetros, ejecuta el descenso de gradiente para varios casos de prueba y reproduce un video sorpresa al final.

## Casos de Prueba

El programa incluye varios casos de prueba con diferentes valores para el punto inicial (\( x_0 \)), el tamaño de paso (\( \alpha \)) y el número máximo de iteraciones:

1. Caso 1: \( x_0 = 0.0 \), \( \alpha = 0.1 \), \( \text{maxIter} = 1000 \)
2. Caso 2: \( x_0 = 5.0 \), \( \alpha = 0.01 \), \( \text{maxIter} = 500 \)
3. Caso 3: \( x_0 = -10.0 \), \( \alpha = 0.05 \), \( \text{maxIter} = 2000 \)
4. Caso 4: \( x_0 = 10.0 \), \( \alpha = 0.001 \), \( \text{maxIter} = 10000 \)
5. Caso 5: \( x_0 = 2.0 \), \( \alpha = 0.5 \), \( \text{maxIter} = 100 \)

## Ejecución

Para ejecutar el programa, simplemente compila y ejecuta el archivo Go:

```bash
go run main.go
```

## Reproducción de Video

Al final de la ejecución, el programa intentará reproducir un video llamado sorpresa.mp4 que debe estar ubicado en la carpeta raíz del script. Si el archivo de video no está presente o hay un error al reproducirlo, se mostrará un mensaje de error.

## Notas
- Asegúrate de tener el archivo `sorpresa.mp4` en la carpeta raíz del script para que se reproduzca correctamente al final.
- Si estás en un sistema operativo diferente a Windows, deberás ajustar el comando en la función playVideo para que funcione en tu entorno.
