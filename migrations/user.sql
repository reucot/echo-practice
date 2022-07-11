CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	first_name text NOT NULL,
	second_name text NOT NULL,
	date_of_birth date NOT NULL,
	income_per_year integer NOT NULL
);