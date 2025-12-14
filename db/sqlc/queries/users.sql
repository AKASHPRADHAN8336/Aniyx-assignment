
-- create user
INSERT INTO users (name, dob)
VALUES (?, ?);

-- GetUserByID 
SELECT id, name, DATE_FORMAT(dob, '%Y-%m-%d') as dob FROM users WHERE id = ?;

-- ListUsers 
SELECT id, name, DATE_FORMAT(dob, '%Y-%m-%d') as dob FROM users ORDER BY id;

-- UpdateUser 
UPDATE users
SET name = ?, dob = ?
WHERE id = ?;

-- DeleteUser
DELETE FROM users WHERE id = ?;