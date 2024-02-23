CREATE TABLE usr (
    id bigserial PRIMARY KEY,
    usr_name varchar(20) NOT NULL,
    pass varchar(20) NOT NULL
);

CREATE TABLE news (
    id bigserial PRIMARY KEY,
    title varchar(255) NOT NULL,
    cnt text NOT NULL 
);

CREATE TABLE news_categories (
    id bigserial PRIMARY KEY,
    news_id bigint REFERENCES news(id),
    category_id bigint NOT NULL,

    UNIQUE(news_id, category_id)
);
