<h1 align="center">
  ITU-minitwit
</h1>
<p align="center">
ITU Minitwit application rewritten in GO with the Gin library.
<br/>

## ⚡️ Quickstart

```sh
cd minitwit-go
docker-compose up
```
`docker-compose up` runs both the database and the server


If files have been changed, then instead run
```sh
docker-compose up --build
```




## 📝 Notes from lectures!

### **Lecture 1️⃣**
<hr/>

* We converted from python 2 -> 3. Simply by trying to run it and then correct where the compliler complained.
* We changed print statement! Just itereated over each complie error. Also we changed the .control file (added a 3 and ""). 
* We forgot to change the test file!
* (Good to use 2To3 on that test file)

### **Lecture 2️⃣**
<hr/>

* We went though the docker exercise and started talking about how to structure our git (having a folder for the old minitwit application and another one for the one we are refactoring). 
* We creating a docker file and refactoring to Go.
* We also want to use postgres because it is a well known database and easy to implement in docker.
