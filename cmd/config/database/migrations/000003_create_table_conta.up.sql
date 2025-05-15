CREATE TABLE IF NOT EXISTS conta (
    id SERIAL PRIMARY KEY,
    usuario_id INT REFERENCES usuario(id),
    banco_id INT REFERENCES banco(id),
    agencia VARCHAR(20) NOT NULL,
    numero VARCHAR(20) NOT NULL UNIQUE
);
