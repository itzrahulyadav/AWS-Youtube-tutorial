from flask import Flask
app = Flask(__name__)

@app.route('/')
def main():
    return "Hello from flask"

@app.route('/home')
def welcome():
    return "Welcome to the home page"


if __name__ == '__main__':
    app.run(debug=True)
