#!/bin/bash

HOST="localhost:8080"
API_ENDPOINT="http://${HOST}/api/v1/ad"

curl -X POST -H "Content-Type: application/json" \
  "${API_ENDPOINT}" \
  --data '{
     "title": "AD 55",
     "startAt": "2023-12-10T03:00:00.000Z",
     "endAt": "2023-12-31T16:00:00.000Z",
     "conditions": {
        "ageStart": 20,
        "ageEnd": 30,
        "country": ["TW", "JP"],
        "platform": ["android", "ios"]
     }  
  }'
