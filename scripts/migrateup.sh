#!/bin/bash

if [ -f .env ]; then
    source .env
fi

goose -dir ./sql/schema turso $DATABASE_URL up
