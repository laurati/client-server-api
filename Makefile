#sqlite3 cotacoes.db

#CREATE TABLE cotacoes (
#    id INTEGER PRIMARY KEY AUTOINCREMENT,
#    valor TEXT
#);

run:
	go run main.go

client-run:
	go run client/main.go