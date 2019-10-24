# VideoLocalManager

![alt text](https://github.com/leopedroso45/VideoLocalManager/blob/master/web/frontWeb.PNG)

## Getting Started

Local platform for managing and consuming videos stored on the hard disk.
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Golang
- mySQL
- Bootstrap

### Installing

- You'll need to install [Golang](https://golang.org/dl/) on your machine.
  You will need to define your GOPATH and those imports in the .go files. (Make sure the path of the imported files conforms to your GOPATH.) 
- You'll need to install [Xampp](https://www.apachefriends.org) (Or manually install MySQL)
- You can now clone this [repository](https://github.com/leopedroso45/VideoLocalManager) and import bootstrap into the web project.

- [ **Finally, you can now set the path to the folder where you store your videos and the path in the feedDB.go file so that it can identify them.** ]

- Now you are available to enable apache server, golang Rest API server and mySQL to open the web interface.

```
go run main.go model.go feedDB.go
```

## Built With

* [Go SQL Driver](https://github.com/go-sql-driver/mysql) - The mysql driver used
* [Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers
* [Bootstrap](https://getbootstrap.com/) - Front-end component library used

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Leonardo Severo** - *Nebullus* - [leopedroso45](https://github.com/leopedroso45)

See also the list of [contributors](https://github.com/leopedroso45/VideoLocalManager/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
