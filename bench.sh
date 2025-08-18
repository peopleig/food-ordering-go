#!/bin/bash

source bench.env

echo "--------------------------------------------------"
echo "Running GET /menu benchmark with cookies..."
echo "--------------------------------------------------"
ab -n 100000 -c 1000 \
  -C "jwt_token=$(grep token $COOKIE_FILE | awk '{print $NF}')" \
  "$BASE_URL/api/menu"

echo "--------------------------------------------------"
echo "Running POST /menu benchmark with cookies..."
echo "--------------------------------------------------"

echo '{
  "cart": [
    {
      "quantity": 2,
      "itemId": 1
    },
    {
      "quantity": 1,
      "itemId": 2
    }
  ],
  "special_instructions": "No onion no garlic",
  "order_type": "takeaway",
  "table_number": "7"
}' > payload.json

ab -n 100000 -c 100 \
  -T application/json \
  -p payload.json \
  -C "jwt_token=$(grep token $COOKIE_FILE | awk '{print $NF}')" \
  "$BASE_URL/menu"

rm payload.json
