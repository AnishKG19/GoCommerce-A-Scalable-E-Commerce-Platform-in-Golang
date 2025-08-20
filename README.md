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

🚀 Getting Started
1. Clone the Repository
git clone https://github.com/yourusername/go-ecommerce.git
cd go-ecommerce

2. Initialize Go Modules
go mod tidy

3. Setup Environment Variables

Create a .env file in the root directory:

MONGO_URI=mongodb://localhost:27017
JWT_SECRET=your_secret_key
PORT=8080

4. Run the Application
go run main.go


The server will start at:
👉 http://localhost:8080

📡 API Endpoints
🔑 Authentication

POST /signup – Register a new user

POST /login – Login user and return JWT

🛍️ Products

GET /products – Fetch all products

POST /products – Add a new product

PUT /products/:id – Update product details

🛒 Cart

POST /cart/add – Add item to cart

GET /cart – Get user’s cart

DELETE /cart/:id – Remove item from cart

📦 Orders

POST /orders – Place a new order

GET /orders – View user’s orders

🧪 Example Request (Add to Cart)
POST /cart/add
Content-Type: application/json
Authorization: Bearer <JWT_TOKEN>

{
  "product_id": "64c9a1e3f",
  "quantity": 2
}

🤝 Contributing

Contributions are welcome! Please fork the repo and submit a pull request.

📜 License

This project is licensed under the MIT License.

🔥 With this backend, you can extend the project into a full-stack e-commerce app by connecting it with a React/Next.js frontend.

Would you like me to also generate a sample main.go starter code (with routes + Gin setup) so the README links directly to runnable code?
