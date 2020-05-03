import json
import hashlib
from flask import g, Response
from flask_expects_json import expects_json
from model import Admin
from security import hashing_password

message = {
    'message': ''
}

admin_schema = {
    'type': 'object',
    'properties': {
        'username': {'type': 'string'},
        'full_name': {'type': 'string'},
        'location': {'type': 'string'},
        'password': {'type': 'string'},
    },
    'required': ['username', 'full_name', 'location', 'password']
}


def get_one_admin(username):
    admin = Admin.objects(username=username).exclude('id', 'password')
    if admin.count() <= 0:
        message['message'] = "Admin is not found"
        return Response(json.dumps(message), status=404, mimetype='application/json')
    admin = admin.first()
    return Response(admin.to_json(), status=200, mimetype='application/json')

def get_all_admin():
    admin = Admin.objects.exclude('id', 'password')
    return Response(admin.to_json(), status=200, mimetype='application/json')



@expects_json(admin_schema)
def create_account():
    admin = Admin().from_json(json.dumps(g.data))
    existed_username_count = Admin.objects(username=admin['username']).count()
    if existed_username_count > 0:
        message['message'] = 'Admin already existed'
        return Response(json.dumps(message), status=400, mimetype='application/json')
   
    admin.password = hashing_password(admin.password)
    is_save = admin.save()
    if is_save is None:
        message['message'] = 'Failed to create new admin account'
        return Response(json.dumps(message), status=400, mimetype='application/json')
    message['message'] = 'Admin successfully created'
    return Response(json.dumps(message), status=200, mimetype='application/json')


def delete_account(username):
    admin = Admin.objects(username=username).first()
    if admin is None:
        message['message'] = "Account does not exist"
        return Response(json.dumps(message), status=400, mimetype='application/json')
    admin.delete()
    message['message'] = 'Account is successfully deleted'
    return Response(json.dumps(message), status=200, mimetype='application/json')