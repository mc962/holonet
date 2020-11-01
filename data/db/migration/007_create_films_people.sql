CREATE TABLE films_people(
    film_id BIGINT REFERENCES films,
    person_id BIGINT REFERENCES people,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_films_people_on_film_id" ON films_people USING btree(
    film_id
);
CREATE UNIQUE INDEX "index_films_people_on_person_id" ON films_people USING btree(
    person_id
);

---- create above / drop below ----

DROP TABLE films_people;