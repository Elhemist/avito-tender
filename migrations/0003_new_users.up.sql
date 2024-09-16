INSERT INTO employee (username, first_name, last_name)
VALUES 
    ('ivanov', 'Иван', 'Иванов'),
    ('petrov', 'Петр', 'Петров'),
    ('saveliy', 'Савелий', 'Савельев'),
    ('nikolaev', 'Николай', 'Николаев');

INSERT INTO organization (name, description, type)
VALUES 
    ('Рога и Копыта', 'Рога и Копыта', 'LLC'),
    ('Сибирские Леса', 'Лиса', 'JSC'),
    ('Торговый Дом "Голандский сыр"', 'Я не безуумец.', 'IE');

INSERT INTO organization_responsible (organization_id, user_id)
VALUES 
    ((SELECT id FROM organization WHERE name = 'Рога и Копыта'), (SELECT id FROM employee WHERE username = 'ivanov')),
    ((SELECT id FROM organization WHERE name = 'Сибирские Леса'), (SELECT id FROM employee WHERE username = 'petrov')),
    ((SELECT id FROM organization WHERE name = 'Торговый Дом "Голандский сыр"'), (SELECT id FROM employee WHERE username = 'sidorov')),
    ((SELECT id FROM organization WHERE name = 'Рога и Копыта'), (SELECT id FROM employee WHERE username = 'nikolaev'));
