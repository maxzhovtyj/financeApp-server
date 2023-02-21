## FinanceApp REST API

* Go 1.18
* Echo framework
* MongoDB

Create .env file in root directory with the following values:
```
HTTP_HOST=localhost

MONGO_URI=mongodb://mongodb:2717
MONGO_USER=admin
MONGO_PASS=qwerty

PASSWORD_SALT=<random string>
JWT_SIGNING_KEY=<random string>
```