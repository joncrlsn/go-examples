
INSERT INTO org(org_id, org_name) VALUES (1, 'Default');

-- password is 'supersecret'
INSERT INTO user (user_id, first_name, last_name, email, password, org_id) VALUES (1, 'Joe', 'Carter', 'joe@example.com', 'c44tqq+T5R57KwVuAJKfNiriIzmxR+uUcyJNuvAA7PvI84OE7xUpASLiGyMp4/vgYDVJu49u99ugxujEH4g7hw==', 1);
INSERT INTO user (user_id, first_name, last_name, email, password, org_id) VALUES (2, 'Jon', 'Carlson', 'jon@example.com', 'c44tqq+T5R57KwVuAJKfNiriIzmxR+uUcyJNuvAA7PvI84OE7xUpASLiGyMp4/vgYDVJu49u99ugxujEH4g7hw==', 1); 
INSERT INTO user (user_id, first_name, last_name, email, org_id) VALUES (3, 'Jane', 'Citizen', 'jane@example.com', 1);
