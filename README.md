# 📚 Library Book Rent API

**Library Book Rent API** is a powerful backend service built with **Go (Golang)** to manage book rentals in a library system. It handles core functionalities such as book inventory, user management, rental workflows, and integrates with a secure **payment gateway** for fee processing.

---

## 🚀 Features

- 🔎 Book search & availability tracking  
- 👤 User registration & authentication  
- 📅 Rent and return of books  
- 💳 Integration with payment gateway for deposits and payments
- 📈 Real-time status updates and transaction history  
- 🛡️ Secure RESTful APIs with token-based authentication  

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)  
- **Database**: PostgreSQL / MySQL (configurable)  
- **API**: RESTful using Echo Framework
- **Authentication**: JWT  
- **Payments**: App Balance / Xendit 

---

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/rayhanadri/book-rent-api.git
cd book-rent-api

# Install dependencies
go mod tidy

# Run the server
go run main.go
