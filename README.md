# Cohort 2 - Mini Challenge #2
### Calculating Words Appear

1. Buatlah looping dengan variable yang berisi string suatu kalimat dan pecahlah kalimat tersebut menjadi 1 per 1 kata
2. Setelah sudah dipecah, lakukan perhitungan munculnya kata dari variable tersebut dengan cara mapping golang

##### Parameter:
Program ini bisa dijalankan dengan atau tanpa parameter/argumen
Jika menggunakan argumen, maka teks yang diproses adalah argumen tersebut.
Jika tidak mengguanakan, argumen maka teks yang diproses adalah "selamat siang/pagi/malam" tergantung waktu eksekusi program.

<!-- ```bash -->

##### Sample Output Tanpa Argumen:
```bash
$ go run main.go
s
e
l
a
m
a
t

m
a
l
a
m

map[ :1 a:4 e:1 l:2 m:3 s:1 t:1]
```

##### Sample Output Dengan Argumen:
```bash
$ go run main.go "Ular Melingkar Di Atas Pagar"
U
l
a
r

M
e
l
i
n
g
k
a
r

D
i

A
t
a
s

P
a
g
a
r

map[ :4 A:1 D:1 M:1 P:1 U:1 a:5 e:1 g:2 i:2 k:1 l:2 n:1 r:3 s:1 t:1]
```
