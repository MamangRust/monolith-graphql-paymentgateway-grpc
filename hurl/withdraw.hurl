# Withdraw Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new withdrawal
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createWithdraw(input: { card_number: \"4111111111111111\", withdraw_amount: 30000, withdraw_time: \"2026-05-20T12:00:00Z\" }) { status message data { withdraw_id withdraw_no card_number withdraw_amount withdraw_time created_at updated_at } } }"
}

HTTP 200

### Find all withdrawals (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllWithdraw(input: { page: 1, page_size: 10, search: \"\" }) { status message data { withdraw_id withdraw_no card_number withdraw_amount withdraw_time created_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find withdrawal by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByIdWithdraw(input: { withdraw_id: 1 }) { status message data { withdraw_id withdraw_no card_number withdraw_amount withdraw_time created_at updated_at } } }"
}

HTTP 200

### Update withdrawal
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateWithdraw(input: { withdraw_id: 1, card_number: \"4111111111111111\", withdraw_amount: 35000, withdraw_time: \"2026-05-20T13:00:00Z\" }) { status message data { withdraw_id withdraw_no card_number withdraw_amount withdraw_time created_at updated_at } } }"
}

HTTP 200

### Find monthly withdraw status success
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyWithdrawStatusSuccess(input: { year: 2026, month: 5 }) { status message data { year month total_success total_amount } } }"
}

HTTP 200

### Find yearly withdraw status success
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findYearlyWithdrawStatusSuccess(input: { year: 2026 }) { status message data { year total_success total_amount } } }"
}

HTTP 200

### Find monthly withdraws
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMonthlyWithdraws(input: { year: 2026 }) { status message data { month total_amount } } }"
}

HTTP 200

### Find by card number
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findByCardNumberWithdraw(input: { card_number: \"4111111111111111\" }) { status message data { withdraw_id withdraw_no card_number withdraw_amount withdraw_time created_at updated_at } } }"
}

HTTP 200

### Trash withdrawal (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedWithdraw(input: { withdraw_id: 1 }) { status message data { withdraw_id withdraw_no card_number withdraw_amount deleted_at } } }"
}

HTTP 200

### Restore withdrawal
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreWithdraw(input: { withdraw_id: 1 }) { status message data { withdraw_id withdraw_no card_number withdraw_amount deleted_at } } }"
}

HTTP 200

### Delete withdrawal permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteWithdrawPermanent(input: { withdraw_id: 1 }) { status message } }"
}

HTTP 200