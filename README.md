# Quickstart Guide to Set up Dev Workspace

1. Git clone https://github.com/mrmurazza/ijah_store.git
2. Make sure you have Go installed and set up Go workspace
3. Install sqlite3 driver library for database driver and Gin-Gonic for web framework using `go get` like this:

   `go get github.com/mattn/go-sqlite3`
   
    `go get -u github.com/gin-gonic/gin`
    
4. Go to your project folder using `cd` and run your project using:

    `$ go run main.go`
    
5. For deployment, run this command to build binary file

    `$ go build main.go`
    
    
# Project Structure 
```
.
├── src/app
│   ├── item                         > package contains Item-related code 
│   ├── purchase                     > package contains purchase-related code
│   ├── request                      > package contains request struct code
│   ├── response                     > package contains response struct code
│   ├── restock                      > package contains restock-related code
│   ├── util                         
│   └── controllers.go               > file contains input/output preparation
├── README.md
├── ijah_store.db                   > SQLite DB file
├── inventory_report_17_01_19.csv   > Example output of inventory reports in CSV
├── main.go                         > Main Go file
└── sales_report_17_01_19.csv       > Example output of sales reports in CSV 

```

# Project Component

Basically, this project is separated into several components: 
1. Controller : responsible on preparing the inputs and serving the outputs and the one to call the business logic code.
2. Handler : responsible on handling the main core of the business logic. (ex: `item_handler.go, restock_handler.go, & purchase.handler)`
3. Model : responsible on containing the struct and queries. (ex: `item.go, restock_resception.go, etc`)


# Notes Regarding the Assessment:

Info Test BE
1. Dengan NodeJS (atau Golang), agar saudara membuat Rest API CRUD User dan User Login.
2. Jika menggunakan NodeJS maka disarankan menggunakan ExpressJS. Database bebas, tetapi disarankan MongoDB.
3. User Login digunakan user (username, password) untuk mengakses API CRUD (token, tetapi mendapatkan nilai tambahan jika menggunakan refresh token).
4. Bikin 2 users dengan role: 1 Admin, 1 User.
5. Admin bisa melakukan/mengakses semua API CRUD, sedangkan User hanya bisa mengakses data user bersangkutan saja (Read)
6. Implementasi arstektur Microservices, menggunakan Kubernetes dengan Docker container deploy di VPS (1 node dengan beberapa pod di dalamnya). Bagi yang belum memiliki VPS, maka cukup (a) menyiapkan semua YML agar aplikasi bisa dijalankan secara containerize dan siap di deploy di Kubernetes dan (b) di-deploy di lokal dan sertakan screenshoot.
7. Upload source code ke Github beserta script YML Kubernetes.
8. Bikin dokumentasi API nya (Postman atau Swagger) yang bisa diakses ke server Rest API nya.
9. Bikin diagram arstektur nya yang menjelaskan flow API CRUD dan Login.
10. Lampirkan credential Admin di Readme.

Mohon submit via whatsapp ini paling lambat Selasa, 13 September jam 20.10 WIB.*

Working here is really fulfilling & the salary makes it extra rewarding. Goodluck and let me know if you need anything!! :)
