-- name: CreateProfile :one
INSERT INTO profiles
(user_id, first_name, last_name, dob, 
address_line1, address_line2, city, state, country, postal_code, 
primary_phone, secondary_phone, 
primary_email, secondary_email, created_at, updated_at) 
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
RETURNING *;
-- name: GetProfile :one
SELECT * FROM profiles WHERE user_id = $1;

-- name: ListProfiles :many
SELECT * FROM profiles LIMIT $1 OFFSET $2;


-- name: UpdateProfile :exec
UPDATE profiles
SET
first_name= $1,
last_name= $2,
dob= $3,
address_line1= $4,
address_line2= $5,
city= $6,
state= $7,
country= $8,
postal_code= $9,
primary_phone= $10,
secondary_phone= $11,
primary_email= $12,
secondary_email= $13,
updated_at = $14
WHERE user_id = $15;