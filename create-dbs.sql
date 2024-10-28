-- Create the database 'music_db'
CREATE DATABASE music_db;

-- Switch to 'music_db'
\c music_db;

-- Create 'album' table if it doesn't already exist
CREATE TABLE album (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    genre VARCHAR(100) NOT NULL,
    release_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create 'track' table if it doesn't already exist
CREATE TABLE track (
    id SERIAL PRIMARY KEY,
    album_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    duration INT NOT NULL,  -- Duration in seconds
    track_number INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (album_id) REFERENCES album (id) ON DELETE CASCADE
);

-- Insert sample data into 'album' table
INSERT INTO album (title, artist, genre, release_date) VALUES
('Thriller', 'Michael Jackson', 'Pop', '1982-11-30 00:00:00'),
('Back in Black', 'AC/DC', 'Rock', '1980-07-25 00:00:00'),
('The Dark Side of the Moon', 'Pink Floyd', 'Progressive Rock', '1973-03-01 00:00:00'),
('The Wall', 'Pink Floyd', 'Rock', '1979-11-30 00:00:00');

-- Insert sample data into 'track' table
INSERT INTO track (album_id, title, duration, track_number) VALUES
(1, 'Wanna Be Startin', 362, 1),
(1, 'Thriller', 357, 2),
(1, 'Beat It', 258, 3),
(2, 'Hells Bells', 312, 1),
(2, 'Back in Black', 255, 2),
(3, 'Speak to Me', 90, 1),
(3, 'Time', 413, 2),
(4, 'Another Brick in the Wall, Pt. 2', 240, 3);
