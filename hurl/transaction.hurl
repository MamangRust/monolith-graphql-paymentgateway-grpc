# Transaction Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new transaction
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createTransaction(input: { api_key: \"test-api-key-123\", card_number: \"4111111111111111\", amount: 10000, payment_method: \"credit_card\", merchant_id: 1, transaction_time: \"2026-05-20T12:00:00Z\" }) { status message data { id card_number transaction_no amount payment_method merchant_id transaction_time created_at updated_at } } }"
}

HTTP 200

### Find all transactions (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllTransaction(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id card_number transaction_no amount payment_method merchant_id transaction_time created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find transaction by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdTransaction(input: { transaction_id: 1 }) { status message data { id card_number transaction_no amount payment_method merchant_id transaction_time created_at updated_at } } }"
}

HTTP 200

### Update transaction
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateTransaction(input: { transaction_id: 1, api_key: \"test-api-key-123\", card_number: \"4111111111111111\", amount: 15000, payment_method: \"debit_card\", merchant_id: 1, transaction_time: \"2026-05-20T13:00:00Z\" }) { status message data { id card_number transaction_no amount payment_method merchant_id transaction_time created_at updated_at } } }"
}

HTTP 200

### Find monthly transaction status success
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTransactionStatusSuccess(input: { year: 2026, month: 5 }) { status message data { year month total_success total_amount } } }"
}

HTTP 200

### Find yearly transaction status success
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findYearlyTransactionStatusSuccess(input: { year: 2026 }) { status message data { year total_success total_amount } } }"
}

HTTP 200

### Find monthly payment methods
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyPaymentMethods(input: { year: 2026 }) { status message data { month payment_method total_transactions total_amount } } }"
}

HTTP 200

### Find monthly amounts
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyAmounts(input: { year: 2026 }) { status message data { month total_amount } } }"
}

HTTP 200

### Find transaction by merchant ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findTransactionByMerchantId(input: { merchant_id: 1 }) { status message data { id card_number transaction_no amount payment_method merchant_id transaction_time created_at updated_at } } }"
}

HTTP 200

### Trash transaction (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedTransaction(input: { transaction_id: 1 }) { status message data { id card_number transaction_no amount payment_method merchant_id deleted_at } } }"
}

HTTP 200

### Restore transaction
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreTransaction(input: { transaction_id: 1 }) { status message data { id card_number transaction_no amount payment_method merchant_id deleted_at } } }"
}

HTTP 200

### Delete transaction permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteTransactionPermanent(input: { transaction_id: 1 }) { status message } }"
}

HTTP 200