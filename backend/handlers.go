package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Expense struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Date     string  `json:"date"`
}

func getExpenses(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, amount, category, date FROM expenses ORDER BY date DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Category, &e.Date)
		expenses = append(expenses, e)
	}

	if expenses == nil {
		expenses = []Expense{}
	}

	c.JSON(http.StatusOK, expenses)
}

func createExpense(c *gin.Context) {
	var expense Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	result, err := db.Exec(
		"INSERT INTO expenses (title, amount, category, date) VALUES (?, ?, ?, ?)",
		expense.Title, expense.Amount, expense.Category, expense.Date,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	id, _ := result.LastInsertId()
	expense.ID = int(id)
	c.JSON(http.StatusCreated, expense)
}

func deleteExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = db.Exec("DELETE FROM expenses WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}
