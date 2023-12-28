
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

###### Cara menjalankan applikasi:

1. Clone project repository
```
git clone git@github.com:MarioTiara/todolist-app-echo.git
```
2. Masuk ke project folder
```
 cd todolist-app-echo
```
3. Jalankan docker-compose (pastikan docker sudah berjalan di komputer anda)
```
  docker-compose up
```
4. Apply database migrations
```
  make migration-up
``` 
