# Welcome to Assignment 1: GoLang Basics (for Newbies Who Love Overkill)

## Table of Contents

- [What Even Is This?](#about)
- [How to Survive](#getting_started)
- [How to Use This Beast](#usage)
  - [Task 1 - Library](#task_1)
  - [Task 2 - Shapes](#task_2)
  - [Task 3 - Employees](#task_3)
  - [Task 4 - Bank](#task_4)

## What Even Is This? <a name="about"></a>

Welcome to my glorious interpretation of Assignment 1, where "GoLang Basics" somehow evolved into a Distributed Domain-Driven REST API with microservices. You know, for fun.

## How to Survive <a name="getting_started"></a>

So you want to run this project on your local machine? Brave soul. Here's what you need:

### Prerequisites

Before you dive in, install:

- Go (latest version because old ones are lame)
- A terminal that doesn't crash when you type `go run .`
- Infinite patience

```bash
# Install Go (if you still believe in free will):
https://go.dev/dl/
```

### Installing

1. Clone this repository because originality is overrated:

```bash
git clone https://github.com/ExonegeS/AP1_Assignment1.git 
cd AP1_Assignment1
```

2. Build the app, and may the odds be ever in your favor:

```bash
go build ./cmd/api/
```

If it fails, it's a feature, not a bug.

3. Celebrate with coffee, because Go apps love caffeine-driven developers.

## How to Use This Beast <a name="usage"></a>

1. Launch the server with `./api.exe` and hope nothing explodes.
2. Navigate to `http://localhost:8888/library/book/add` and try not to cry when you get a 404 (or success, if fate is kind).
3. Use Postman or curl because GUI tools are for the weak:

```bash
curl -X POST http://localhost:8888/library/book/add \
    -H "Content-Type: application/json" \
    -d '{"title": "The Go Gospels", "author": "Anonymous"}'
```

4. Rejoice or debug — your destiny awaits!

## Task 1 - Library <a name="task_1"></a>

Build a package to manage a library's book collection using structures, maps, and methods.

Define a struct Book with the following fields:

- ID (string)
- Title (string)
- Author (string)
- IsBorrowed (bool)

Create a Library struct that maintains a collection of books using a map[string]Book where the key is ID.

Implement the following methods for Library:

- `AddBook(book Book)` - Add a new book.
  - Endpoint: `POST http://localhost:8888/library/books/add`
- `ListBooks()` - List all books in the library.
  - Endpoint: `GET http://localhost:8888/library/books/list`
- `GetBook(id string)` - Get a book by ID.
  - Endpoint: `GET http://localhost:8888/library/books/{book_uuid}`
- `PutBook(id string)` - Update existing book.
  - Endpoint: `PUT http://localhost:8888/library/books/{book_uuid}`
- `DeleteBook(id string)` - Delete existing book.
  - Endpoint: `DELETE http://localhost:8888/library/books/{book_uuid}`
- `BorrowBook(id string)` - Borrow a book.
  - Endpoint: `http://localhost:8888/library/books/{book_uuid}/borrow`
- `ReturnBook(id string)` - Return a borrowed book.
  - Endpoint: `PUSThttp://localhost:8888/library/books/{book_uuid}/return`

Use loops to iterate over the books in the map.

## Task 2 - Shapes <a name="task_2"></a>

Create a package to calculate the area and perimeter of different shapes using structs and interfaces.

Define an interface `Shape` with two methods:

- `Area() float64`
- `Perimeter() float64`

Create the following structs:

- `Rectangle` with Length and Width
- `Circle` with Radius
- `Square` with Length
- `Triangle` with SideA, SideB, SideC

Implement the `Shape` interface for all structs.

Use a slice of `Shape` to store multiple shapes and iterate using a loop to print details.

- `PrintShapeDetails(s Shape)`
  - Endpoint: `http://localhost:8888/shapes/print`
  
## Task 3 - Employees <a name="task_3"></a>

Implement an employee management system using interfaces, maps, and methods.

Define an interface `Employee` with a method:

- `GetDetails() string`

Create structs:

- `FullTimeEmployee` with ID, Name, and Salary
- `PartTimeEmployee` with ID, Name, HourlyRate, and HoursWorked

Implement the `Employee` interface for both structs.

Implement methods:
- `AddEmployee(employee Employee)` - Add a new employee.
  - Endpoint: `POST http://localhost:8888/employees/add`
- `ListEmployees()` - List all employees in the company.
  - Endpoint: `GET http://localhost:8888/employees/list`
- `GetEmployeeDetails(id string)` - Get a employee details in readable format.
  - Endpoint: `GET http://localhost:8888/employees/{employee_id}/details`
- `GetEmployee(id string)` - Get a employee by ID.
  - Endpoint: `GET http://localhost:8888/employees/{employee_id}`
- `PutEmployee(id string)` - Update existing book.
  - Endpoint: `PUT http://localhost:8888/employees/{employee_id}`
- `DeleteEmployee(id string)` - Delete existing employee.
  - Endpoint: `DELETE http://localhost:8888/employees/{employee_id}`


Use a loop to display employee details.

## Task 4 - Bank <a name="task_4"></a>

Simulate a bank account system using structs, methods, and interfaces.

Define a struct `BankAccount` with:

- `AccountNumber string`
- `HolderName string`
- `Balance float64`

Implement methods:

- `Deposit(amount float64)` - Add money.
- `Withdraw(amount float64)` - Deduct money if sufficient balance exists.
- `GetBalance()` - Print the current balance.

Write a function `Transaction(account *BankAccount, transactions []float64)` that processes a slice of transaction amounts:

- Positive amounts are deposits.
- Negative amounts are withdrawals.

Implement methods:
- `AddWallet(wallet Wallet)` - Add a new wallet.
  - Endpoint: `POST http://localhost:8888/bank/add`
- `GetBalance()` - Retrieve a wallet's current contents.
  - Endpoint: `GET http://localhost:8888/bank/{wallet_id}/balance`
- `CreateTransation(transation Transation)` - Create a request for later processing.
  - Endpoint: `POST http://localhost:8888/bank/{wallet_id}/add`
- `ProcessTransactions(walletID uuid)` - Process all transactions for the wallet.
  - Endpoint: `POST http://localhost:8888/bank/{wallet_id}/commit`
- `GetTransactions(walletID uuid)` - Retrieve all remaining transactions.
  - Endpoint: `GET http://localhost:8888/bank/{wallet_id}`
