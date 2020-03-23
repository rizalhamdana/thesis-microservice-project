from mongoengine import connect, Document, StringField

root_password = os.environ('mongodb-root-password')
host_db = 'mongodb://root:{}@mongodb-server/admin-service'

connect(host=host_db.format(root_password))


class Admin(Document):
    full_name = StringField(required=True, max_length=50)
    username = StringField(required=True, max_length=20)
    location = StringField(default="")
    password = StringField(required=True)

    def __str__(self):
        return self.username
