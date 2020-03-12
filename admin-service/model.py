from mongoengine import connect, Document, StringField
connect(db="admin-service")

class Admin(Document):
    full_name = StringField(required=True, max_length=50)
    username = StringField(required=True, max_length=20)
    location = StringField(default="")
    password = StringField(required=True)

    def __str__(self):
        return self.username