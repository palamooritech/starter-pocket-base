# PocketBase

PocketBase is an open-source backend framework built in Go. It offers features such as an embedded SQLite database, real-time subscriptions, built-in authentication management, a user-friendly dashboard UI, and a simple REST-ish API. It's designed to provide a secure and cost-effective foundation for web applications.

# Intro
This is a starter template, to extend Pocketbase as a GO framework. Just clone it and extend it as per your needs.

## Getting Started

To use PocketBase, you can clone this repository and add your custom paths. Here's a brief overview of the directory structure:


- **main.go**: Your project's main application file.

- **models/**: This directory contains all the model files, likely representing your data structures.

- **helpers/**: This directory contains helper functions and middleware files used in your application.

- **pkg/**: The `pkg` directory contains the main codebase of your application, organized into subdirectories.

  - **adapters/**: This directory may contain various adapters, and it includes a subdirectory `pocketbase/` for PocketBase-related code.

  - **domain/**: This directory houses the core and main components of your application.

project_root/
│
├── main.go
│
├── models/
│   ├── (model files)
│
├── helpers/
│   ├── (helper and middleware files)
│
├── pkg/
│   ├── adapters/
│   │   ├── pocketbase/
│   │   │   ├── (PocketBase-related files)
│   │
│   ├── domain/
│   │   ├── (core and main component files)
│
├── README.md
│
├── other_project_files...


## Usage

1. Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/pocketbase.git
```

2. Start the application
```bash
make serve
```