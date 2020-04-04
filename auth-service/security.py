import hashlib
import jwt
import os
import redis
import json

redis_password = os.getenv('redis-password')
redis = redis.Redis(host='redis-master', password=redis_password)


def hashing_password(password):
    hasher = hashlib.sha384(password.encode())
    return hasher.hexdigest()

def create_token(user):
    payload = {
        'username': user.username,
        'password': user.password,
        'location': user.location
    }
    private_key = "PRIVATE_KEY:" + os.environ['JWT_PRIVATE_KEY']
    algorithm = os.environ['JWT_ALGORITHM']
    encode = jwt.encode(payload, private_key, algorithm=algorithm)
    token = encode.decode('ascii')
    cache_data = json.dumps({
        'cached': True,
        'data': str(payload)
    })
    caching(token, cache_data)
    return token


def caching(token, data):
    redis.setex(token, 86400, str(data))

