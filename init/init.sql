-- Table: public.song

DROP TABLE IF EXISTS public.song;
CREATE TABLE IF NOT EXISTS public.song
(
    author text COLLATE pg_catalog."default",
    title text COLLATE pg_catalog."default",
    link text COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.song
    OWNER to postgres;

INSERT INTO song (author, title, link) VALUES
('Lola Young', 'Messy', ''),
('Rosé & Bruno Mars', 'APT.', ''),
('Gracie Abrams', 'Thats So True', ''),
('Lady Gaga & Bruno Mars', 'Die With a Smile', ''),
('Chrystal', 'The Days', ''),
('Billie Eilish', 'Birds of a Feather', ''),
('Chappell Roan', 'Good Luck, Babe!', ''),
('Gigi Perez', 'Sailor Song', ''),
('Benson Boone', 'Beautiful Things', ''),
('Kendrick Lamar & SZA', 'Luther', '');