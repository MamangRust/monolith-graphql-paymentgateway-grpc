# Transfer Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new transfer
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createTransfer(input: { transfer_from: \"4111111111111111\", transfer_to: \"4222222222222222\", transfer_amount: 25000 }) { status message data { id transfer_no transfer_from transfer_to transfer_amount transfer_time created_at updated_at } } }"
}

HTTP 200

### Find all transfers (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllTransfer(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id transfer_no transfer_from transfer_to transfer_amount transfer_time created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find transfer by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdTransfer(input: { transfer_id: 1 }) { status message data { id transfer_no transfer_from transfer_to transfer_amount transfer_time created_at updated_at } } }"
}

HTTP 200

### Update transfer
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateTransfer(input: { transfer_id: 1, transfer_from: \"4111111111111111\", transfer_to: \"4333333333333333\", transfer_amount: 30000 }) { status message data { id transfer_no transfer_from transfer_to transfer_amount transfer_time created_at updated_at } } }"
}

HTTP 200

### Find transfer by sender
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findTransferByTransferFrom(input: { transfer_from: \"4111111111111111\" }) { status message data { id transfer_no transfer_from transfer_to transfer_amount transfer_time created_at updated_at } } }"
}

HTTP 200

### Find transfer by receiver
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findTransferByTransferTo(input: { transfer_to: \"4222222222222222\" }) { status message data { id transfer_no transfer_from transfer_to transfer_amount transfer_time created_at updated_at } } }"
}

HTTP 200

### Find monthly transfer status success
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTransferStatusSuccess(input: { year: 2026, month: 5 }) { status message data { year month total_success total_amount } } }"
}

HTTP 200

### Find monthly transfer amounts
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTransferAmounts(input: { year: 2026 }) { status message data { month total_amount } } }"
}

HTTP 200

### Trash transfer (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedTransfer(input: { transfer_id: 1 }) { status message data { id transfer_no transfer_from transfer_to transfer_amount deleted_at } } }"
}

HTTP 200

### Restore transfer
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreTransfer(input: { transfer_id: 1 }) { status message data { id transfer_no transfer_from transfer_to transfer_amount deleted_at } } }"
}

HTTP 200

### Delete transfer permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteTransferPermanent(input: { transfer_id: 1 }) { status message } }"
}

HTTP 200