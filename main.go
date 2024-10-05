package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Transaction struct to hold each transaction's details
type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string // "income" or "expense"
}

// BudgetTracker struct to manage transactions
type BudgetTracker struct {
	transactions []Transaction
	nextID       int
}

// Interface for common behavior
type FinancialRecord interface {
	GetAmount() float64
	GetType() string
}

// Implementing interface methods for Transaction
func (t Transaction) GetAmount() float64 {
	return t.Amount
}

func (t Transaction) GetType() string {
	return t.Type
}

// Add a new transaction
func (bt *BudgetTracker) AddTransaction(amount float64, category, tType string) {
	newTransaction := Transaction{
		ID:       bt.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}
	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextID++
}

// Display all transactions
func (bt BudgetTracker) DisplayTransactions() {
	fmt.Println("ID\tAmount\tCategory\tDate\t\tType")
	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n", transaction.ID, transaction.Amount, transaction.Category, transaction.Date.Format("2006-01-02"), transaction.Type)
	}
}

// Get total income or expenses
func (bt BudgetTracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

// Save transactions to a CSV file
func (bt BudgetTracker) SaveToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"})

	// Write data
	for _, t := range bt.transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("2006-01-02"),
			t.Type,
		}
		writer.Write(record)
	}
	fmt.Println("Transactions saved to", filename)
	return nil
}

// Main function
func main() {
	bt := BudgetTracker{}

	for {
		fmt.Println("\n--- Personal Budget Tracker ---")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Display Transactions")
		fmt.Println("3. Show Total Income")
		fmt.Println("4. Show Total Expenses")
		fmt.Println("5. Save Transactions to CSV")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount: ")
			var amount float64
			fmt.Scanln(&amount)

			fmt.Print("Enter category: ")
			var category string
			fmt.Scanln(&category)

			fmt.Print("Enter type (income/expense): ")
			var tType string
			fmt.Scanln(&tType)

			bt.AddTransaction(amount, category, tType)
			fmt.Println("Transaction added!")

		case 2:
			bt.DisplayTransactions()

		case 3:
			fmt.Printf("Total Income: %.2f\n", bt.CalculateTotal("income"))

		case 4:
			fmt.Printf("Total Expenses: %.2f\n", bt.CalculateTotal("expense"))

		case 5:
			fmt.Print("Enter filename (e.g., transactions.csv): ")
			var filename string
			fmt.Scanln(&filename)
			if err := bt.SaveToCSV(filename); err != nil {
				fmt.Println("Error saving transactions:", err)
			}

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
