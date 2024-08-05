create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('Bruno', 'Homem Top'),
('Polna', 'Homem do Fuzil');