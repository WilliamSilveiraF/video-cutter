INSERT INTO use_terms (version, description)
SELECT '1.0', 'Default Use Terms Template'
WHERE NOT EXISTS (SELECT 1 FROM use_terms);
