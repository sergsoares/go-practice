# Echo Quick start

Following guide https://echo.labstack.com/guide/

- http://localhost:8080/show?team=x-men&member=wolverine

# validate /save
curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:8080/save

# upload image
<!-- $ curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:808/save -->
curl -F "name=Joe Smith" -F "avatar=https://cdn.labstack.com/images/echo-logo.svg" http://localhost:8080/savefile
