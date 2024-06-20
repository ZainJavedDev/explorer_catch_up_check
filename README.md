# Matches Explorer Monitoring Microservice
This application serves as a monitoring microservice for the Matches Explorer application. It runs as an hourly cron job and performs two main functions:

## Match Logging
The application retrieves and logs the last match played and the time elapsed since it was played. This helps in determining the amount of latest data missing from the Matches Explorer application. Given the large amount of data generated by the Dota 2 API, even constant operation of the explorer can result in lagging behind the played matches.

## Storage Monitoring
The second function of the application is to monitor the available storage within the volume. The large amount of data can consume significant storage, necessitating occasional clean-up or resizing. This application enables monitoring of the available storage to take appropriate action as needed.


Please make sure to update your .env file with the necessary environment variables before running the application.

You would need the following values:

- `MYSQL_DB`: The name of the database.
- `MYSQL_USER`: The user for the database.
- `MYSQL_PASSWORD`: The password for the database.
- `WEBHOOK_ID`: The ID of the Discord webhook.
- `WEBHOOK_TOKEN`: The Token of the Discord webhook.