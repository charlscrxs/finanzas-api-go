CREATE DATABASE finanzas_personales;


CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    contrasena TEXT NOT NULL  -- Contraseña cifrada
);


CREATE TABLE categorias (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE,    -- Nombre de la categoría (ej. "Alimentación")
    descripcion TEXT                        -- Descripción opcional de la categoría
);


CREATE TABLE movimientos (
    id SERIAL PRIMARY KEY,                  -- Identificador único
    tipo VARCHAR(20) NOT NULL,              -- 'ingreso' o 'gasto'
    monto DECIMAL(10, 2) NOT NULL,          -- Monto del movimiento
    categoria_id INT REFERENCES categorias(id) ON DELETE SET NULL, -- Relación con la categoría
    descripcion TEXT,                       -- Descripción opcional (por ejemplo, "compra en supermercado")
    fecha TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Fecha del movimiento
    usuario_id INT REFERENCES usuarios(id) ON DELETE CASCADE -- Relación con el usuario (si tienes autenticación)
);



-- Insertar usuarios
INSERT INTO usuarios (nombre, email, contrasena) VALUES
('Juan Pérez', 'juan@example.com', 'hashedpassword1'),
('María López', 'maria@example.com', 'hashedpassword2'),
('Carlos García', 'carlos@example.com', 'hashedpassword3'),
('Ana Torres', 'ana@example.com', 'hashedpassword4'),
('Luis Fernández', 'luis@example.com', 'hashedpassword5');

-- Insertar categorías
INSERT INTO categorias (nombre, descripcion) VALUES
('Alimentación', 'Gastos en comida y supermercado'),
('Transporte', 'Gastos en gasolina, pasajes y mantenimiento de vehículo'),
('Entretenimiento', 'Gastos en ocio, cine, conciertos, etc.'),
('Salud', 'Gastos en medicina, consultas médicas y seguros'),
('Educación', 'Gastos en cursos, libros y colegiaturas');

-- Insertar movimientos
INSERT INTO movimientos (tipo, monto, categoria_id, descripcion, usuario_id) VALUES
('gasto', 50.00, 1, 'Compra en supermercado', 1),
('gasto', 15.00, 2, 'Pasaje de autobús', 2),
('ingreso', 2000.00, NULL, 'Salario mensual', 3),
('gasto', 120.00, 4, 'Consulta médica', 4),
('gasto', 80.00, 3, 'Salida al cine', 5);
