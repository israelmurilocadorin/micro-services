from flask import Flask, jsonify, request
from flask_cors import CORS

app = Flask(__name__)
cors = CORS(app)

 
cars = [
    {   
        'id': 1,
        'model': 'Gol Turbo',
        'year': '2010'
 
    },
    {
        'id': 2,
        'model': 'Fiat 147',
        'year': '2013'
 
    },
    {
        'id': 3,
        'model': 'Fusca',
        'year': '2015'
 
    },
    {
        'id': 5,
        'model': 'Fox turbo',
        'year': '2019'
 
    }
]
 
@app.route('/cars', methods=['GET'])
def home():
    return jsonify(cars), 200



@app.route('/cars/<int:id>', methods=['PUT'])
def change_model_cars(id):
    for car in cars:
        if car['id'] == id:
            car['model'] = request.get_json().get('model')
 
            return jsonify(car), 200
    
    return jsonify({'error': 'not found'}), 404



@app.route('/cars/<int:id>', methods=['GET'])
def cars_availability_id(id):
    for car in cars:
        if car ['id'] == id:
            return jsonify(car), 200
    
    return jsonify({'error':'not found'}), 404



@app.route('/cars', methods=['POST'])
def save_cars():
    data = request.get_json()
    cars.append(data)
 
    return jsonify(data), 201
 


if __name__ == '__main__':
    app.run(debug=True, host='localhost', port=8888)