# 📚 Library Book Rent API

**Library Book Rent API** is a powerful backend service built with **Go (Golang)** to manage book rentals in a library system. It handles core functionalities such as book inventory, user management, rental workflows, and integrates with a secure **payment gateway** for fee processing.

---

## 🚀 Features

- 🔎 Book search & availability tracking  
- 👤 User registration & authentication  
- 📅 Rent, return, and renewal of books  
- 💳 Integration with payment gateway for deposits and payment
- 📈 Real-time status updates and transaction history  
- 🛡️ Secure RESTful APIs with token-based authentication  

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)  
- **Database**: PostgreSQL / MySQL (configurable)  
- **API**: RESTful using `gorilla/mux` or `gin`  
- **Authentication**: JWT  
- **Payments**: Stripe / Razorpay (modular integration)  

---

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/library-book-rent-api.git
cd library-book-rent-api

# Install dependencies
go mod tidy

# Run the server
go run main.go
