ALTER TABLE transactions 
ADD COLUMN IF NOT EXISTS tenor_months int;

UPDATE transactions SET tenor_months = 2;