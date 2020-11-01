CREATE TABLE planets_films(
    planet_id BIGINT REFERENCES planets,
    film_id BIGINT REFERENCES films,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_planets_films_on_planets_id" ON planets_films USING btree(
    planet_id
);
CREATE UNIQUE INDEX "index_planets_films_on_film_id" ON planets_films USING btree(
    film_id
);

---- create above / drop below ----

DROP TABLE planets_films;