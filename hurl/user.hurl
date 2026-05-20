# User Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new user
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createUser(input: { firstname: \"Jane\", lastname: \"Smith\", email: \"jane.smith@example.com\", password: \"password123\", confirm_password: \"password123\" }) { status message data { id firstname lastname email created_at updated_at } } }"
}

HTTP 200

### Find all users (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllUser(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id firstname lastname email created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find user by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdUser(input: { id: 1 }) { status message data { id firstname lastname email created_at updated_at } } }"
}

HTTP 200

### Update user
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateUser(input: { id: 1, firstname: \"Jane Marie\", lastname: \"Smith-Johnson\", email: \"jane.marie@example.com\", password: \"newpass123\", confirm_password: \"newpass123\" }) { status message data { id firstname lastname email created_at updated_at } } }"
}

HTTP 200

### Find active users
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByActiveUser(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id firstname lastname email created_at updated_at deleted_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find trashed users
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByTrashedUser(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id firstname lastname email created_at updated_at deleted_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Trash user (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedUser(input: { id: 1 }) { status message data { id firstname lastname email created_at updated_at deleted_at } } }"
}

HTTP 200

### Restore user
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreUser(input: { id: 1 }) { status message data { id firstname lastname email created_at updated_at deleted_at } } }"
}

HTTP 200

### Delete user permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteUserPermanent(input: { id: 1 }) { status message } }"
}

HTTP 200