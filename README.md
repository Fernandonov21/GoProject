# Go "Hello World"

This project is a simple "Hello World" web application built using **Go (Golang)**. It demonstrates how to set up a basic web server with Go and how to run it locally or in a Docker container.

## Table of Contents

- [Project Description](#project-description)
- [Requirements](#requirements)
- [Installation and Local Setup](#installation-and-local-setup)
- [How to Use the Project](#how-to-use-the-project)
- [Docker Setup](#docker-setup)
- [License](#license)

## Project Description

This project creates a simple Go web server that listens on a specified port and displays a "Hello World!" message when you visit the homepage. This project is ideal for beginners looking to understand the basics of web development with Go.

### Features:
- Minimal setup with Go
- A single route that displays "Hello World!"
- Ready to be run locally or in a Docker container

## Requirements

Before you begin, make sure you have the following installed:

- **Go** (version 1.16 or higher): [Download Go](https://golang.org/dl/)

## Installation and Local Setup

Follow these steps to set up the project on your local machine:

1. **Clone the repository** to your local machine:

   ```bash
   https://github.com/Fernandonov21/GoProject.git
   ```
2. **Open a PowerShell in the Project Directory**
Navigate to the project directory where the repository was cloned. To do this:
Open the PowerShell or Terminal (on macOS/Linux).
If you work with Visual Studio Code, open the folder project and open a PowerShell:

   ![image](https://github.com/user-attachments/assets/1eef3341-14c9-404e-8573-0c232141f719)

  
   
4. **Run the Flask application**
Once the dependencies are installed, run the app with the following command:

    ![image](https://github.com/user-attachments/assets/60bf38aa-63b3-4fb7-93a6-703f7a843d75)

    ```bash
    go run main.go
    ```
 

**Access the application**

Open your web browser and go to http://127.0.0.1:8080 to see the "Hello, World!" message.

## How to configure a docker image
The project has already a docker file so the only thing that you need is up the container
1. In the terminal (you can use Command Prompt, PowerShell, or the VS Code integrated terminal), navigate to the project folder and execute the next command:
```bash
docker build -t goptproject .
```
2. run the image
```bash
docker run -p 8080:8080 goproject
```
## IMPORTANT
To access the application at http://localhost:8080, ensure that port 8080 is available. If the port is already in use, stop any process occupying it and try running the command again.

