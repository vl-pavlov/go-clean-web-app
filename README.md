# Golang Clean Architecture Web App

You can use this template for building your web app to ensure that our application is easy to maintain, scalable, and robust. 

We also recommend exploring the [basic layout](https://github.com/golang-standards/project-layout) for Go application projects  

## Getting Started

To get started, follow these steps:

```bash
# Clone the repository
git clone https://github.com/vl-pavlov/go-clean-web-app.git

# Navigate to the project directory
cd go-clean-web-app

# Build the project
go build

# Run the web server
./go-clean-web-app
```

## Project Structure

### `/cmd`
Main applications for this project.

### `/pkg`
Library code that's ok to use by external applications.

### `/internal`
Private application and library code.

### `/api`
API protocol definition files.

## License
Distributed under the MIT License. See [LICENSE](LICENSE) for more information.