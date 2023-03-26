# GDSC_Backend
An android application using AR technology to create a safe living environment for the aging.

## Get Start
1. Download all necessary dependencies.

    ``go mod tidy `` 

2. Install necessary DB.

    ``brew install mysql``

3. Create local database `gdsc`

4. Execute the [init.sql](config%2Finit.sql) file to construct table

    ``go run init.sql``

5. Start server

   ``go run main.go``

Then, you can test the endpoint by port: http://localhost:8080/