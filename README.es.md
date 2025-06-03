# 📊 github-user-activity

Una herramienta de línea de comandos escrita en Go que muestra la actividad reciente de un usuario de GitHub directamente en la terminal. Inspirada en el proyecto de roadmap.sh: [GitHub User Activity](https://roadmap.sh/projects/github-user-activity).

## 🚀 ¿Qué hace?

Este programa consume la API pública de GitHub para obtener eventos recientes de un usuario y los muestra en la terminal con mensajes personalizados según el tipo de evento.

## 🛠️ Instalación

```bash
git clone https://github.com/S3ergio31/github-user-activity.git
cd github-user-activity
go build -o github-activity
```

## ⚙️ Uso

```bash
./github-activity <nombre-de-usuario>
```

Ejemplo:

```bash
./github-activity S3ergio31
```

## 📋 Ejemplo de salida

```
- Pushed 1 commits to S3ergio31/github-user-activity
- A new 'branch' was created in S3ergio31/github-user-activity
- A new 'repository' was created in S3ergio31/url-shortener
- A new 'tag' was created in S3ergio31/url-shortener
```

## 🧪 Tests

```bash
go test ./...
```

## 🌱 Aprendizajes

Este proyecto me permitió:

- Practicar el consumo de APIs REST en Go.
- Manejar y parsear datos JSON.
- Aplicar principios de diseño limpio y modular.

## 📚 Tecnologías

- [Go](https://golang.org/)
- API de GitHub

## 📄 Licencia

MIT License © [S3ergio31](https://github.com/S3ergio31)
