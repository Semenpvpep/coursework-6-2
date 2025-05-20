#!/bin/bash

# Base API URL
BASE_URL="http://localhost:8000/api/v1"
CONF_URL="${BASE_URL}/appointments"
AUTH_URL="${BASE_URL}/auth"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Starting API tests..."

# Test 1: Register user
echo -e "\n${GREEN}Test 1: Register user${NC}"
REGISTER_RESPONSE=$(curl -s -X POST "${AUTH_URL}/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpass"
  }')
echo "Response: $REGISTER_RESPONSE"

# Test 2: Try to create appointment without auth
echo -e "\n${GREEN}Test 2: Create appointment without auth${NC}"
NO_AUTH_RESPONSE=$(curl -s -X POST "${CONF_URL}/" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Python Web Development Workshop",
    "description": "Learn about modern Python web development",
    "start_time": "2024-03-20T10:00:00Z",
    "end_time": "2024-03-20T18:00:00Z",
    "max_participants": 100,
    "registration_deadline": "2024-03-19T23:59:59Z",
    "timezone": "UTC",
    "tags": ["python", "web", "workshop"]
  }')
echo "Response: $NO_AUTH_RESPONSE"

# Test 3: Create appointment with auth
echo -e "\n${GREEN}Test 3: Create appointment with auth${NC}"
CREATE_RESPONSE=$(curl -s -X POST "${CONF_URL}/" \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic $(echo -n 'testuser:testpass' | base64)" \
  -d '{
    "title": "Python Web Development Workshop",
    "description": "Learn about modern Python web development",
    "start_time": "2024-03-20T10:00:00Z",
    "end_time": "2024-03-20T18:00:00Z",
    "status": "scheduled",
    "meeting_link": "https://zoom.us/j/123456789",
    "organizer": "testuser",
    "max_participants": 100,
    "registration_deadline": "2024-03-19T23:59:59Z",
    "timezone": "UTC",
    "tags": ["python", "web", "workshop"]
  }')
echo "Response: $CREATE_RESPONSE"

# Extract appointment ID from response
appointment_ID=$(echo $CREATE_RESPONSE | jq -r '.appointment_id')

# Test 4: Get created appointment
echo -e "\n${GREEN}Test 4: Get appointment${NC}"
GET_RESPONSE=$(curl -s -X GET "${CONF_URL}/${appointment_ID}" \
  -H "Authorization: Basic $(echo -n 'testuser:testpass' | base64)")
echo "Response: $GET_RESPONSE"

# Test 5: Update appointment
echo -e "\n${GREEN}Test 5: Update appointment${NC}"
UPDATE_RESPONSE=$(curl -s -X PUT "${CONF_URL}/${appointment_ID}" \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic $(echo -n 'testuser:testpass' | base64)" \
  -d '{
    "title": "Advanced Python Web Development Workshop",
    "description": "Deep dive into modern Python web development",
    "start_time": "2024-03-21T10:00:00Z",
    "end_time": "2024-03-21T18:00:00Z",
    "status": "scheduled",
    "meeting_link": "https://zoom.us/j/987654321",
    "organizer": "testuser",
    "max_participants": 150,
    "registration_deadline": "2024-03-20T23:59:59Z",
    "timezone": "UTC",
    "tags": ["python", "web", "advanced", "workshop"]
  }')
echo "Response: $UPDATE_RESPONSE"

# Test 6: Get all appointments
echo -e "\n${GREEN}Test 6: List all appointments${NC}"
LIST_RESPONSE=$(curl -s -X GET "${CONF_URL}/" \
  -H "Authorization: Basic $(echo -n 'testuser:testpass' | base64)")
echo "Response: $LIST_RESPONSE"

# Test 7: Try to get non-existent appointment
echo -e "\n${GREEN}Test 7: Get non-existent appointment${NC}"
NOT_FOUND_RESPONSE=$(curl -s -X GET "${CONF_URL}/non-existent-id" \
  -H "Authorization: Basic $(echo -n 'testuser:testpass' | base64)")
echo "Response: $NOT_FOUND_RESPONSE"

# Test 8: Delete appointment
echo -e "\n${GREEN}Test 8: Delete appointment${NC}"
DELETE_RESPONSE=$(curl -s -X DELETE "${CONF_URL}/${appointment_ID}" \
  -H "Authorization: Basic $(echo -n 'testuser:testpass' | base64)")
echo "Response: $DELETE_RESPONSE"

# Test 9: Verify deletion
echo -e "\n${GREEN}Test 9: Verify deletion${NC}"
GET_DELETED_RESPONSE=$(curl -s -X GET "${CONF_URL}/${appointment_ID}" \
  -H "Authorization: Basic $(echo -n 'testuser:testpass' | base64)")
echo "Response: $GET_DELETED_RESPONSE"

echo -e "\n${GREEN}Testing completed${NC}" 