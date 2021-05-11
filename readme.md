# Auth Tjänst 🔐
Detta repository är en prototyp av autensieringstjänsten för Bee Well. Detta program sköter användare och skapandet av JWT tokens.  

## Tech 💻
Servern är skriven i Go och använder sig utav ramverket `Gin` för att sköta HTTP förfrågningar. Användare sparas i en PostgreSQL-databas. Alla lösenord som sparas i PostgreSQL är hashade med BCrypt. Programmet körs i en Docker-container på Heroku Cloud och deployas ✨ automagiskt ✨ med CD. I detta repot, precis som alla andra under [bee-well](https://github.com/bee-well), så sköts utveckling på en separat branch med pull requests till main efter att en feature är klar. När en pull request skapas så körs alla tester ✨ automagiskt ✨ här på GitHub (CI). Både CI och CD är implementerat med hjälp av GitHub Actions ([se .github/workflows](https://github.com/bee-well/auth/tree/main/.github/workflows)). Efter att en pull request har förts till main branchen så deployas applikationen i Heroku. 

## Docker 🐳
Detta program finns inte på Docker Hub utan för att köra det i Docker måste det laddas ned härifrån. För att enklast starta upp programmet med dependencies såsom PostgreSQL kan [docker-compose](https://docs.docker.com/compose/) användas. För att se programmets dependencies och miljövariablar, se [`docker-compose.yaml`](https://github.com/bee-well/auth/blob/main/docker-compose.yaml). Följande kommando bör användas för att starta upp programmet: `docker-compose up --build`, en annan variant som använder sig av Docker CLI:t är `docker compose up --build`, men jag rekommenderar att den förstnämnda används. 

## JWT
De JWT tokens som skapas innehåller följande data efter att de avkodas:
```
{
  "id": Integer,
  "issued": String
}
```
Notera att `issued` inte används för tillfället utan istället fungerar som en framtidsförsäkring för att kunna sätta en tidsgräns på hur länge en token är giltig. Gruppen är medveten om att JWT tokens enkelt kan avkodas och sparar därför inte någon känslig data där i.

## Endpoints 👇
Nedan finns en förklaring av alla endpoints som denna tjänst innehåller. 
### GET `/users` 🧮
Denna endpoint returnerar antalet användare som är registrerade i systemet. Den används utav UI:n för att skriva ut detta antal på förstasidan.
#### Förfrågan
Förfrågan kräver ingen data.
#### Respons
Det enda som skulle resultera i att något annat än `200 OK` returneras från tjänsten är om databasen ligger nere, i det fallet returneras `500 Internal Server Error`. Datan som skickas tillbaka med responses är presenterad nedan.
```
{
  "count": Integer
}
```
#### Flöde
![users](https://user-images.githubusercontent.com/36814950/117797711-9e6cf280-b250-11eb-9c44-a7d1fc463da5.png)

### POST `/sign-up` ➕
Denna endpoint används för att registrera en ny användare. 
#### Förfrågan
Följande data måste inkluderas i förfrågningens body.
```
{
  "email": String,
  "firstName": String,
  "lastName": String,
  "password": String,
}
```
#### Respons
Tjänsten besvarar förfrågningen med `200 OK` om allting gick som det skulle. Om användaren har försett tjänsten med bristfällig data så svarar tjänsten med `400 Bad Request` tillsammans med ett errormeddelande som förklarar felet i datan. Kraven för att valideringen ska gå igenom är ett lösenord på minst 8 karaktärer samt en giltig email adress, utöver detta så måste även förnamn och efternamn vara ifyllt. Valideringen för att mailadressen ska vara unik sker på databasen, om mailadressen inte är unik så svarar tjänsten med `400 Bad Request`. 

![sign-up](https://user-images.githubusercontent.com/36814950/117798719-b133f700-b251-11eb-9c5a-5858adbc0950.png)

### POST `/sign-in` 🔑
Denna endpoint används för att generera en JWT token som användaren kan skicka vidare till andra tjänster för att autensiera sig.
#### Förfrågan
Förfrågan måste innehålla följande data:
```
{
  "email": String,
  "password": String
}
```
#### Respons
Om förfrågan inte innehåller korrekt data så besvarar tjänsten med `400 Bad Request`, med "korrekt data" menas här att både mailadress och lösenord måste skickas med. Om uppgifterna inte stämmer överens med något registrerat konto skickas `401 Not Authorized` tillbaka. Vid en lyckad autensiering så skickas `200 OK` tillbaka tillsammans med en JWT token i plain-text. 

#### Flöde

![sign-in](https://user-images.githubusercontent.com/36814950/117806517-b9dcfb00-b25a-11eb-8547-842fe75b6190.png)

### GET `/me` 👨
Detta är den enda endpointen i denna tjänst som kräver att `Authorization`-headern är satt och innehåller en giltig JWT token. Endpointen används för att hämta data angående den inloggade användaren.
#### Förfrågan
Ingen body data krävs för förfrågan, men `Authorization`-headern måste innehålla en giltig JWT token.
#### Respons
Om användaren har skickar med en giltig token så returneras hens data tillsammans med `200 OK` status. Datan är formatterad som följande:
```
{
  "email": String,
  "firstName": String,
  "lastName": String
}
```
Om användaren inte är inloggad så skickas `401 Not Authorized` tillbaka utan någon body data.

#### Flöde
![me](https://user-images.githubusercontent.com/36814950/117807093-75059400-b25b-11eb-86b5-626791e98b2f.png)

## Contribution 👨‍👦‍👦
*För att göra ändringar i detta eller andra repositories för Bee Well, vänligen utveckla på en separat branch och skapa en PR till main branchen då CI/CD är aktiverat så fungerande kod måste alltid befinna sig på `main`*
