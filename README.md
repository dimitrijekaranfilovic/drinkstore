# alkostore
Predmetni projekat iz Naprednih tehnologija programiranja.

## Student
Dimitrije Karanfilović, SW-39/2018

## Opis sistema
Sistem je namijenjen pretrazi, ocjenjivanju i naručivanju alkoholnih pića. 

## Funkcionalnosti
Sistem sadrži tri kategorije korisnika:
 * neautentifikovani korisnik
 * autentifikovani korisnik
 * administrator


Neautentifikovani korisnik može da pretražuje i sortira pića, ali ne može da ih naručuje, ocjenjuje niti može da ostavlja komentare (za svako piće postoji sekcija sa komentarima). Takođe će imati uvid u to koja su pića bila najprodavanija u određenom vremenskom periodu. Neutentifikovani korisnik može i da se registruje.

Autentifikovani korisnik može da radi sve što i neutentifikovani, a dodatno može da naručuje pića, da ih ocjenjuje i da ostavlja komentare, s tim što to može da uradi samo za pića koja je naručio barem jednom. Na svakih 1500 RSD korisnik dobija 1 poen, a svaki poen nosi 3% popusta. Korisnik može da koristi poene u kupovinama, ali maksimalno 10 po kupovini. Takođe će imati uvid u svoju istoriju kupovine. Korisnici takođe mogu da prijave komentare drugih korisnika (u slučaju vrijeđanja ili neprimjerenog sadržaja).

Administrator ne može da naručuje niti ocjenjuje, ali može da dodaje pića u katalog, da ih briše i mijenja. Takođe, može da briše komentare i da banuje korisnike. Pored toga, moći će da pregleda prijave korisnika i da odlučuje da li će ih prihvatiti ili ne.

Piće će imati minimum sledeće podatke:
 * naziv
 * slika
 * opis
 * zapremina
 * procenat alkohola
 * prosječna ocjena
 * komentari

Još neki podaci se mogu dodati po potrebi.

## Arhitektura sistema
Sistem je zasnovan na mikroservisnoj arhitekturi.


**Mikroservisi**
|Naziv servisa|Opis|Tehnologije|
|-------------|----|-----------|
|**UserService**|servis za autentifikaciju, autorizaciju, dodavanje i banovanje korisnika|Go, PostgreSQL|
|**CommentService**|servis za upravljanje komentarima i njihovim prijavama|Go, MongoDB|
|**DrinkService**|servis za upravljanje pićima|Go, PostgreSQL|
|**PurchaseService**|servis za obavljanje same kupovine|Pharo, PostgreSQL|

**Ostale komponente**
|Naziv komponente|Opis|Tehnologije|
|----------------|----|-----------|
|**Frontend aplikacija**|Monolitna frontend aplikacija sa svim funkcionalnostima|VueJS|
|**API gateway**|api gateway za mikroservise|Apache APISIX ili Python Django ili SpringBoot|





