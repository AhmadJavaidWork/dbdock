# DBDock

DBDock is a modern cross-platform desktop database client built with Go, Wails, and Vue.

It focuses on performance, usability, and a lightweight native desktop experience while supporting large datasets with virtualized rendering and responsive UI interactions.

---

# Features

## Database Management

* Connect to databases from a native desktop application
* Browse tables and schemas
* Open multiple tables in tabs
* Reorder tabs
* Attached and detached table tabs

## Table Viewer

* Virtualized table rendering for large datasets
* Sticky table headers
* Resizable columns
* Pagination support
* Configurable limit and offset
* Horizontal and vertical scrolling

## UI/UX

* Light and dark theme support
* Keyboard shortcuts
* Fast tab navigation
* Responsive layout
* Resizable sidebar

## Developer Experience

* Built with Go + Vue + Wails
* GitHub Actions automated beta releases
* Cross-platform builds
* Modular frontend architecture

---

# Tech Stack

## Backend

* Go
* Wails

## Frontend

* Vue 3
* TypeScript
* Pinia
* TailwindCSS
* Vuelidate

---

# Screenshots


---

# Installation

## Download Releases

Download the latest beta release from the GitHub Releases page.

---

# Development Setup

## Prerequisites

Install:

* Go
* Node.js
* Wails CLI

### Install Wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

---

# Clone Repository

```bash
git clone https://github.com/AhmadJavaidWork/dbdock.git
cd dbdock
```

---

# Install Frontend Dependencies

```bash
cd frontend
npm install
```

---

# Run Development Mode

```bash
wails dev
```

---

# Production Build

## Build Current Platform

```bash
wails build
```

## Build Specific Platforms

### Linux

```bash
wails build -platform linux/amd64
```

### Windows

```bash
wails build -platform windows/amd64
```

### macOS

```bash
wails build -platform darwin/amd64
```

---

# Keyboard Shortcuts

| Shortcut           | Action           |
| ------------------ | ---------------- |
| Ctrl + W           | Close active tab |
| Ctrl + Tab         | Next tab         |
| Ctrl + Shift + Tab | Previous tab     |

---

# Project Structure

```text
.
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── stores/
│   │   ├── services/
│   │   ├── composables/
│   │   └── types/
├── build/
├── app.go
├── main.go
└── wails.json
```

---

# Roadmap

## Planned Features

* Query editor
* SQL autocomplete
* Saved connections
* Query history
* Data editing
* Query execution plans
* Export data
* Multiple database drivers
* Connection groups
* Filters and sorting
* Persistent tabs

---

# Releases

DBDock uses automated GitHub Actions releases.

Current releases are published as beta versions while the project is under active development.

---

# Contributing

Contributions, suggestions, and bug reports are welcome.

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Open a pull request

---

# License

MIT License

---

# Author

Ahmad Javaid

