# Role Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new role
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createRole(input: { name: \"admin\" }) { status message data { id name created_at updated_at } } }"
}

HTTP 200

### Find all roles (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllRole(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id name created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find role by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdRole(input: { role_id: 1 }) { status message data { id name created_at updated_at } } }"
}

HTTP 200

### Find roles by user ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByUserIdRole(input: { user_id: 1 }) { status message data { id name created_at updated_at } } }"
}

HTTP 200

### Update role
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateRole(input: { id: 1, name: \"super_admin\" }) { status message data { id name created_at updated_at } } }"
}

HTTP 200

### Find active roles
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByActiveRole(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id name created_at updated_at deleted_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find trashed roles
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByTrashedRole(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id name created_at updated_at deleted_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Trash role (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedRole(input: { role_id: 1 }) { status message data { id name created_at updated_at deleted_at } } }"
}

HTTP 200

### Restore role
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreRole(input: { role_id: 1 }) { status message data { id name created_at updated_at deleted_at } } }"
}

HTTP 200

### Delete role permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteRolePermanent(input: { role_id: 1 }) { status message } }"
}

HTTP 200