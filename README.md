# Loan Engine

## Project Description
The Loan Engine is a comprehensive application designed to manage the lifecycle of loans. It facilitates the movement of loans through various states—**Proposed**, **Approved**, **Invested**, and **Disbursed**. Each state has specific rules and requirements that must be adhered to, ensuring a structured and reliable lending process.
   
## API Design

This project includes a RESTful API that supports the following operations:

- **Get Available Loan**
- **Get Loan Detail**
- **Create a Loan**
- **Approve a Loan**
- **Invest in a Loan**
- **Disburse a Loan**

## Getting Started

### Prerequisites

- Go 1.23
- MySQL 8.3.0

### Installation

1. Execute migration file
   ```bash
   cd migration
   ```
   execute DDL in database
2. Clone the repository:
   ```bash
   git clone https://github.com/martinus490/loan-engine.git
   cd loan-engine
2. Run the program
   ```bash
   go run main.go
3. Run the unit test
   ```bash
   cd tests
   go test ./...
