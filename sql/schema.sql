-- schema.sql
CREATE TABLE public.dictionary (
    kanji character varying NOT NULL,
    onyomi text,
    kunyomi text,
    imi text,
    itaiji text,
    bushu text
);
