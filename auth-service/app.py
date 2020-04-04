from flask import Flask, g, Response, request
from flask_expects_json import expects_json
import json
from admin_model import Admin
from security import hashing_password, create_token, token_check
app = Flask(__name__)
prefix = "/api/v1/auth"
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


@app.route(prefix, methods=['POST'])
@expects_json(admin_auth_schema)
def authentication_administrator():
    username = g.data['username']
    password = g.data['password']
    hashed_password = hashing_password(password)
    admin = Admin.objects(username=username, password=hashed_password)
    if admin.count() <= 0:
        message['message'] = "Admin is not found"
        return Response(json.dumps(message), status=404, mimetype='application/json')
    admin = admin.first()
    message['message'] = 'Login success'
    message['token'] = create_token(admin)
    return Response(json.dumps(message), status=200, mimetype='application/json')


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=5500, debug=True)
