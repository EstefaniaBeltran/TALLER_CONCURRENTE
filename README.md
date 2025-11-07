# TALLER_CONCURRENTE
Integrantes: Estefanía Beltrán, Marc Suarez y Alejandro Poveda
# Ejercicio 1 – Ejecución concurrente con Goroutines

Descripción

Crea un programa que ejecute tres tareas concurrentes:
1. Imprimir los números del 1 al 5.
2. Imprimir las letras de la A a la E.
3. Mostrar un mensaje de estado cada segundo (“Procesando…”).
   
Usa sync.WaitGroup para que el programa principal espere a que todas las goroutines finalicen antes de terminar.
Conceptos:
Lanzamiento de goroutines con go .
Sincronización con sync.WaitGroup .

# Ejercicio 2 – Comunicación concurrente con canales

Descripción

Implementa un programa que cree dos goroutines:
1. Una productora, que genere los números del 1 al 10 y los envíe a un canal.
2. Una consumidora, que reciba esos números, filtre los pares y los muestre en pantalla.
3. Cierra el canal correctamente cuando no haya más datos.
   
Conceptos:
Comunicación entre procesos concurrentes con channels.
Recepción, filtrado y cierre de canales.

# Ejercicio 3 – Scheduling concurrente con prioridad simulada

Descripción

Implementa un programa que simule un planificador concurrente (scheduler) que gestiona tareas con diferentes niveles de prioridad.
Cada tarea será una goroutine que realiza un trabajo (por ejemplo, una operación matemática o una simulación de espera) y reporta su finalización.

Instrucciones

1. Define una estructura Tarea con los campos:
id int
prioridad int
duracion time.Duration
2. Crea una lista de al menos 5 tareas con distintas prioridades (1 = alta, 3 = baja) y duraciones aleatorias (usa time.Sleep() para simular tiempo de
ejecución).
3. Implementa una función ejecutarTarea(t Tarea, resultados chan<- string) que:
Simule la ejecución de la tarea (durmiendo según su duración).
Envíe un mensaje al canal de resultados indicando que la tarea terminó.
4. En la función main() :
Lanza una goroutine planificadora que envíe las tareas al canal de ejecución empezando por las de mayor prioridad .
Lanza varias goroutines “worker” que tomen tareas del canal y las ejecuten concurrentemente.
Muestra los resultados a medida que llegan a través del canal de resultados.
Usa sync.WaitGroup para esperar que todas las tareas finalicen antes de imprimir el mensaje final:

Todas las tareas han sido completadas.

Conceptos:

Planificación de tareas concurrentes con prioridad simulada.
Comunicación y coordinación entre goroutines.
Sincronización mediante sync.WaitGroup .


