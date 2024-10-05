# Budget Tracker App 📊

A simple **Personal Budget Tracker** application written in **Go** that helps you manage your **income** and **expenses**, categorize transactions, and save them in a **CSV file** for future reference.

## Features
- Add and categorize transactions as **income** or **expense**
- Display a list of all transactions with **amount**, **category**, **date**, and **type**
- Calculate **total income** and **total expenses**
- Save transactions to a **CSV file** for easy management

## Technologies Used
- **Go** (Golang)
- **CSV** for file storage
- **fmt**, **os**, **time**, and **encoding/csv** Go standard libraries

## Getting Started

### Prerequisites
- **Go** installed on your system (version 1.16 or higher)
- Basic knowledge of Go programming

### Installation
1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/budget-tracker-app.git
   cd budget-tracker-app
   go run main.go```

Also, here is the file structure:
.

├── main.go           # Go source code for the budget tracker app

├── transactions.csv  # CSV file where your transactions will be saved
