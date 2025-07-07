# E-COMMERCE-BASIC-GO-ECHO

_Simplify Commerce, Empower Growth, Accelerate Success_

[![Last Commit](https://img.shields.io/github/last-commit/nawwardiab/e-commerce-basic-go-echo)](https://github.com/nawwardiab/e-commerce-basic-go-echo) [![Go](https://img.shields.io/badge/Go-82.1%25-blue)](https://golang.org/) [![Languages](https://img.shields.io/github/languages/count/nawwardiab/e-commerce-basic-go-echo)](https://github.com/nawwardiab/e-commerce-basic-go-echo)
Built with the tools and technologies:
![Markdown](https://img.shields.io/badge/Markdown-000000?logo=markdown) ![Go Modules](https://img.shields.io/badge/Go%20Modules-000000?logo=go) ![YAML](https://img.shields.io/badge/YAML-000000?logo=yaml) ![Echo](https://img.shields.io/badge/Echo-Framework-blue)

---

## Table of Contents

- [E-COMMERCE-BASIC-GO-ECHO](#e-commerce-basic-go-echo)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Testing](#testing)

---

## Overview

_e-commerce-basic-go-echo_ is a lightweight, modular template for building a complete e-commerce platform using the Echo web framework. It provides a clean, organized architecture that simplifies development of secure user workflows, product management, and shopping-cart functionalities with Echo's powerful routing and middleware capabilities.
**Why \*\*\***e-commerce-basic-go-echo**\***?\*\*  
This project helps developers rapidly prototype and build scalable online stores with core features like user registration, product browsing, and cart management. Built on Echo's high-performance HTTP router, it offers enhanced middleware support and streamlined request handling. The core features include:

- 🔐 **User Authentication:** Secure registration, login, and session handling with Echo's context-aware middleware.
- 🎨 **Dynamic Templates:** Rendered views for product listings, details, and shopping cart with Echo's template engine integration.
- 🚀 **Modular Architecture:** Clear separation of concerns with dedicated handlers, services, repositories, and models optimized for Echo's structure.
- ⚙️ **Configurable Setup:** Environment flexibility via YAML configuration for deployment across different environments.
- 🌐 **Static Asset Middleware:** Efficient static resource serving using Echo's built-in static middleware.
- 🗄️ **Database Integration:** Reliable PostgreSQL connection management supporting persistent data storage.
- ⚡ **High Performance:** Built on Echo's optimized HTTP router for superior performance and minimal memory footprint.
- 🛡️ **Middleware Support:** Comprehensive middleware stack including CORS, logging, recovery, and custom authentication.

---

## Getting Started

### Prerequisites

This project requires the following dependencies:

- **Programming Language:** [Go](https://golang.org/)
- **Package Manager:** Go modules
- **Web Framework:** [Echo](https://echo.labstack.com/)

### Installation

Build _e-commerce-basic-go-echo_ from source and install dependencies:

1. **Clone the repository:**

```bash
git clone https://github.com/nawwardiab/e-commerce-basic-go-echo
```

2. **Navigate to the project directory:**

```bash
cd e-commerce-basic-go-echo
```

3. **Install the dependencies (using Go modules):**

```bash
go mod download
```

### Usage

Run the project with:

```bash
go run main.go
```

Or build and then run:

```bash
go build ./e-commerce-basic-go-echo
```

### Testing

_e-commerce-basic-go-echo_ uses the built-in Go testing framework. Run the full test suite with:

```bash
go test ./...
```
