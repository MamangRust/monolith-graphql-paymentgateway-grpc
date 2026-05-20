# Saldo (Balance) Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new saldo
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createSaldo(input: { card_number: \"4111111111111111\", total_balance: 100000 }) { status message data { saldo_id card_number total_balance created_at updated_at } } }"
}

HTTP 200

### Find all saldos (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllSaldo(input: { page: 1, page_size: 10, search: \"\" }) { status message data { saldo_id card_number total_balance created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find saldo by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdSaldo(input: { saldo_id: 1 }) { status message data { saldo_id card_number total_balance created_at updated_at } } }"
}

HTTP 200

### Find saldo by card number
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByCardNumberSaldo(card_number: \"4111111111111111\") { status message data { saldo_id card_number total_balance created_at updated_at } } }"
}

HTTP 200

### Update saldo
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateSaldo(input: { saldo_id: 1, card_number: \"4111111111111111\", total_balance: 150000 }) { status message data { saldo_id card_number total_balance created_at updated_at } } }"
}

HTTP 200

### Find monthly total saldo balance
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTotalSaldoBalance(input: { year: 2026, month: 5 }) { status message data { month year total_balance } } }"
}

HTTP 200

### Find yearly total saldo balance
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findYearTotalSaldoBalance(input: { year: 2026 }) { status message data { year total_balance } } }"
}

HTTP 200

### Trash saldo (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedSaldo(input: { saldo_id: 1 }) { status message data { saldo_id card_number total_balance deleted_at } } }"
}

HTTP 200

### Restore saldo
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreSaldo(input: { saldo_id: 1 }) { status message data { saldo_id card_number total_balance deleted_at } } }"
}

HTTP 200

### Delete saldo permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteSaldoPermanent(input: { saldo_id: 1 }) { status message } }"
}

HTTP 200