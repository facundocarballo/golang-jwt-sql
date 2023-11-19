package db

const INSERT_USER_STATEMENT = "INSERT INTO User (name, email, password) VALUES (?, ?, ?)"
const INSERT_TASK_STATEMENT = "INSERT INTO Task (name, description, owner) VALUES (?, ?, ?)"
