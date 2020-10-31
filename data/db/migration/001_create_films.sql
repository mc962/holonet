CREATE TABLE films(
    id BIGSERIAL PRIMARY KEY,
    title CHARACTER VARYING NOT NULL,
    -- based on title
    slug CHARACTER VARYING NOT NULL,
    episode_id INTEGER NOT NULL,
    opening_crawl TEXT NOT NULL,
    director CHARACTER VARYING NOT NULL,
    producer CHARACTER VARYING NOT NULL,
    release_date DATE,
    created TIMESTAMP WITH TIME ZONE NOT NULL,
    edited TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX "index_films_on_title" ON films USING btree(
    title
);
CREATE UNIQUE INDEX "index_films_on_slug" ON films USING btree(
    slug
);

---- create above / drop below ----

DROP TABLE films;
