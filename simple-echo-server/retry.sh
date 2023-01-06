#!/bin/bash

echo "Ejecutando curl con retry"

http_response=""
http_response=$(curl --retry-delay 15 --retry 1 --retry-all-errors -m 6 -o /dev/null -s -I -w "%{http_code}\n" -X GET "http://localhost:8086/x" -H 'X-application-id:ftc-slots-ms-scheduled-tasks' -H "API_KEY:test")

echo "Request to modify Rts capacities finished with response status code: ${http_response}"


