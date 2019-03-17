DATA_DIR=./data

run:
	go run main.go

start_db:
	gcloud beta emulators datastore start --data-dir=${DATA_DIR} --no-store-on-disk