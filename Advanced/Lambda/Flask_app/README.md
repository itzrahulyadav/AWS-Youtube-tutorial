# Create a simple flask application

```
from flask import Flask

app = Flask(__name__)

@app.route('/')
def home():
    return "Hello, Flask!"

@app.route('/welcome')
def welcome():
    return "Welcome to the home page"


if __name__ == '__main__':
    app.run(debug=True)

```
-  Create requirements.txt

```
Flask==3.1.0
Werkzeug==3.1.3

```

- Create a Dockerfile

```
# Use an official Python runtime as a parent image
FROM python:3.12

# Set the working directory in the container
WORKDIR /app

# Copy the application files to the container
COPY . /app

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose the port Flask runs on
EXPOSE 5000

# Set the environment variable to run Flask
ENV FLASK_APP=app.py
ENV FLASK_RUN_HOST=0.0.0.0

# Command to run the Flask application
CMD ["flask", "run"]

```
# Create a ECR Repo

```


```
