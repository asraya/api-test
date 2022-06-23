migrate db
DBEVENT=migration go run main.go

seed
DBEVENT=seed go run main.go

rollback
DBEVENT=seed go run main.go

News and Topic management
CRUD ON NEWS AND Tags
One news can contains multiple tags e.g. "Safe investment" might contains tags
"investment", "mutual fund", etc
One news can contains multiple tags e.g. "Safe investment" might contains tags
"investment", "mutual fund", etc
Enable filter by news status ("draft", "deleted", "publish")
Enable filter news by its topics