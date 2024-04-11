start:
	export DB_HOST=aws-0-ap-southeast-1.pooler.supabase.com; \
    export DB_PORT=5432; \
    export DB_USER=postgres.obbzmgleelshmaficbrh; \
    export DB_PASSWORD=Speedattack2107; \
    export DB_NAME=postgres; \
    export CLD_NAME=dtitvei0p; \
    export CLD_API_KEY=865515316225168; \
    export CLD_API_SECRET=xJaTdJi4pW3JaAjp2Q4sPIuqG-U; \
    export MONGO_HOST=localhost; \
    export MONGO_PORT=27017; \
	go run cmd/main.go

gen-swag:
	swag init -g ./cmd/main.go -o cmd/docs

