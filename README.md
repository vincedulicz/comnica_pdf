# Comnica Sign-Backend Integráció

Go nyelven írt alkalmazás, amely integrálja a Comnica aláíró alkalmazását

a [https://sign-test.comnica.com](https://sign-test.comnica.com) szerveren keresztül. 

A program az alábbi feladatokat valósítja meg:

- **Aláíró rekord létrehozása:**  
  Létrehoz egy új aláíró rekordot a Comnica rendszerében
  
- **PDF feltöltése:**  
  Feltölt egy tetszőleges PDF fájlt az adott aláíró rekordhoz
  
- **Rekord véglegesítése:**  
  Véglegesíti a rekordot, és visszaadja a rekord azonosítóját, valamint a felhasználók számára elérhető aláírási linket.
  
- **HTTP szerver:**  
  A program maga is egy szerver, mely HTTP POST kérések segítségével fogadja a PDF feltöltéseket

---

## Feladat Leírása

A projekt célja, hogy a Comnica OpenAPI specifikációja szerint kommunikáljon 

a [sign-test.comnica.com](https://sign-test.comnica.com) szerverrel. 

A feladat során:

1. **Session létrehozása:**  
   Egy új aláíró session jön létre, amelyhez autentikációs token és session ID tartozik.
   
3. **Dokumentum feltöltése:**  
   A feltöltött PDF fájl a session-hoz kapcsolódik, a feltöltés során a dokumentum adatai elküldésre kerülnek.
   
5. **Visszajelzés:**  
   A rendszer visszaadja a feltöltött dokumentum azonosítóját és az aláírás elvégzéséhez szükséges URL-t.

---

## Használat

go run main.go

## PDF Feltöltés

1. Nyisd meg a böngészőt és navigálj a következő címre: [http://localhost:8080](http://localhost:8080).
2. Töltsd fel a PDF fájlt az űrlapon keresztül.
3. A feltöltés után a program a következőket hajtja végre:
   - Létrehozza az aláíró session-t.
   - Feltölti a PDF fájlt a Comnica rendszerébe.
   - Visszaadja a session ID-t, a dokumentum azonosítóját, és az aláírás elvégzéséhez szükséges URL-t.
4. A naplózás a terminálban történik, ahol a session és a dokumentum azonosítók is megjelennek.

## Kódstruktúra

- **`services/document.go`**: A PDF feltöltés logikája a Comnica szerver felé, hibakezeléssel és multipart request összeállítással.
- **`services/session.go`**: A session létrehozásának logikája, amely JSON payload segítségével kommunikál a szerverrel.
- **`handlers/upload.go`**: A HTTP feltöltési űrlap és feltöltési folyamat kezelése. Itt történik a fájl olvasása, ideiglenes tárolása, és a session indítása.
- **`templates/index.html`**: Egy egyszerű webes felület a PDF feltöltéséhez.
