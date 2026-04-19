# Expense Tracker

A full-stack expense tracking application built with Flutter and Go.

## Tech Stack

- **Frontend:** Flutter (Dart) — cross-platform mobile/web UI
- **Backend:** Go with Gin framework — REST API
- **Database:** SQLite — lightweight local persistence

## Features

- Add and delete expenses
- Categorize expenses (Food, Transport, Housing, Health, Entertainment)
- Real-time total calculation
- REST API with full CRUD operations
- Clean architecture separating frontend, backend and data layers

## Project Structure
expense_tracker/     # Flutter frontend
lib/
models/          # Data models
services/        # API communication layer
screens/         # UI screens
expense_api/         # Go backend
main.go            # Server setup and routes
database.go        # Database initialization
handlers.go        # CRUD request handlers

## Running Locally

### Backend (Go API)
```bash
cd expense_api
go run .
# Server starts on http://localhost:8080
```

### Frontend (Flutter)
```bash
cd expense_tracker
flutter run -d chrome
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /expenses | List all expenses |
| POST | /expenses | Create new expense |
| DELETE | /expenses/:id | Delete an expense |