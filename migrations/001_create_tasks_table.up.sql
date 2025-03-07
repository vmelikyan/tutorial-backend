CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    task TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    deleted BOOLEAN DEFAULT FALSE
);

-- Insert sample data
INSERT INTO tasks (task) VALUES
    ('Complete project documentation'),
    ('Review pull requests'),
    ('Setup monitoring system'),
    ('Update dependencies'),
    ('Write unit tests'), 
    ('Pray to the demo god offering a USB sacrifice'),
    ('Practice your surprised face for when something does not work');
