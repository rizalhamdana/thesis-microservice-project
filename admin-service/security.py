import hashlib
import jwt
import os


def hashing_password(password):
    hasher = hashlib.sha384(password.encode())
    return hasher.hexdigest()

