from flask import Flask, g, Response
from flask_expects_json import expects_json
import json
from admin_model import Admin
from security import hashing_password, create_token

app = Flask(__name__)
admin_prefix = "/api/v1/admin-auth"
citizen_prefix = "/api/v1/citizen-auth"
message = {
    'message': "Login Success"
}

admin_auth_schema = {
    'type': 'object',
    'properties': {
        'username': {'type': 'string'},
        'password': {'type': 'string'}
       
    },
    'required': ['username', 'password']
}

citizen_auth_schema = {
    'type': 'object',
    'properties': {
        'nik': {'type': 'string'},
        'password': {'type': 'string'}
       
    },
    'required': ['username', 'password']
}

@app.route(admin_prefix, methods=['POST'])
@expects_json(auth_schema)
def authentication_administrator():
    username = g.data['username']
    password = g.data['password']
    hashed_password = hashing_password(password) 
    admin = Admin.objects(username=username, password=hashed_password)
    if admin.count() <= 0:
        message['message'] = "Admin is not found"
        return Response(json.dumps(message), status=404, mimetype='application/json')
    admin = admin.first()
    message['token'] = create_token(admin)
    return Response(json.dumps(message), status=200, mimetype='application/json')

@app.route(admin_prefix, methods=['POST'])
@expects_json(citizen_auth_schema)
def authentication_citizen():
    nik = g.data['nik']
    password = g.data['password']
    hashed_password = hashing_password(password)
    
    

if __name__ == '__main__':
    app.run(debug=True)