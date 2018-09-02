.PHONY: api deploy-infra deploy-api start-api token-observer
.SILENT: api deploy-infra deploy-api start-api token-observer

api:
	rm -rf bin/* && cd api && env GOOS=linux go build -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o ../bin/networth-api .

token-observer:
	rm -rf bin/* && cd api/token_observer && env GOOS=linux go build -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o ../../bin/networth-token-observer .

deploy-infra:
	aws cloudformation deploy --template-file cloud/aws.infra.yml --stack-name networth --capabilities CAPABILITY_IAM --region us-east-1

deploy-api:
	make api
	aws cloudformation package --template-file cloud/aws.rest.api.yml --s3-bucket lambda.networth.app --output-template-file /tmp/aws.rest.api.yml --s3-prefix networth-api
	aws cloudformation deploy --template-file /tmp/aws.rest.api.yml --stack-name networth-api --capabilities CAPABILITY_IAM --region us-east-1

deploy-token-observer:
	make token-observer
	aws cloudformation package --template-file cloud/aws.token.observer.yml --s3-bucket lambda.networth.app --output-template-file /tmp/aws.token.observer.yml --s3-prefix networth-token-observer
	aws cloudformation deploy --template-file /tmp/aws.token.observer.yml --stack-name networth-token-observer --capabilities CAPABILITY_IAM --region us-east-1

start-api:
	cd api && gin --appPort 8000
	# make api
	# cd api && sam local start-api --env-vars .env.json

start-web:
	cd web && npm run start

start-token-observer:
	make token-observer
	sam local generate-event dynamodb | sam local invoke "NetWorthTokenObserverFunction" --template cloud/aws.token.observer.yml

token:
	# reset pass
	# aws cognito-idp admin-respond-to-auth-challenge --user-pool-id us-east-1_5cJz62UiG --challenge-name NEW_PASSWORD_REQUIRED --client-id 2tam11a22g38in2vqcd5kge3cu --challenge-responses USERNAME=demo@networth.app,NEW_PASSWORD=Testing!!1234. --session "kvdceEkVdbV6hI-_8Gpk4--kgIIu89Fic5GYW-F5BiL94WvWrgVB4_ZDjOUWGmQRSpNdC7qxjjgfThHaRAxspPIYTnOUql9RIiSVBtIyzc3sMU8gtQYjyjkJs04zy4gZROp8FA6GUe41Se_J9I5s_J9zWa7g36OWIpFRivJ2KGmVVpxAZnePYA3NzG01sTOBDd2XXWP_j4wb-c27FzCBlUSba3dlBZORysRhtEg-mtoGeRxu3KMc5RzSJELw-56fqcWZOu22aFV0qbH4sixWfhdZQmSkdAce50CYQt5UQCN6ZF-IgLU2Itghb2LaYO_rcICo5WKLnhF61iQ9Vqn9d9_VA7rljVsGuCofXNAwnqXhKfSMxbJT8hHVcYuX90XUnT--CvcU9tY51F_pa_KYsulrKRWvrkizbVRcIvnvIo8tInI6YQE5yWOzHN6UAB3rb9RK_6bqAcLO6yTAgYjJ-sLBqsozN7DLijYnIJ52-Ch74mSJATbfaBHAu22mT8hsMvhxS0NvbpY_NjJl5sJAAoaW4mea4p4i0LDfI3k8hKkp_JUav_R6yz8LZ01jW7Cco10qHbHO6RYoiPD6V8xPF-CfZaOiWQw9LceoJA_RbOiFKkgTGuoSR1p6YdtTE0G_5PNqsO4oF2bt6Yo6uAnTYb5fZYYVHDfD-kC5uN3lQ_tIp6ojEI70xhm1XfYh0FwpAYuHyje9ucfp2vlNFI6sXyyJBFD7rInp6dFwTia0KgUfxNvpFJrCa9ls3VS92qh5Hbv6lo_2arf0zbemq_oS2xP2-Rflyscbi1qPRjUdCu4vkT9zYg6tnAxMGWmEzoFwZ2Ju6-4Kcv3dkyImAw-DqQ"
	aws cognito-idp initiate-auth --client-id 2tam11a22g38in2vqcd5kge3cu --auth-flow USER_PASSWORD_AUTH --auth-parameters USERNAME=demo@networth.app,PASSWORD=Testing!!1234. | jq -r .AuthenticationResult.IdToken

