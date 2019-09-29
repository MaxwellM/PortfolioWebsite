SELECT 'CREATE DATABASE portfoliowebsite_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'portfoliowebsite_db')\gexec

