# Merchant Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new merchant
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createMerchant(input: { name: \"Test Store\", user_id: 1 }) { status message data { id name api_key status user_id created_at updated_at } } }"
}

HTTP 200

### Find all merchants (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllMerchant(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id name api_key status user_id created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find merchant by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdMerchant(input: { merchant_id: 1 }) { status message data { id name api_key status user_id created_at updated_at } } }"
}

HTTP 200

### Find merchant by API key
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByApiKey(input: { api_key: \"test-api-key-123\" }) { status message data { id name api_key status user_id created_at updated_at } } }"
}

HTTP 200

### Find merchants by user ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByMerchantUserId(input: { user_id: 1 }) { status message data { id name api_key status user_id created_at updated_at } } }"
}

HTTP 200

### Update merchant
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateMerchant(input: { merchant_id: 1, name: \"Updated Store Name\", user_id: 1, status: \"active\" }) { status message data { id name api_key status user_id created_at updated_at } } }"
}

HTTP 200

### Update merchant status
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateMerchantStatus(input: { merchant_id: 1, status: \"active\" }) { status message data { id name api_key status user_id created_at updated_at } } }"
}

HTTP 200

### Find all merchant transactions
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllTransactionMerchant(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id card_number amount payment_method merchant_id merchant_name transaction_time created_at updated_at deleted_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find monthly payment methods (merchant)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyPaymentMethodsMerchant(input: { year: 2026 }) { status message data { month payment_method total_amount } } }"
}

HTTP 200

### Find yearly amount (merchant)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findYearlyAmountMerchant(input: { year: 2026 }) { status message data { year total_amount } } }"
}

HTTP 200

### Find monthly total amount (merchant)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTotalAmountMerchant(input: { year: 2026 }) { status message data { month year total_amount } } }"
}

HTTP 200

### Trash merchant (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedMerchant(input: { merchant_id: 1 }) { status message data { id name api_key status user_id created_at updated_at deleted_at } } }"
}

HTTP 200

### Restore merchant
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreMerchant(input: { merchant_id: 1 }) { status message data { id name api_key status user_id created_at updated_at deleted_at } } }"
}

HTTP 200

### Delete merchant permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteMerchantPermanent(input: { merchant_id: 1 }) { status message } }"
}

HTTP 200