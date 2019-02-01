#!/bin/bash

set -e

cd web
./buildwebsite.py
cd ..

go install
PortfolioWebsite
