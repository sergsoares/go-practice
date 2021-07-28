CREATE TABLE todo(
    id int GENERATED ALWAYS AS IDENTITY,
    description text NOT NULL,
    created_date timestamp NOT NULL,
    completed_date timestamp NULL
);