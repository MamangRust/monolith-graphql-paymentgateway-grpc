# Merchant Document Service GraphQL Tests
# Endpoint: POST /query (GraphQL)

### Create a new merchant document
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { createMerchantDocument(input: { merchant_id: 1, document_type: \"business_license\", document_url: \"https://example.com/documents/business_license.pdf\" }) { status message data { document_id merchant_id document_type document_url status note uploaded_at updated_at } } }"
}

HTTP 200

### Find all merchant documents (paginated)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllMerchantDocuments(input: { page: 1, page_size: 10, search: \"\" }) { status message data { document_id merchant_id document_type document_url status note uploaded_at updated_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Find merchant document by ID
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findMerchantDocumentById(input: { document_id: 1 }) { status message data { document_id merchant_id document_type document_url status note uploaded_at updated_at } } }"
}

HTTP 200

### Update merchant document
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateMerchantDocument(input: { document_id: 1, merchant_id: 1, document_type: \"business_license\", document_url: \"https://example.com/documents/updated_license.pdf\", note: \"Updated document\", status: \"approved\" }) { status message data { document_id merchant_id document_type document_url status note uploaded_at updated_at } } }"
}

HTTP 200

### Update merchant document status
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { updateMerchantDocumentStatus(input: { document_id: 1, merchant_id: 1, note: \"Verified\", status: \"approved\" }) { status message data { document_id merchant_id document_type document_url status note uploaded_at updated_at } } }"
}

HTTP 200

### Find active merchant documents
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "query { findAllActiveMerchantDocuments(input: { page: 1, page_size: 10, search: \"\" }) { status message data { document_id merchant_id document_type document_url status note uploaded_at updated_at deleted_at } pagination { current_page page_size total_pages total_records } } }"
}

HTTP 200

### Trash merchant document (soft delete)
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { trashedMerchantDocument(input: { document_id: 1 }) { status message data { document_id merchant_id document_type document_url status note uploaded_at updated_at deleted_at } } }"
}

HTTP 200

### Restore merchant document
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { restoreMerchantDocument(input: { document_id: 1 }) { status message data { document_id merchant_id document_type document_url status note uploaded_at updated_at deleted_at } } }"
}

HTTP 200

### Delete merchant document permanently
POST http://localhost:5000/query
Content-Type: application/json
Authorization: Bearer {{auth_token}}

{
  "query": "mutation { deleteMerchantDocumentPermanent(input: { document_id: 1 }) { status message } }"
}

HTTP 200