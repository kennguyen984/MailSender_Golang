
# Config SMTP to send mail on file /.env

Create file .env on src/ folder with variables on below

    EMAIL_TYPE="SMTP" # SMTP or API
    EMAIL_API_KEY="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"
    EMAIL_HOST="smtp-relay.sendinblue.com"
    EMAIL_PORT=587
    EMAIL_AUTH_USER="ntnpro@gmail.com"
    EMAIL_AUTH_PASSWORD="DYcS7U2GktBxQvVg"


# How to setup and run by command line

1. Run GO initital command line to download third modules
    go get
    go test

2. Build app and run send email by command line
    go build -o main .
    ./main ./SampleData/email_template.json ./SampleData/customers.csv ./SampleData/ ./SampleData/errors.csv


# How to setup and run by Docker

1. Build product to a container
    docker build -t sendmail .

2. Run send email by command line
    docker run sendmail ./SampleData/email_template.json ./SampleData/customers.csv ./SampleData/ ./errors.csv
