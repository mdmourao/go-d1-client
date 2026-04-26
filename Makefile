run:
	go run .

dev:
	npx wrangler dev

deploy:
	npx wrangler deploy

secret:
	npx wrangler secret put D1_DSN
	npx wrangler secret put CF_ACCESS_CLIENT_ID
	npx wrangler secret put CF_ACCESS_CLIENT_SECRET

tail:
	npx wrangler tail