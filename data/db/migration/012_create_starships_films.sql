CREATE TABLE starships_films(
    starship_id BIGINT REFERENCES starships,
    film_id BIGINT REFERENCES films,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_starships_films_on_starships_id" ON starships_films USING btree(
    starship_id
);
CREATE UNIQUE INDEX "index_starships_films_on_film_id" ON starships_films USING btree(
    film_id
);

---- create above / drop below ----

DROP TABLE starships_films;