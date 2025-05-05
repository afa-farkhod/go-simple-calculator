# 🧮 Simple Calculator (Golang)

A beginner-level CLI calculator built with Go that supports basic arithmetic operations:
- Addition
- Subtraction
- Multiplication
- Division (with zero-division check)

---

## 📂 Project Structure

```
simple-calculator
|-- calculator
|   |-- add.go
|   |-- divide.go
|   |-- multiply.go
|   `-- substract.go
|-- go.mod
`-- main.go
```

## 🚀 Getting Started

1. Clone the repo

```bash
git clone https://github.com/afa-farkhod/simple-calculator.git
cd simple-calculator
```

2. Initialize the module
```
go mod init simple-calculator
```

3. Run the program
```
go run main.go
```

## 💡 Example Usage
```
Enter first number: 12
Enter operator (+, -, *, /): /
Enter second number: 4
✅ Result: 3.00
```

## 📦 Features

- Clear folder structure with separate files per operation
- Proper error handling for division by zero
- Modular and extensible Go code

## 🛠 Technologies

- Go (Golang) 1.22+
- Standard Library only (fmt, errors)

