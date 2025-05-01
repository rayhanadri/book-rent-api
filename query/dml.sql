/*
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    published_at DATE NOT NULL,
    isbn VARCHAR(255) NOT NULL UNIQUE,
    category VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    available BOOLEAN NOT NULL,
    price INT NOT NULL,
    description TEXT
);

*/

INSERT INTO books (title, author, publisher, published_at, isbn, category, stock, available, price, description) VALUES
('The Great Gatsby', 'F. Scott Fitzgerald', 'Scribner', '1925-04-10', '9780743273565', 'Fiction', 2, TRUE, 5000, 'A novel set in the Jazz Age.'),
('To Kill a Mockingbird', 'Harper Lee', 'J.B. Lippincott & Co.', '1960-07-11', '9780061120084', 'Fiction', 1, TRUE, 10000, 'A story of racial injustice in the Deep South.'),
('1984', 'George Orwell', 'Secker & Warburg', '1949-06-08', '9780451524935', 'Dystopian', 2, TRUE, 15000, 'A dystopian novel about totalitarianism.'),
('Pride and Prejudice', 'Jane Austen', 'T. Egerton', '1813-01-28', '9780141439518', 'Romance', 0, FALSE, 5000, 'A classic romance novel.'),
('The Catcher in the Rye', 'J.D. Salinger', 'Little, Brown and Company', '1951-07-16', '9780316769488', 'Fiction', 2, TRUE, 10000, 'A story about teenage rebellion.'),
('The Hobbit', 'J.R.R. Tolkien', 'George Allen & Unwin', '1937-09-21', '9780547928227', 'Fantasy', 2, TRUE, 15000, 'A fantasy adventure novel.'),
('Moby-Dick', 'Herman Melville', 'Harper & Brothers', '1851-10-18', '9781503280786', 'Adventure', 0, FALSE, 5000, 'A story about the quest for a giant whale.'),
('War and Peace', 'Leo Tolstoy', 'The Russian Messenger', '1869-01-01', '9780199232765', 'Historical', 1, TRUE, 10000, 'A novel about the Napoleonic Wars.'),
('The Alchemist', 'Paulo Coelho', 'HarperOne', '1988-05-01', '9780061122415', 'Adventure', 2, TRUE, 15000, 'A philosophical adventure story.'),
('The Road', 'Cormac McCarthy', 'Alfred A. Knopf', '2006-09-26', '9780307387899', 'Post-apocalyptic', 1, TRUE, 5000, 'A story of survival in a post-apocalyptic world.');
