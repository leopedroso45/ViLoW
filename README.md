<p align="center">
	<!-- ![alt text](https://github.com/leopedroso45/ViLoW/blob/master/app/web/ViLoWofc.png) --> 
	<a href="#"><img src="https://github.com/leopedroso45/ViLoW/blob/master/web/ViLoWofc.png" alt="ViLoW">	
	</a>
</p>

<div align="center">
	
[![Go Report Card](https://goreportcard.com/badge/github.com/leopedroso45/vilow)](https://goreportcard.com/report/github.com/leopedroso45/ViLoW)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c402f859fde040c9af9f5b9e2e272d31)](https://app.codacy.com/manual/leopedroso45/ViLoW?utm_source=github.com&utm_medium=referral&utm_content=leopedroso45/ViLoW&utm_campaign=Badge_Grade_Dashboard)
[![Maintainability](https://api.codeclimate.com/v1/badges/974ea2165b7a5ac93d63/maintainability)](https://codeclimate.com/github/leopedroso45/ViLoW/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/974ea2165b7a5ac93d63/test_coverage)](https://codeclimate.com/github/leopedroso45/ViLoW/test_coverage)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/leopedroso45/vilow)

</div>

<h3 align="center">ViLoW is a system that displays videos shared by the user on a platform for access on mobile devices or desktop on home network.</h3>

<p align="center">These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.</p>
 

### Prerequisites

- [Docker](https://docs.docker.com/)
- [Docker-compose](https://docs.docker.com/compose/install/)

### Installing

- After installing the tools mentioned above you are ready to run the application and / or contribute to the project.

- Now you can clone this [ViloW Server](https://github.com/leopedroso45/ViLoW), [ViLoW Client](https://github.com/leopedroso45/ViLoW_Client) and open it in your favorite editor, I recommend [VSCode](https://code.visualstudio.com/)!  :)

### You need to modify some paths for everything to work properly

- By default, the volume folder where videos should be stored is **C:\data**, but you can easily modify this in the Docker-compose file.

- Now you can run: 
```docker-compose run --build```

## Built With

* [Go SQL Driver](https://github.com/go-sql-driver/mysql) - The mysql driver used
* [Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers
* [MySQL](https://www.mysql.com/) - is a relational database management system based on SQL – Structured Query Language.
* [Docker](https://docs.docker.com/) - Docker provides a way to run applications securely isolated in a container, packaged with all its dependencies and libraries.

The client was built using Reactjs, check it out:

- [ViLoW](https://github.com/leopedroso45/ViLoW_Client)

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Leonardo Severo** - *Nebullus* - [leopedroso45](https://github.com/leopedroso45)

See also the list of [contributors](https://github.com/leopedroso45/ViLoW/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
