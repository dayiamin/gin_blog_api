# 📝 Gin Blog API

A simple and extensible blog backend built using [Golang](https://golang.org/) and the [Gin](https://github.com/gin-gonic/gin) web framework. This API supports features such as user registration, login with JWT authentication, creating posts, commenting, and managing user profiles — all powered by a SQLite database and GORM ORM.

---

## 🚀 Features

- 🔐 JWT-based Authentication (Login & Register)
- 🧑 User CRUD (Create, Read, Update)
- 📝 Post CRUD with user association
- 💬 Comment system on posts
- 🗃 SQLite as the database
- 📦 GORM for ORM and migrations
- 🧪 Basic structure ready for unit testing

---

## 🗂️ Project Structure

```
.
├── database/         # DB connection & migration logic
├── docs/             # Swagger Docs
├── handlers/         # Route handler functions (user, post)
├── middleware/       # JWT auth middleware
├── models/           # GORM models (User, Post, Comment)
├── routes/           # Route setup
├── utils/            # JWT utilities (token generation, parsing)
├── main.go           # App entry point
├── go.mod            # Go modules
├── .env              # Environment variables
```

---

## 🛠️ Installation & Run

1. **Clone the repository**

```bash
git clone https://github.com/dayiamin/gin_blog_api.git
cd gin_blog_api
```

2. **Set up environment variables**

Create a `.env` file in the root directory:

```env
JWT_SECRET=your_jwt_secret
```

3. **Run the application**

```bash
go run main.go
```

4. **Access the API**

API will run at: `http://localhost:8080`

---

## 📬 API Endpoints

### Authentication
| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| POST   | `/register`      | Register new user   |
| POST   | `/login`         | Login user and get token |

### Users
| Method | Endpoint     | Description         |
|--------|--------------|---------------------|
| GET    | `/users`     | Get all users       |
| GET    | `/user/:user_name`  | Get user by user_name      |

### Posts
| Method | Endpoint     | Description             |
|--------|--------------|-------------------------|
| GET    | `/posts`     | Get all posts           |
| POST   | `/post`      | Create new post (auth)  |
| DELETE | `/post/:id`  | Delete post (auth)      |

### Comments (optional depending on implementation)
| Method | Endpoint           | Description          |
|--------|--------------------|----------------------|
| POST   | `/post/:id/comment`| Add comment to post  |
| DELETE | `/post/:id/comment/:id`  | Delete comment (auth)      |

---


## Swagger
you can go to /docs and use swagger docs for viewing endpoints 


## 🧑‍💻 Author

Developed by [@dayiamin](https://github.com/dayiamin)  
Feel free to contribute, raise issues, or fork the project.

---
