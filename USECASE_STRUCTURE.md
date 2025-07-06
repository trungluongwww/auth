# Use Case Structure Documentation

## Overview

The use cases have been reorganized with a single interface per domain but implementations separated into multiple files based on their functionality. This provides better code organization while maintaining a clean API.

## Structure

### User Use Cases

#### `pkg/usecase/user.go`
**Purpose**: Main interface and constructor
- Defines the complete `User` interface with all operations
- Contains the `userImpl` struct and `NewUser()` constructor
- Single source of truth for user operations

#### `pkg/usecase/user_find.go`
**Purpose**: Implementation of read-only operations
- `GetMe()` - Get current user information

#### `pkg/usecase/user_upsert.go`
**Purpose**: Implementation of create, update, and authentication operations
- `Register()` - User registration
- `Login()` - Email/password login
- `LoginWithFacebook()` - Facebook OAuth login
- `RefreshToken()` - Token refresh

### Post Use Cases

#### `pkg/usecase/post.go`
**Purpose**: Main interface and constructor
- Defines the complete `Post` interface with all operations
- Contains the `postImpl` struct and `NewPost()` constructor
- Single source of truth for post operations

#### `pkg/usecase/post_find.go`
**Purpose**: Implementation of read-only operations
- `GetPost()` - Get specific post with view count increment
- `GetPosts()` - Get paginated list of public posts with search
- `GetUserPosts()` - Get user's posts with privacy controls

#### `pkg/usecase/post_upsert.go`
**Purpose**: Implementation of create, update, delete, and like operations
- `CreatePost()` - Create new post
- `UpdatePost()` - Update existing post (authorization required)
- `DeletePost()` - Delete post (authorization required)
- `LikePost()` - Like a post
- `UnlikePost()` - Unlike a post

## Benefits

### 1. **Single Interface**
- One clear interface per domain
- All operations visible in one place
- Easy to understand the complete API

### 2. **Organized Implementations**
- Related functionality grouped in separate files
- Easier to locate specific implementations
- Better file structure

### 3. **Maintainability**
- Changes to find operations don't affect upsert operations
- Easier to add new operations to specific categories
- Reduced merge conflicts

### 4. **Code Organization**
- Interface definition in main file
- Implementations split by functionality
- Clear separation of concerns

### 5. **Backward Compatibility**
- Existing code continues to work unchanged
- Same constructor and interface usage
- No breaking changes

## Usage

### Creating Use Cases
```go
// Create user use case
user := NewUser(repository, authService, userService, social)

// Create post use case
post := NewPost(repository, postService, userService)
```

### Using the Interface
```go
// All operations available through single interface
user.GetMe(ctx, userID)           // Find operation
user.Register(ctx, payload)       // Upsert operation
user.Login(ctx, loginPayload)     // Upsert operation

post.GetPost(ctx, userID, postID) // Find operation
post.CreatePost(ctx, userID, payload) // Upsert operation
post.LikePost(ctx, userID, likePayload) // Upsert operation
```

## File Organization

```
pkg/usecase/
├── user.go           # User interface, struct, constructor
├── user_find.go      # User find implementations
├── user_upsert.go    # User upsert implementations
├── post.go           # Post interface, struct, constructor
├── post_find.go      # Post find implementations
└── post_upsert.go    # Post upsert implementations
```

## Future Extensions

This structure makes it easy to add new operation categories:

- `user_profile.go` - Profile management implementations
- `user_settings.go` - Settings management implementations
- `post_comment.go` - Comment-specific implementations
- `post_share.go` - Sharing implementations

Each new file would contain implementations for the same interface, keeping the API clean and organized. 