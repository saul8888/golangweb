### Create keys

it is created in the tokens folder

**public key**

`openssl rsa -in private.rsa -pubout > public.rsa.pub`

**private key**

`openssl genrsa -out private.rsa 1024`

**compilate**

`sudo docker-compose up`