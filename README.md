# Go Todo App

A simple full-stack Todo application built with:

- **Backend:** Go (Fiber, MongoDB)
- **Frontend:** React + Chakra UI

## Features

- Add, update, and delete todos
- Persistent storage with MongoDB
- Responsive UI with Chakra UI
- Light/Dark mode support

## Getting Started

### Backend

1. **Install dependencies:**
   ```sh
   go mod tidy
   ```

2. **Set up environment variables:**
   - Create a `.env` file in the root directory:
     ```
     MONGODB_URI=your_mongodb_connection_string
     PORT=3000
     ```

3. **Run the server:**
   ```sh
   go run main.go
   ```
   Or, for automatic reload on code changes, you can use [Air](https://github.com/cosmtrek/air):

   ```sh
   # Install Air if you don't have it
   go install github.com/air-verse/air@latest

   # Then run
   air
   ```

### Frontend

1. **Install dependencies:**
   ```sh
   cd client
   npm install
   ```

2. **Start the React app:**
   ```sh
   npm run dev
   ```

3. **Visit:** http://localhost:5173