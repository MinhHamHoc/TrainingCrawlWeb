1. Place the readme texts like: setup step of project, or any other information about the project

2. Create .env file, where we keep all environment related constant, like database, username, password, secret key of jwt 

3. Create the "main.go" file in root, this will server as the project/webs-server entry point

5. Create these folders:
    1. Controllers: contains controller that accepts a request and call particular service to process

    2. Services: contains services that has the logic for doing all process like manipulating DB and give back the desired result

    3. Modals: contains the struct for storing data in DB or fetching data from DB

    4. Database: Keep Database related operation here, these can be used in services to fetch/update/insert/delete data from DB
        *****Place DB connection file as well inside this folder

    5. Routes: Keep all routes here and pass the name of controllers

    6. Config: Keep all your configuration things here like, fetching variable from .env file, or even DB config can be place here

    8. Ultilities: Keep commonly used function here like capitalising first letter, etc