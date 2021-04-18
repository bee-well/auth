# Auth Tjänst
Detta repository är en prototyp av autensieringstjänsten för Bee Well. Detta program kommer att sköta användare och genereringen av JWT-tokens. 
## Tech
Servern är skriven i Go och använder sig utav ramverket `Gin` för att sköta HTTP förfrågningar. Användare sparas i en PostgreSQL-databas med följande fält: 
* ID
* Förnamn
* Efternamn
* Lösenord
* Email
## Köra för utveckling
För att starta programmet kan du använda Docker Compose (se `docker-compose.yaml`). Se till att du har detta installerat (på MacOS kommer detta installerat med Docker). Sedan är det bara att navigera sig till projektmappen i terminalen och köra `docker-compose up` för att starta RabbitMQ, PostgreSQL och Auth-tjänsten.
