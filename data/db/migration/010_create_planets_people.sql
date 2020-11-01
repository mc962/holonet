CREATE TABLE planets_people(
    planet_id BIGINT REFERENCES planets,
    person_id BIGINT REFERENCES people,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_planets_people_on_planets_id" ON planets_people USING btree(
    planet_id
);
CREATE UNIQUE INDEX "index_planets_people_on_person_id" ON planets_people USING btree(
    person_id
);

---- create above / drop below ----

DROP TABLE planets_people;