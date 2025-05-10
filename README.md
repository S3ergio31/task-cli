# 🧩 task-cli

Un **task tracker minimalista para la terminal**, escrito en Go. Inspirado en [roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker), este proyecto me sirvió para aprender y practicar el lenguaje Go desde cero.

## 🚀 ¿Qué es task-cli?

`task-cli` es una pequeña herramienta de línea de comandos para gestionar tareas pendientes. Permite agregar, modificar la descripción, listar, cambiar el estado y eliminar tareas desde la terminal, sin necesidad de base de datos ni dependencias externas.

## 📦 Versiones

### ✅ v1.0.0
> Primera versión funcional.
- Todo el código en un solo archivo.
- Familiarización con la sintaxis y estructura básica de Go.

### 📁 v2.0.0
> Organización y calidad.
- Código modularizado en paquetes.
- Tests unitarios con **100% de cobertura**.

### 🧠 v3.0.0
> Mejora arquitectónica.
- Introducción de una interfaz genérica `Command` para representar cada comando.
- Código más limpio, mantenible y extensible.

## 🛠️ Instalación

```bash
git clone https://github.com/tuusuario/task-cli.git
cd task-cli/src/task-cli
go build -o task
```

## ⚙️ Uso

```bash
./task add "Aprender Go"
./task update <uuid> "aprender go"
./task list
./task mark <uuid> done
./task remove <uuid>
```

## 🧾 Comandos disponibles

- `add <description>` – Agrega una nueva tarea.
- `update <uuid> <description>` – Actualiza la descripción de una tarea.
- `list <status>` – Lista todas las tareas por estado: todo, in-progress, done. Si no se especifica, se muestran todas.
- `mark <uuid> <status>` – Marca una tarea como todo, in-progress o done.
- `remove <uuid>` – Elimina una tarea.

## 🧪 Tests

```bash
go test .
```

## 🌱 Aprendizajes

Este proyecto me permitió:
- Entender cómo funciona Go desde sus bases.
- Aplicar buenas prácticas de testing y organización del código.
- Diseñar una arquitectura extensible usando interfaces.

## 📚 Tecnologías

- [Go](https://golang.org/)
- Terminal / CLI

## 📄 Licencia

MIT License © [Sergio Fidelis](https://github.com/S3ergio31)
