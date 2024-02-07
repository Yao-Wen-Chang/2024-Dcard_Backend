#!/bin/bash

# Define variables for the API endpoint and query parameters
HOST="localhost:8080"
ENDPOINT="/api/v1/ad"
OFFSET="10"
LIMIT="3"
AGE="24"
GENDER="F"
COUNTRY="TW"
PLATFORM="ios"

# Construct the URL with query parameters
URL="http://${HOST}${ENDPOINT}?offset=${OFFSET}&limit=${LIMIT}&age=${AGE}&gender=${GENDER}&country=${COUNTRY}&platform=${PLATFORM}"

# Make the GET request using curl
curl -X GET -H "Content-Type: application/json" "${URL}"