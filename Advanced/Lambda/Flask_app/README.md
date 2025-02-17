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
# Use a minimal Python image
FROM python:3.12-slim

# Add AWS Lambda Web Adapter as an extension
COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.9.0 /lambda-adapter /opt/extensions/lambda-adapter

# Set the working directory inside the container
WORKDIR /app

# Copy application code
COPY . /app

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Set environment variables for Flask & Lambda
ENV FLASK_APP=app.py
ENV PORT=8080  

# Expose the correct port
EXPOSE 8080

# Start the Flask app using Gunicorn with dynamic port
CMD ["gunicorn", "-b", "0.0.0.0:8080", "app:app"]
```
# Create a ECR Repo

```


```

# Create a lambda function

```
aws lambda create-function --function-name flask-lambda --package-type Image  --code <uri>
:latest --role <role_name>

```
