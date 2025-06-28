**1. go.mod y go.sum**

go.mod define el nombre del módulo y gestiona las dependencias. El archivo go.mod es la raíz
de la gestión de dependencias en GoLang. Todos los módulos que se necesitan o se van a
utilizar en el proyecto se mantienen en el archivo go.mod. Para todos los paquetes que vayamos
a importar/utilizar en nuestro proyecto, creará una entrada de esos módulos en go.mod. Tener un
archivo go mod ahorra el esfuerzo de ejecutar el comando go get para cada módulo
dependiente para ejecutar el proyecto con éxito.
