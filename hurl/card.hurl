# Card Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new card
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createCard(input: { user_id: 1, card_type: \"credit\", expire_date: \"12/25\", cvv: \"123\", card_provider: \"visa\" }) { status message data { id user_id card_number card_type expire_date cvv card_provider created_at updated_at } } }"
}

HTTP 200

### Find all cards (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllCard(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id user_id card_number card_type expire_date cvv card_provider created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find card by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdCard(input: { card_id: 1 }) { status message data { id user_id card_number card_type expire_date cvv card_provider created_at updated_at } } }"
}

HTTP 200

### Find card by user ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByUserIdCard(input: { user_id: 1 }) { status message data { id user_id card_number card_type expire_date cvv card_provider created_at updated_at } } }"
}

HTTP 200

### Update card
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateCard(input: { card_id: 1, user_id: 1, card_type: \"debit\", expire_date: \"12/26\", cvv: \"456\", card_provider: \"mastercard\" }) { status message data { id user_id card_number card_type expire_date cvv card_provider created_at updated_at } } }"
}

HTTP 200

### Dashboard card
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { dashboardCard { status message data { total_balance total_topup total_withdraw total_transaction total_transfer } } }"
}

HTTP 200

### Dashboard card by card number
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { dashboardCardNumber(input: { card_number: \"4111111111111111\" }) { status message data { total_balance total_topup total_withdraw total_transaction total_transfer_send total_transfer_receiver } } }"
}

HTTP 200

### Find monthly balance
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyBalance(input: { year: 2026 }) { status message data { month total_balance } } }"
}

HTTP 200

### Find yearly balance
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findYearlyBalance(input: { year: 2026 }) { status message data { year total_balance } } }"
}

HTTP 200

### Find monthly topup amount
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTopupAmount(input: { year: 2026 }) { status message data { month total_amount } } }"
}

HTTP 200

### Find monthly transaction amount
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTransactionAmount(input: { year: 2026 }) { status message data { month total_amount } } }"
}

HTTP 200

### Find monthly withdraw amount
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyWithdrawAmount(input: { year: 2026 }) { status message data { month total_amount } } }"
}

HTTP 200

### Trash card (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedCard(input: { card_id: 1 }) { status message data { id user_id card_number card_type expire_date cvv card_provider created_at updated_at deleted_at } } }"
}

HTTP 200

### Restore card
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreCard(input: { card_id: 1 }) { status message data { id user_id card_number card_type expire_date cvv card_provider created_at updated_at deleted_at } } }"
}

HTTP 200

### Delete card permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteCardPermanent(input: { card_id: 1 }) { status message } }"
}

HTTP 200