=== CURL API Documentation ===

This folder contains curl commands for testing all API endpoints in the authentication and social media application.

=============================================================================
FILE STRUCTURE
=============================================================================

1. auth.txt - Authentication endpoints (login, register, Facebook OAuth, refresh token)
2. user.txt - User-related endpoints (get current user profile)
3. post.txt - Post-related endpoints (CRUD operations, likes, search, pagination)
4. all_endpoints.txt - Complete reference of all current and future endpoints
5. README.txt - This file

=============================================================================
HOW TO USE
=============================================================================

1. Start your server (usually on http://localhost:8080)
2. Copy the curl command you want to test
3. Replace placeholders:
   - YOUR_ACCESS_TOKEN_HERE -> Your actual JWT token
   - user@example.com -> Valid email
   - password123 -> Valid password
   - facebook_access_token_here -> Valid Facebook token

=============================================================================
QUICK START
=============================================================================

1. Test server health:
   curl -X GET http://localhost:8080/api/v1/ping

2. Register a new user:
   curl -X POST http://localhost:8080/api/v1/register \
     -H "Content-Type: application/json" \
     -d '{"email":"test@example.com","name":"Test User","password":"password123"}'

3. Login to get access token:
   curl -X POST http://localhost:8080/api/v1/login \
     -H "Content-Type: application/json" \
     -d '{"email":"test@example.com","password":"password123"}'

4. Use the access token for authenticated requests:
   curl -X GET http://localhost:8080/api/v1/users/me \
     -H "Authorization: Bearer YOUR_ACCESS_TOKEN_HERE"

=============================================================================
AUTHENTICATION FLOW
=============================================================================

1. Register -> Get user account
2. Login -> Get access token and refresh token
3. Use access token in Authorization header for protected endpoints
4. When access token expires, use refresh token to get new tokens

=============================================================================
COMMON HEADERS
=============================================================================

For JSON requests:
-H "Content-Type: application/json"

For authenticated requests:
-H "Authorization: Bearer YOUR_ACCESS_TOKEN_HERE"

=============================================================================
ERROR HANDLING
=============================================================================

- 400: Bad Request (validation errors)
- 401: Unauthorized (invalid/missing token)
- 403: Forbidden (insufficient permissions)
- 404: Not Found (resource doesn't exist)
- 500: Internal Server Error

=============================================================================
TESTING TIPS
=============================================================================

1. Use a tool like Postman or Insomnia for easier testing
2. Save your access token in a variable for reuse
3. Test error cases (invalid data, missing fields, etc.)
4. Test authentication edge cases (expired tokens, invalid tokens)
5. Test pagination with different page/limit values
6. Test search functionality with various keywords

=============================================================================
ENVIRONMENT SETUP
=============================================================================

Make sure your server is running with:
- Database properly configured
- JWT secret configured
- Facebook OAuth configured (if testing Facebook login)
- All required environment variables set

=============================================================================
NOTES
=============================================================================

- All endpoints return JSON responses
- Dates are in ISO 8601 format
- IDs are integers
- Pagination uses page and limit parameters
- Search is case-insensitive
- Private posts are only visible to their owners
- View counts are automatically incremented
- Like counts are updated in real-time 