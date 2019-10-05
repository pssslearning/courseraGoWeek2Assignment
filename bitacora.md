# BITACORA

## 1. Creating project folder by cloning an scheleton github project formerly created

```sh
devel1@vbxdeb10mate:~$ mkdir -pv $GOPATH/src/github.com/pssslearning/
mkdir: se ha creado el directorio '/home/devel1/gowkspc/src/github.com'
mkdir: se ha creado el directorio '/home/devel1/gowkspc/src/github.com/pssslearning/'

devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning$ git clone https://github.com/pssslearning/courseraGoWeek2Assignment.git
Clonando en 'courseraGoWeek2Assignment'...
remote: Enumerating objects: 10, done.
remote: Counting objects: 100% (10/10), done.
remote: Compressing objects: 100% (8/8), done.
remote: Total 10 (delta 1), reused 0 (delta 0), pack-reused 0
Desempaquetando objetos: 100% (10/10), listo.
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning$ ls -la courseraGoWeek2Assignment/
total 20
drwxr-xr-x 3 devel1 devel1 4096 oct  3 19:18 .
drwxr-xr-x 3 devel1 devel1 4096 oct  3 19:18 ..
drwxr-xr-x 8 devel1 devel1 4096 oct  3 19:18 .git
-rw-r--r-- 1 devel1 devel1 1069 oct  3 19:18 LICENSE
-rw-r--r-- 1 devel1 devel1  596 oct  3 19:18 README.md
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning$ mkdir -pv courseraGoWeek2Assignment/trunc
mkdir: se ha creado el directorio 'courseraGoWeek2Assignment/trunc'
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning$ touch -v courseraGoWeek2Assignment/trunc/trunc.go
touch: opción inválida -- 'v'
Pruebe 'touch --help' para más información.
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning$ touch courseraGoWeek2Assignment/trunc/trunc.go
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning$ ls -la courseraGoWeek2Assignment/
total 24
drwxr-xr-x 4 devel1 devel1 4096 oct  3 19:19 .
drwxr-xr-x 3 devel1 devel1 4096 oct  3 19:18 ..
drwxr-xr-x 8 devel1 devel1 4096 oct  3 19:18 .git
-rw-r--r-- 1 devel1 devel1 1069 oct  3 19:18 LICENSE
-rw-r--r-- 1 devel1 devel1  596 oct  3 19:18 README.md
drwxr-xr-x 2 devel1 devel1 4096 oct  3 19:19 trunc
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning$ ls -la courseraGoWeek2Assignment/trunc
total 8
drwxr-xr-x 2 devel1 devel1 4096 oct  3 19:19 .
drwxr-xr-x 4 devel1 devel1 4096 oct  3 19:19 ..
-rw-r--r-- 1 devel1 devel1    0 oct  3 19:19 trunc.go

devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning$ cd courseraGoWeek2Assignment/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek2Assignment$ code .
```

## 2. Compliling and testing

```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/
devel1@vbxdeb10mate:~/gowkspc/src$ cd github.com/pssslearning/courseraGoWeek2Assignment/trunc/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek2Assignment/trunc$ go build trunc.go 
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek2Assignment/trunc$ ./trunc 

Welcome user to Assignment program for Week 2,

--------------------------------------------------------------
Please enter a float number to trunc ... 
      (press <CTRL+C> or '0.0' to exit.) 
--------------------------------------------------------------
189.3456 
The number you entered once parsed as a float number is now: [189.345600]
The number you entered once transformed in integer and then truncated is now: [189]
--------------------------------------------------------------
Please enter a float number to trunc ... 
      (press <CTRL+C> or '0.0' to exit.) 
--------------------------------------------------------------
456.789
The number you entered once parsed as a float number is now: [456.789000]
The number you entered once transformed in integer and then truncated is now: [456]
--------------------------------------------------------------
Please enter a float number to trunc ... 
      (press <CTRL+C> or '0.0' to exit.) 
--------------------------------------------------------------
5e-5
The number you entered once parsed as a float number is now: [0.000050]
The number you entered once transformed in integer and then truncated is now: [0]
--------------------------------------------------------------
Please enter a float number to trunc ... 
      (press <CTRL+C> or '0.0' to exit.) 
--------------------------------------------------------------
0.0

Goodbye !!!
```

## 3. Testing behaviour for an incorrect input (not parseable as float number)

```sh
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek2Assignment/trunc$ ./trunc

Welcome user to Assignment program for Week 2,

--------------------------------------------------------------
Please enter a float number to trunc ... 
      (press <CTRL+C> or '0.0' to exit.) 
--------------------------------------------------------------
This is not a floating point number !!!
An error has happenned. You probably entered something that cannot be interpreted as a float number
The error detected was [strconv.ParseFloat: parsing "": invalid syntax]
The program will be exited.
Please note that may be remaining input not consumed and sent to your input. 
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek2Assignment/trunc$ his is not a floating point number !!!
his is not a floating point number ./trunc!
bash: his: orden no encontrada
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek2Assignment/trunc$ 
```