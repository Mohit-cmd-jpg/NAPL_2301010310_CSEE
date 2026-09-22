# NAPL Go Lab Assignments

This repository contains Go programming lab assignments for NAPL.

## Repository Structure

- `Lab 1/` - introductory Go programs (`Q1.go` and `Q2.go`)
- `Lab 2/exp1/` - mathematical utility functions and user input
- `Lab 2/exp2/` - collection management using slices and maps
- `Lab 3/exp1/` - `Person` struct with associated methods

## Requirements

- Go 1.23 or newer
- PowerShell, Command Prompt, or a terminal

## Running the Programs

Run Lab 1 programs from the repository root:

```powershell
go run ".\Lab 1\Q1.go"
go run ".\Lab 1\Q2.go"
```

Run Lab 2 Experiment 1:

```powershell
Push-Location ".\Lab 2\exp1"
go run .
Pop-Location
```

Run Lab 2 Experiment 2:

```powershell
Push-Location ".\Lab 2\exp2"
go run .
Pop-Location
```

Run Lab 3 Experiment 1:

```powershell
go run ".\Lab 3\exp1\main.go"
```

Lab 2 experiments use their own `go.mod` files. Lab 3 Experiment 1 is a standalone Go source file and does not require a module file.
