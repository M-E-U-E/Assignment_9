# Assignment09


## Table of Contents
- [Description](#description)
- [Git Clone Instructions](#git-clone-instructions)
- [Test](#test)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Code Coverage Report](#code_coverage_report)



## Description

The project aims to recreate the features of TheCatAPI website, focusing on displaying cat-related content and functionality. The backend will be implemented using Beego, leveraging its controllers, template rendering, and configuration capabilities, while the frontend will be built using Vue.js, HTML, and CSS to deliver an interactive user interface.
## Git Clone Instructions

To clone this project to your local machine, follow these steps:

1. **Open terminal (Command Prompt, PowerShell, or Terminal)**

2. **Clone the repository**:
   
         git clone https://github.com/M-E-U-E/Assignment_9.git or git clone git@github.com:M-E-U-E/Assignment_9.git
   
    Go to the Directory:
    ```bash
    cd Assignment_9
    ```

    **The conf/app.conf file should already be set up in your project root with configurations such as appname, httpport, runmode, cat_api_key, and StaticDir. If you do not want to use the existing file, you can manually create one with the following content:
    Install Dependencies**
    ```
    appname = cat_api
    httpadr = "localhost"
    httpport = 8080
    runmode = dev
    staticfile_cdn_headers = Cache-Control: no-cache, no-store, must-revalidate
    api_key = "live_G9HhPFIdEQOKjegSsyNOQ8lRoWxFkzxttwXgu7F0gCFLhYVlX0F1cIVGMADW6rtg"
    api_base_url = https://api.thecatapi.com/v1
    ```
3. **To Run the Application:**
    ```
    go install github.com/beego/bee/v2@latest
    go mod tidy
    bee run
    ```
4. **Access Endpoints**
    - http://localhost:8080/voting
    - http://localhost:8080/breeds
    - http://localhost:8080/favs



## Test
  Run the testing file:
  Ensure the beego is install:
   Use this comman to view code coverage  report
   ```
   go test ./... -coverprofile=coverage.out
   go tool cover -html=coverage.out                                      |
     
   ```
 
    

## Project Structure
```
Assignment_9/                          # Root project directory
├── conf/                         # Configuration files
│   └── app.conf                  # Main application configuration
│
├── controllers/                  # Application controllers
│   ├── breeds_controller.go
│   ├── breeds_controller_test.go
│   ├── config_controller.go
│   ├── default.go
│   ├── favs_controller.go
│   ├── favs_controller_test.go
│   ├── main_controller.go
│   ├── main_controller_test.go
│   ├── voting_controller.go
│   └── voting_controller_test.go
│
├── models/                       # Data models
│   └── models.go
│
├── routers/                      # Route definitions
│   └── router.go
│
├── static/                       # Static files for the application
│   ├── css/
│   │   └── all css            # CSS files
│   ├── img/
│   │   └── none              # Image assets
│   └── js/
│       └── all js              # JavaScript files
│
├── tests/                        # Unit and integration tests
│   ├── default_test.go
│
├── views/                        # Template files
│   ├── breeds.tpl
│   ├── favs.tpl
│   ├── index.tpl
│   └── voting.tpl
│
├── .gitignore                    # Files and directories to ignore in Git
├── go.mod                        # Go modules file
├── go.sum                        # Dependency checksums
├── main.go                       # Application entry point
├── README.md                     # Project documentation

```
## Technologies Used

- ![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  
  **Go (Golang)**: The main programming language for developing the application.

- ![Beego](https://img.shields.io/badge/Beego-%23F15A24.svg?style=for-the-badge&logo=beego&logoColor=white)  
  **Beego Framework**: A popular Go web framework, indicated by the directory structure (e.g., controllers, models, routers).

- ![HTML](https://img.shields.io/badge/HTML5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)  
  **HTML Templates**: For rendering views, as suggested by the `.tpl` files in the `views` directory.

- ![CSS](https://img.shields.io/badge/CSS3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white)  
  ![JavaScript](https://img.shields.io/badge/JavaScript-%23F7DF1E.svg?style=for-the-badge&logo=javascript&logoColor=black)  
  **CSS/JavaScript**: For styling and interactivity, located in the `static` folder.

- ![Go Testing](https://img.shields.io/badge/Go%20Testing-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  
  **Go Testing Package**: For unit and integration tests, as seen in the `*_test.go` files.

- ![Beego ORM](https://img.shields.io/badge/Beego%20ORM-%23F15A24.svg?style=for-the-badge&logo=beego&logoColor=white)  
  **Beego ORM**: If Beego is used, its built-in ORM is likely used for database interactions.


  

## Code Coverage Report:
Use this comman to view code coverage  report
```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out                                      |
  
```
  

