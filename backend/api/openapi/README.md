# Angidi API Documentation

This directory contains the OpenAPI specification for the Angidi e-commerce platform API.

## OpenAPI Specification

The complete API specification is available in:
- [`openapi.yaml`](./openapi.yaml) - OpenAPI 3.0.3 specification

## Viewing the Documentation

### Option 1: Swagger UI (Online)

Visit [Swagger Editor](https://editor.swagger.io/) and paste the contents of `openapi.yaml`.

### Option 2: Swagger UI (Local)

If you have Docker installed:

```bash
docker run -p 8081:8080 -e SWAGGER_JSON=/openapi/openapi.yaml -v $(pwd):/openapi swaggerapi/swagger-ui
```

Then open http://localhost:8081 in your browser.

### Option 3: Redoc (Local)

If you have Node.js installed:

```bash
npx redoc-cli serve openapi.yaml
```

## API Overview

### Base URL

- **Local Development**: `http://localhost:8080`
- **API Version 1**: `http://localhost:8080/api/v1`

### Authentication

Most endpoints require authentication using JWT tokens. Include the access token in the Authorization header:

```
Authorization: ******
```

Tokens can be obtained via:
- `POST /api/v1/users/login` - Returns access and refresh tokens
- `POST /api/v1/users/refresh-token` - Refreshes expired access tokens

### API Endpoints

#### Health
- `GET /health` - Health check

#### Authentication
- `POST /api/v1/users/register` - Register new user
- `POST /api/v1/users/login` - Login user
- `POST /api/v1/users/refresh-token` - Refresh access token

#### Users
- `GET /api/v1/users/me` - Get current user profile (protected)
- `PUT /api/v1/users/me` - Update user profile (protected)

#### Products
- `GET /api/v1/products` - List products (public)
- `GET /api/v1/products/{id}` - Get product by ID (public)
- `POST /api/v1/products` - Create product (admin only)
- `PUT /api/v1/products/{id}` - Update product (admin only)
- `DELETE /api/v1/products/{id}` - Delete product (admin only)

### Error Responses

All error responses follow a consistent format:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": [
      {
        "field": "field_name",
        "message": "Validation error message"
      }
    ],
    "request_id": "unique-request-id"
  }
}
```

### Common Error Codes

- `VALIDATION_ERROR` (400) - Request validation failed
- `INVALID_CREDENTIALS` (401) - Invalid email or password
- `UNAUTHORIZED` (401) - Missing or invalid authentication token
- `FORBIDDEN` (403) - Insufficient permissions
- `NOT_FOUND` (404) - Resource not found
- `CONFLICT` (409) - Resource already exists (e.g., duplicate email)
- `RATE_LIMIT_EXCEEDED` (429) - Too many requests
- `INTERNAL_ERROR` (500) - Internal server error

## Examples

### Register a New User

```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "SecurePass123!",
    "name": "John Doe"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "SecurePass123!"
  }'
```

### Get User Profile

```bash
curl -X GET http://localhost:8080/api/v1/users/me \
  -H "Authorization: ******"
```

### List Products

```bash
curl -X GET "http://localhost:8080/api/v1/products?page=1&page_size=10&search=phone"
```

### Create Product (Admin)

```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -H "Authorization: ******" \
  -d '{
    "name": "New Product",
    "description": "Product description",
    "price": 99.99,
    "stock": 100,
    "category_id": "electronics",
    "image_url": "https://example.com/image.jpg"
  }'
```

## Validation Rules

### User Registration
- **Email**: Valid email format, max 255 characters
- **Password**: Min 8 characters, max 128 characters
- **Name**: Min 2 characters, max 100 characters

### Product Creation
- **Name**: Min 3 characters, max 255 characters
- **Description**: Max 2000 characters
- **Price**: Must be greater than 0
- **Stock**: Must be >= 0
- **Category ID**: Required
- **Image URL**: Valid URL format (optional)

## Rate Limiting

The API implements rate limiting to prevent abuse:
- **Default**: 100 requests per minute
- **Login endpoint**: 5 attempts per minute

Rate limit information is included in response headers:
- `X-RateLimit-Limit`: Maximum requests allowed
- `X-RateLimit-Remaining`: Requests remaining in current window
- `X-RateLimit-Reset`: Time when the rate limit resets

## Versioning

The API uses URL-based versioning:
- Current version: `v1`
- Base path: `/api/v1`

Future versions will be available at `/api/v2`, etc.

## Support

For issues or questions about the API:
- GitHub Issues: https://github.com/yesoreyeram/angidi-demo-app/issues
- Documentation: See the main README.md

## License

MIT License - see LICENSE file for details
