# Hurl GraphQL API Test Files

This directory contains Hurl test files for all GraphQL operations in the monolith payment gateway. Each file sends `POST` requests to the GraphQL endpoint (`/query`) with query/mutation payloads.

## Available Test Files

| Service | File | Description |
|---------|------|-------------|
| Auth | `auth/auth.hurl` | Full auth flow: login → getMe → refreshToken |
| Auth (Login) | `auth/login.hurl` | Login mutation with token capture |
| Auth (Register) | `auth/register.hurl` | User registration mutation |
| Card | `card.hurl` | Card CRUD, dashboard, balance/topup/transaction stats |
| Merchant | `merchant.hurl` | Merchant CRUD, transaction stats, payment methods |
| Merchant Document | `merchantdocument.hurl` | Merchant document management |
| Role | `role.hurl` | Role CRUD with soft-delete/restore |
| Saldo | `saldo.hurl` | Balance management, monthly/yearly stats |
| Topup | `topup.hurl` | Top-up operations, status & method analytics |
| Transaction | `transaction.hurl` | Transaction CRUD, payment methods, amounts |
| Transfer | `transfer.hurl` | Transfer operations, sender/receiver lookups |
| User | `user.hurl` | User CRUD with soft-delete/restore |
| Withdraw | `withdraw.hurl` | Withdrawal operations, status analytics |

## Prerequisites

1. **Install Hurl**: Follow the installation guide at [https://hurl.dev/docs/installation.html](https://hurl.dev/docs/installation.html)

2. **Start the GraphQL API Gateway**: Make sure the API Gateway is running on `http://localhost:5000`

## Usage

### Running Individual Tests

```bash
# Test auth flow
hurl auth/auth.hurl

# Test card operations
hurl card.hurl

# Test transactions
hurl transaction.hurl
```

### Running All Tests

```bash
# Run all root-level test files
hurl *.hurl

# Run all tests including auth subdirectory
./run_tests.sh
```

### Variables

The test files use placeholder variables:

- `{{auth_token}}` — Valid JWT access token
- `{{refresh_token}}` — Valid JWT refresh token

### Getting Auth Tokens

1. Register a user:
```bash
hurl auth/register.hurl
```

2. Login to get tokens:
```bash
hurl auth/login.hurl
```

The login test captures `auth_token` and `refresh_token` automatically.

### Using Environment Variables

```bash
export AUTH_TOKEN="your_jwt_token_here"
export REFRESH_TOKEN="your_refresh_token_here"

hurl card.hurl
```

## Test Structure

All tests use a single endpoint — `POST http://localhost:5000/query` — with GraphQL payloads:

```
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createCard(input: { ... }) { status message data { ... } } }"
}

HTTP 200
```

Each test file covers:

1. **Mutations** — Create, update, soft-delete (trash), restore, permanent delete
2. **Queries** — Find all (paginated), find by ID, analytics (monthly/yearly stats)

## GraphQL Endpoint

| Property | Value |
|----------|-------|
| **URL** | `http://localhost:5000/query` |
| **Method** | `POST` |
| **Content-Type** | `application/json` |
| **Auth** | `Authorization: Bearer <token>` |

## Troubleshooting

### Connection Issues
- Ensure the GraphQL API Gateway is running on port 5000
- Test with: `curl -X POST http://localhost:5000/query -H "Content-Type: application/json" -d '{"query":"{ __typename }"}'`

### Authentication Issues
- Run `auth/login.hurl` first to capture a valid token
- Ensure tokens are not expired
- Check the `Authorization: Bearer` header format

### GraphQL Errors
- Response errors appear in `$.errors` (not HTTP status codes)
- Use `[Asserts]` with `jsonpath` to validate `$.data.<operation>.status`

For more advanced Hurl features, see the [Hurl documentation](https://hurl.dev/docs/).