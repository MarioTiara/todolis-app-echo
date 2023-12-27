
## Moonlay Academy - Backend Test (GOLANG)

Anda akan membuat backend handler untuk aplikasi **Todo List**.
>**Todo List App**
Aplikasi yang umumnya digunakan untuk memelihara tugas sehari-hari atau membuat daftar semua yang harus dilakukan, dengan urutan prioritas tugas tertinggi hingga terendah. Sangat membantu dalam merencanakan jadwal harian.

###### Spesifikasi Bisnis:
```
- Menampilkan, menambahkan, mengubah dan menghapus data.
- Dapat menambahkan sub list untuk setiap list yang terdaftar. Sub list bisa ditambah, diubah dan dihapus. ( hanya 2 level )

Data Input untuk masing-masing list/sub list :
- title [required | maximum 100 karakter | alphanumeric (including space)]
- description [required | maximum 1000 karakter]
- file [optional | hanya menerima file dengan extension .txt dan .pdf]
  Bisa upload lebih dari 1 file per list/sub list
```
Buat API Todo List dengan kriteria sebagai berikut :
1. [METHOD:GET] Menampilkan data all list ( include pagination, filter[Search By: title, description] ) dengan atau tanpa preload sub list (dynamic)
2. [METHOD:GET] Menampilkan data detail list by list id.
3. [METHOD:GET] Menampilkan data all sub list by list id ( include pagination, filter[Search By: title, description] )
4. [METHOD:GET] Menampilkan data detail sub list by sub list id.
5. [METHOD:POST] Menambahkan data list.
6. [METHOD:POST] Menambahkan data sub list untuk spesifik list.
7. [METHOD:POST/PUT] Mengubah data list/sub list dengan kritera input diatas. 
8. [METHOD:DELETE] Menghapus data list/sub list. 
9. API untuk create dan update harus dipisah.

###### Spesifikasi Teknis
1. Terdapat Unit Test dengan coverage min 65%
2. Menyertakan API Specification (Postman Collection atau Swagger Yaml)
3. Migration DB
4. Menyertakan panduan menjalankan program pada file Readme.md

>Jumlah API yang dibutuhkan untuk mengakomodir fitur maksimal 10.

### Teknologi dan library yang wajib digunakan:
| Teknologi   | Version | Link |
| ----------- | ---------------- | ------------------- |
| Golang      | v1.18 or later   | [Go Download](https://go.dev/dl)  |
| Go Echo Framework     | v4     | [Echo Installation](https://echo.labstack.com/guide/#installation) | 
| GORM | v2 | [GORM Installation](https://gorm.io/docs/#Install) |
| PostgreSQL | v13 or later | [PostgreSQL Download](https://www.postgresql.org/download/) |
<br>

```
Catatan :
Gunakan `SOLID` Principle
Reference :
- https://dave.cheney.net/2016/08/20/solid-go-design
- https://s8sg.medium.com/solid-principle-in-go-e1a624290346

```

Aplikasi ini adalah aplikasi berskala besar yang memiliki potensi perubahan atau penambahan fitur secara berkala sehingga membuat aplikasi ini akan menjadi lebih kompleks dan memerlukan lebih dari 1 developer dalam pengerjaannya (kolaborasi team).

**Code quality di aplikasi ini harus reusable, readable dan easy to upgrade.**

Sangat baik jika ada :
1. Migration Database ( database version control )
2. Memiliki kapabilitas untuk merubah DB Engine (misal PostgreSQL ke MySQL)
3. Memiliki kapabilitas untuk mengubah File Storage (misal dari Local ke AWS, atau dari Azure ke Google)

*Good luck and Have Fun.* :rocket:
