🛒 GoCommerce – E-Commerce Platform in Golang

GoCommerce is a scalable and secure E-Commerce backend platform built with Golang and the Gin framework.
It provides APIs for user authentication, product management, shopping cart, and order processing, ensuring a smooth online shopping experience.

✨ Features

🔑 JWT Authentication – Secure user login & registration

🛍️ Product Management – Create, update, and fetch product details

🛒 Shopping Cart – Add, update, and remove items

📦 Order Management – Place and manage customer orders

⚡ High Performance – Powered by Go’s concurrency and Gin’s lightweight router

🗄️ MongoDB Integration – Scalable and flexible database storage

🛠️ Tech Stack

Language: Go (Golang)

Framework: Gin

Database: MongoDB

Authentication: JWT (JSON Web Tokens)

📂 Project Structure
go-ecommerce/
│── main.go               # Entry point of the application
│── go.mod                # Go module file
│── generate/             # Token generation (JWT)
│   └── token.go
│── controllers/          # API logic (users, products, orders, cart)
│── models/               # Data models
│── routes/               # API routes
│── config/               # Database connection & configurations
│── README.md             # Project documentation



