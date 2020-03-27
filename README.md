<p align="center">
	<a href="https://caddyserver.com"><img src="https://github.com/leopedroso45/ViLoW/blob/master/app/web/ViLoWofc.png" alt="ViLoW"></a>
</p>

<!-- ![alt text](https://github.com/leopedroso45/ViLoW/blob/master/app/web/ViLoWofc.png) --> 

<h3 align="center">ViLoW is a system that displays videos shared by the user on a platform for access on mobile devices or desktop on home network.</h3>

<p align="center">These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.</p>

### Prerequisites

- [Docker](https://docs.docker.com/)
- [Docker-compose](https://docs.docker.com/compose/install/)

or

- [Go](https://golang.org/dl/)
- [mySQL](https://www.mysql.com/downloads/)

### Installing

- After installing the tools mentioned above you are ready to run the application and / or contribute to the project.

- Now you can clone this [repository](https://github.com/leopedroso45/ViLoW) and open it in your favorite editor, I recommend [VSCode](https://code.visualstudio.com/)!  :)

### You need to modify some paths for everything to work properly!

- By default, the volume folder where videos should be stored is **C:\data**, but you can easily modify this in the Docker-compose file.

- In the db_connection file you can edit the database settings, such as the database name, root name and password that are standard. (If you're running on docker, you don't have to worry about)

- Now you can run: 
```
docker-compose run --build
```

## Built With

* [Go SQL Driver](https://github.com/go-sql-driver/mysql) - The mysql driver used
* [Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers
* [MySQL](https://www.mysql.com/) - is a relational database management system based on SQL – Structured Query Language.
* [Docker](https://docs.docker.com/) - Docker provides a way to run applications securely isolated in a container, packaged with all its dependencies and libraries.

The client has not yet been developed, some technologies are still being studied. Like: web flutter, ReactJs, VueJs

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Leonardo Severo** - *Nebullus* - [leopedroso45](https://github.com/leopedroso45)

See also the list of [contributors](https://github.com/leopedroso45/ViLoW/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
