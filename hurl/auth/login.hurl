# Auth Login — GraphQL Mutation
# Endpoint: POST /query (GraphQL)

POST http://localhost:5000/query
Content-Type: application/json

{
  "query": "mutation { loginUser(input: { email: \"john.doe@hellodota.com\", password: \"password123\" }) { status message data { access_token refresh_token } } }"
}

HTTP 200
[Asserts]
jsonpath "$.data.loginUser.status" == "success"
jsonpath "$.data.loginUser.data.access_token" exists
jsonpath "$.data.loginUser.data.refresh_token" exists
[Captures]
auth_token: jsonpath "$.data.loginUser.data.access_token"
refresh_token: jsonpath "$.data.loginUser.data.refresh_token"