# Auth Register — GraphQL Mutation
# Endpoint: POST /query (GraphQL)

POST http://localhost:5000/query
Content-Type: application/json

{
  "query": "mutation { registerUser(input: { firstname: \"John\", lastname: \"Doe\", email: \"john.doe@hellodota.com\", password: \"password123\", confirm_password: \"password123\" }) { status message data { id firstname lastname email created_at updated_at } } }"
}

HTTP 200
[Asserts]
jsonpath "$.data.registerUser.status" == "success"