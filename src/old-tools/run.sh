#!/bin/bash

set -e
set -x

export PGUSER=postgres
export PGPASSWORD=Maxwellmorin1
export PGPORT=5432
export PGDATABASE=portfoliowebsite_db
export GOOGLE_APPLICATION_CREDENTIALS="googleKey.json"

cd web
./buildwebsite.py
cd ..

go install
PortfolioWebsite
