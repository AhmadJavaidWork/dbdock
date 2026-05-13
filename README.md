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

<img width="501.25313275" height="250" alt="1" src="https://github.com/user-attachments/assets/ac6dfb62-9d1d-4313-9588-9ba32e7dc986" />
<img width="501.25313275" height="250" alt="2" src="https://github.com/user-attachments/assets/3b0ff834-0d12-4692-8c2f-719e406b929f" />
<img width="501.25313275" height="250" alt="3" src="https://github.com/user-attachments/assets/e122d399-bf54-4504-97c5-830d58065bea" />
<img width="501.25313275" height="250" alt="4" src="https://github.com/user-attachments/assets/7821f3ac-e760-4c70-98ea-b4ed1b1fcbc1" />
<img width="501.25313275" height="250" alt="5" src="https://github.com/user-attachments/assets/98d0b608-4d0c-4cd2-932e-62033b823725" />
<img width="501.25313275" height="250" alt="6" src="https://github.com/user-attachments/assets/097c5b60-8975-43a5-bda8-d5d2272f6394" />
<img width="501.25313275" height="250" alt="7" src="https://github.com/user-attachments/assets/4bc624a2-a9f3-4577-8085-43e6d49a9e47" />
<img width="501.25313275" height="250" alt="8" src="https://github.com/user-attachments/assets/006b44da-5687-4f9e-8d65-31fbb10e9101" />
<img width="501.25313275" height="250" alt="9" src="https://github.com/user-attachments/assets/7e973834-0cc7-4b12-8c29-2cbe94862a7e" />
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

