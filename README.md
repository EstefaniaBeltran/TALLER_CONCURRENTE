# TALLER_CONCURRENTE
#Ejercicio 1 – Ejecución concurrente con Goroutines

Descripción

Crea un programa que ejecute tres tareas concurrentes:
1. Imprimir los números del 1 al 5.
2. Imprimir las letras de la A a la E.
3. Mostrar un mensaje de estado cada segundo (“Procesando…”).
Usa sync.WaitGroup para que el programa principal espere a que todas las goroutines finalicen antes de terminar.
Conceptos:
Lanzamiento de goroutines con go .
Sincronización con sync.WaitGroup .

#Ejercicio 2 – Comunicación concurrente con canales

Descripción

Implementa un programa que cree dos goroutines:
1. Una productora, que genere los números del 1 al 10 y los envíe a un canal.
2. Una consumidora, que reciba esos números, filtre los pares y los muestre en pantalla.
3. Cierra el canal correctamente cuando no haya más datos.
Conceptos:
Comunicación entre procesos concurrentes con channels.
Recepción, filtrado y cierre de canales.

