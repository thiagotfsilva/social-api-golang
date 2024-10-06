
USE devbook;

INSERT INTO users (name, nick, email, password) VALUES
('Alice Santos', 'alice123', 'alice.santos@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('Bruno Ferreira', 'bruno456', 'bruno.ferreira@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('Carla Mendes', 'carla789', 'carla.mendes@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('Diego Silva', 'diego001', 'diego.silva@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('Elena Costa', 'elena999', 'elena.costa@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('Fabio Pereira', 'fabio007', 'fabio.pereira@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('Gisele Rocha', 'gisele654', 'gisele.rocha@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('Henrique Almeida', 'henrique111', 'henrique.almeida@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('Isabela Nunes', 'isabela222', 'isabela.nunes@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu'),
('João Souza', 'joao333', 'joao.souza@example.com', '$2a$10$cAxwMhVorQ/BGH/pdkyrcuG5rlrY8PsHcJiszL.RswtXGFGdyuZTu');

INSERT INTO followers (user_id, follower_id) VALUES
(1, 2),
(1, 3),
(2, 1),
(2, 4),
(3, 1),
(3, 4),
(4, 2),
(4, 3),
(5, 1),
(5, 2),
(6, 3),
(6, 5),
(7, 4),
(7, 6),
(8, 1),
(8, 7),
(9, 8),
(9, 3),
(10, 2),
(10, 6);

SELECT * FROM users;
SELECT * FROM followers;

INSERT INTO publications(title, content, author_id)
VALUES
("Publicação do usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do usuário 2", "Essa é a publicação do usuário 1! Oba!", 2),
("Publicação do usuário 13", "Essa é a publicação do usuário 1! Oba!", 3);