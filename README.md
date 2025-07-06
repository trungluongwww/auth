# Social Authentication API

A comprehensive Go-based authentication and social media API built with clean architecture principles.

## Features

### Authentication
- **Email/Password Registration & Login**
- **Facebook OAuth Integration**
- **JWT Token-based Authentication**
- **Refresh Token Mechanism**

### Social Features
- **Posts**: Create, read, update, delete posts with privacy controls
- **Comments**: Hierarchical comments with replies
- **Likes**: Like/unlike posts and comments
- **User Profiles**: View user posts and profiles
- **Search**: Search posts by title and content
- **Privacy**: Public/private post visibility

## API Endpoints

### Authentication
- `POST /api/v1/register` - User registration
- `POST /api/v1/login` - Email/password login
- `POST /api/v1/facebook` - Facebook OAuth login
- `POST /api/v1/refresh-token` - Token refresh
- `GET /api/v1/users/me` - Get current user info

### Posts (Authenticated)
- `POST /api/v1/posts` - Create a new post
- `GET /api/v1/posts` - Get all public posts (with pagination and search)
- `GET /api/v1/posts/:id` - Get specific post
- `PUT /api/v1/posts/:id` - Update post
- `DELETE /api/v1/posts/:id` - Delete post
- `GET /api/v1/users/:userId/posts` - Get user's posts
- `POST /api/v1/posts/like` - Like a post
- `DELETE /api/v1/posts/like` - Unlike a post

## Database Schema

### Core Tables
- **accounts**: Stores hashed passwords and signup info
- **users**: User profile information linked to accounts
- **account_refresh_tokens**: JWT refresh token management
- **user_facebook_logins**: Facebook OAuth data

### Social Tables
- **posts**: User posts with title, content, privacy settings
- **comments**: Hierarchical comments on posts
- **post_likes**: Post like relationships
- **comment_likes**: Comment like relationships
- **user_follows**: User following relationships

## Steps to Run the Project

1. **Setup Environment Variables**
   ```bash
   # Copy and configure environment variables
   cp .env.example .env
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Database Setup**
   - Configure MySQL database
   - Update `gen.tool` with your database credentials
   - Run migrations: `go run cmd/main.go` (applies migrations automatically)

4. **Generate Models**
   ```bash
   # Install Gentool
   go install gorm.io/gen/tools/gentool@latest
   
   # Generate models
   make gen
   # or
   gentool -c gen.tool
   ```

5. **Start the Project**
   ```bash
   # Using Go
   go run cmd/main.go
   
   # Using Makefile
   make run
   ```

## Technology Stack

- **Framework**: Echo (HTTP framework)
- **ORM**: GORM with MySQL driver
- **Authentication**: JWT with golang-jwt
- **Password Hashing**: bcrypt
- **Validation**: go-playground/validator
- **Configuration**: envconfig
- **Migrations**: Goose
- **OAuth**: Facebook SDK
- **Code Generation**: GORM Gen for models

## Architecture

The project follows clean architecture principles:

- **Handlers**: HTTP request/response handling
- **Use Cases**: Business logic implementation
- **Services**: Domain services and data transformation
- **Repository**: Data access layer
- **Models**: Database models and DTOs

## Development

### Code Generation
After adding new database tables:
1. Add migration files in `mysql/ddl/`
2. Run the application to apply migrations
3. Update `gen.tool` to include new tables
4. Run `make gen` to generate models

### Linting
```bash
make check-lint
```

### Docker
```bash
docker build -t auth-api .
docker run -p 5000:5000 auth-api
```
