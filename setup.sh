#!/bin/bash

echo "Logging in and storing cookies..."

BASE_URL="http://localhost:8001"
COOKIE_FILE="cookies.txt"

curl -s -c "$COOKIE_FILE" -X POST \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "login_type=mobile&identifier=0000000000&password=Wowwhatagreatandsecurepassword123" \
  "$BASE_URL/login" > /dev/null

if grep -q "token" "$COOKIE_FILE"; then
  echo "Login successful. Token stored in $COOKIE_FILE"
else
  echo "Login failed or token cookie not found."
  exit 1
fi

echo "BASE_URL=$BASE_URL" > bench.env
echo "COOKIE_FILE=$COOKIE_FILE" >> bench.env
