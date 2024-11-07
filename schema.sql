DROP TABLE IF EXISTS authors, posts;

CREATE TABLE IF NOT EXISTS authors (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY, 
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now())
);

INSERT INTO authors (name) VALUES ('Макаренко Владимир'), ('Гудимова Светлана Алексеевна'), ('Гусева Наталья Игоревна');
INSERT INTO posts (author_id, title, content) VALUES (2, 'Музыка в контексте культуры', 'Рассматривается музыкальная эстетика Европы, Индии, Китая, Японии'), (3, 'Численное моделирование ламинарных и турбулентных течений в следе за телом', 'Изложены результаты численного моделирования течений в ближнем следе за плоскими и осесимметричными телами'), (1, 'Экономическая аксиология: опыт исследования экономических культур', 'Главную задачу данной статьи автор видит в том, чтобы существенно расширить область обсуждения взаимосвязи экономической и культурной сфер и вписать российские дискуссии в общий тренд экономической компаративистики. ');