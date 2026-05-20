#!/bin/bash

# Hurl GraphQL API Test Runner
# This script runs all Hurl test files for the payment gateway GraphQL API

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Base URL for the GraphQL API
BASE_URL="http://localhost:5000"
GRAPHQL_ENDPOINT="${BASE_URL}/query"

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if hurl is installed
check_hurl() {
    if ! command -v hurl &> /dev/null; then
        print_error "Hurl is not installed. Please install it first:"
        echo "Visit: https://hurl.dev/docs/installation.html"
        exit 1
    fi
}

# Check if GraphQL API is running
check_api() {
    print_status "Checking if GraphQL API Gateway is running on $BASE_URL..."
    if curl -s -X POST "$GRAPHQL_ENDPOINT" \
        -H "Content-Type: application/json" \
        -d '{"query":"{ __typename }"}' | grep -q "Query"; then
        print_success "GraphQL API Gateway is running!"
    else
        print_error "GraphQL API Gateway is not responding on $GRAPHQL_ENDPOINT"
        print_warning "Please start the API Gateway first"
        exit 1
    fi
}

# Run a single test file
run_test_file() {
    local file=$1
    print_status "Running tests in $file..."
    
    if hurl "$file" 2>/dev/null; then
        print_success "All tests in $file passed!"
        return 0
    else
        print_error "Some tests in $file failed!"
        return 1
    fi
}

# Main execution
main() {
    echo "========================================"
    echo "  Payment Gateway GraphQL Test Runner"
    echo "========================================"
    echo

    # Check prerequisites
    check_hurl
    check_api

    echo
    print_status "Starting GraphQL API tests..."
    echo

    # Get all .hurl files (root level)
    test_files=( *.hurl )
    failed_tests=()
    passed_tests=()

    # Run each test file
    for file in "${test_files[@]}"; do
        if [[ -f "$file" ]]; then
            echo "----------------------------------------"
            if run_test_file "$file"; then
                passed_tests+=("$file")
            else
                failed_tests+=("$file")
            fi
            echo
        fi
    done

    # Run auth subdirectory tests
    if [[ -d "auth" ]]; then
        for file in auth/*.hurl; do
            if [[ -f "$file" ]]; then
                echo "----------------------------------------"
                if run_test_file "$file"; then
                    passed_tests+=("$file")
                else
                    failed_tests+=("$file")
                fi
                echo
            fi
        done
    fi

    # Summary
    echo "========================================"
    echo "  Test Summary"
    echo "========================================"
    echo

    if [[ ${#passed_tests[@]} -gt 0 ]]; then
        print_success "Passed tests (${#passed_tests[@]}):"
        for file in "${passed_tests[@]}"; do
            echo "  ✓ $file"
        done
        echo
    fi

    if [[ ${#failed_tests[@]} -gt 0 ]]; then
        print_error "Failed tests (${#failed_tests[@]}):"
        for file in "${failed_tests[@]}"; do
            echo "  ✗ $file"
        done
        echo
    fi

    # Exit with appropriate code
    if [[ ${#failed_tests[@]} -eq 0 ]]; then
        print_success "All tests passed! 🎉"
        exit 0
    else
        print_error "Some tests failed. Please check the output above."
        exit 1
    fi
}

# Handle command line arguments
case "${1:-}" in
    --help|-h)
        echo "Usage: $0 [options]"
        echo
        echo "Options:"
        echo "  --help, -h     Show this help message"
        echo "  --check         Only check prerequisites (hurl installation and API status)"
        echo
        echo "All tests send GraphQL queries/mutations via POST to $GRAPHQL_ENDPOINT"
        echo
        echo "Examples:"
        echo "  $0              # Run all tests"
        echo "  $0 --check      # Check prerequisites only"
        exit 0
        ;;
    --check)
        check_hurl
        check_api
        print_success "All prerequisites met!"
        exit 0
        ;;
    "")
        main
        ;;
    *)
        print_error "Unknown option: $1"
        echo "Use --help for usage information"
        exit 1
        ;;
esac