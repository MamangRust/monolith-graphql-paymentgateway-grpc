# Auth Full Flow — GraphQL Mutations & Queries
# Endpoint: POST /query (GraphQL)

### Step 1: Login
POST http://localhost:5000/query
Content-Type: application/json

{
  "query": "mutation { loginUser(input: { email: \"john.doe@hellodotaa.com\", password: \"password123\" }) { status message data { access_token refresh_token } } }"
}

HTTP 200
[Asserts]
jsonpath "$.data.loginUser.status" == "success"
jsonpath "$.data.loginUser.data.access_token" exists
jsonpath "$.data.loginUser.data.refresh_token" exists
[Captures]
auth_token: jsonpath "$.data.loginUser.data.access_token"
refresh_token: jsonpath "$.data.loginUser.data.refresh_token"


### Step 2: Get Me
POST http://localhost:5000/query
Content-Type: application/json

{
  "query": "query { getMe(input: { access_token: \"{{auth_token}}\" }) { status message data { id firstname lastname email created_at updated_at } } }"
}

HTTP 200
[Asserts]
jsonpath "$.data.getMe.status" == "success"
jsonpath "$.data.getMe.data.email" == "john.doe@hellodotaa.com"


### Step 3: Refresh Token
POST http://localhost:5000/query
Content-Type: application/json

{
  "query": "mutation { refreshToken(input: { refresh_token: \"{{refresh_token}}\" }) { status message data { access_token refresh_token } } }"
}

HTTP 200
[Asserts]
jsonpath "$.data.refreshToken.status" == "success"
jsonpath "$.data.refreshToken.data.access_token" exists
jsonpath "$.data.refreshToken.data.refresh_token" exists