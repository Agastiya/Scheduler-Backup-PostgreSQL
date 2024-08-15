# Backup Database (postgreSQL) With Scheduler and Notifiers

## About
In this project I create a database dump file (postgreSQL) as a backup with a scheduler and send it to one or more platforms as Notifiers.

This is just a fun project. **Enjoy**.

## Getting Started
We will prepare this project, according to what notifier you want.

## Setup Notifiers
* Discord
* Telegram `Comming Soon :D`

### Discord
* We need to create a webhook from your channel.
* The steps for create a webhook are: 
    - Open your **Server Settings** and head into the Integrations tab:
    - Click the **Create Webhook** button to create a new webhook!
    - Choose what channel the Webhook posts to.
    - Set your webhook name.
    - Copy the webhook URL.

## Installation
Several stages that need to be prepared at this stage are as follows :

* Clone this repository

        git clone https://github.com/Agastiya/scheduler-backup-postgresql.git

* Create env.yml file
        
    - You can clone **.env.example.yml** in Environment folder and rename the file into **env.yml**
    - After that you have to setup database configuration and setup discord webhook.
    - Save the file.

* Setup Scheduler

    - We need to setup scheduler time on the **StartSchedulerBackup** function in **Scheduler/backup.go** file according to your needs.

* Path `pg_dump`  
    - if you use windows, you can change pgDump variable on line 28 with pg_dump.exe path. For example, **`pgDump := C:\\Program Files\\PostgreSQL\\16\\bin\\pg_dump.exe" `**

## Executing Program
Once all the steps are complete, **it's showtime**

    go run .

ensure the program running well. Wait until the time set in the scheduler, and the backup file will be sent to the notifiers.

## Acknowledgments
* [gocron: A Golang Job Scheduling Package](https://github.com/go-co-op/gocron?tab=readme-ov-file

