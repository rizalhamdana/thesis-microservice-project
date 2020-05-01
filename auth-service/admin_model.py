from mongoengine import connect, Document, StringField

import os

root_password = os.getenv('mongodb-root-password')
host_db = 'mongodb+srv://rizalhamdana:21mei1998@cluster0-inove.gcp.mongodb.net/test?retryWrites=true&w=majority'
print(host_db)

connect(host=host_db)


class Admin(Document):
    full_name = StringField(required=True, max_length=50)
    username = StringField(required=True, max_length=20)
    location = StringField(default="")
    password = StringField(required=True)

    def __str__(self):
        return self.username
