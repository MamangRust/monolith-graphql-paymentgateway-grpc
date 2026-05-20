# Topup Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new topup
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createTopup(input: { card_number: \"4111111111111111\", topup_no: \"TOPUP123456\", topup_amount: 50000, topup_method: \"bank_transfer\" }) { status message data { id card_number topup_no topup_amount topup_method topup_time created_at updated_at } } }"
}

HTTP 200

### Find all topups (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllTopup(input: { page: 1, page_size: 10, search: \"\" }) { status message data { id card_number topup_no topup_amount topup_method topup_time created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find topup by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdTopup(input: { topup_id: 1 }) { status message data { id card_number topup_no topup_amount topup_method topup_time created_at updated_at } } }"
}

HTTP 200

### Update topup
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateTopup(input: { topup_id: 1, card_number: \"4111111111111111\", topup_amount: 60000, topup_method: \"credit_card\" }) { status message data { id card_number topup_no topup_amount topup_method topup_time created_at updated_at } } }"
}

HTTP 200

### Find monthly topup status success
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTopupStatusSuccess(input: { year: 2026, month: 5 }) { status message data { year month total_success total_amount } } }"
}

HTTP 200

### Find yearly topup status success
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findYearlyTopupStatusSuccess(input: { year: 2026 }) { status message data { year total_success total_amount } } }"
}

HTTP 200

### Find monthly topup methods
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTopupMethods(input: { year: 2026 }) { status message data { month topup_method total_topups total_amount } } }"
}

HTTP 200

### Find monthly topup amounts
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyTopupAmounts(input: { year: 2026 }) { status message data { month total_amount } } }"
}

HTTP 200

### Trash topup (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedTopup(input: { topup_id: 1 }) { status message data { id card_number topup_no topup_amount topup_method deleted_at } } }"
}

HTTP 200

### Restore topup
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreTopup(input: { topup_id: 1 }) { status message data { id card_number topup_no topup_amount topup_method deleted_at } } }"
}

HTTP 200

### Delete topup permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteTopupPermanent(input: { topup_id: 1 }) { status message } }"
}

HTTP 200