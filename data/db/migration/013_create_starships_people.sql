CREATE TABLE starships_people(
    starship_id BIGINT REFERENCES starships,
    person_id BIGINT REFERENCES people,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_starships_people_on_starships_id" ON starships_people USING btree(
    starship_id
);
CREATE UNIQUE INDEX "index_starships_people_on_person_id" ON starships_people USING btree(
    person_id
);

---- create above / drop below ----

DROP TABLE starships_people;