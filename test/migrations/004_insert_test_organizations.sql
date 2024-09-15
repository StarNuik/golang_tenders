INSERT INTO organization
    (name, description, type, id)
VALUES
    ('Org1', 'Description 1', 'LLC', '00000000-0000-0000-0000-000000000011'),
    ('Org2', 'Description 2', 'IE' , '00000000-0000-0000-0000-000000000012'),
    ('Org3', 'Description 3', 'JSC', '00000000-0000-0000-0000-000000000013'),
    ('Org4', 'Description 4', 'LLC', '00000000-0000-0000-0000-000000000014');

-- todo: make this more functional-like
CREATE PROCEDURE add_representative(org_name TEXT, empl_name TEXT) AS $$
    INSERT INTO organization_responsible
        (organization_id, user_id)
    VALUES (
        (SELECT id FROM organization WHERE name = org_name),
        (SELECT id FROM employee WHERE username = empl_name)
    );
$$ LANGUAGE SQL;

CALL add_representative('Org1', 'user01');
CALL add_representative('Org1', 'user02');
CALL add_representative('Org1', 'user03');

CALL add_representative('Org2', 'user04');
CALL add_representative('Org2', 'user05');
CALL add_representative('Org2', 'user06');

CALL add_representative('Org3', 'user07');
CALL add_representative('Org3', 'user08');
CALL add_representative('Org3', 'user09');

CALL add_representative('Org4', 'user10');
CALL add_representative('Org4', 'user11');
CALL add_representative('Org4', 'user12');
