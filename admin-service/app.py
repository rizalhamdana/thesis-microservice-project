from flask import Flask
import admin_controller


app = Flask(__name__)
prefix = '/api/v1/admin'


app.add_url_rule(f'{prefix}/<username>', 'get_one_admin',
                 admin_controller.get_one_admin, methods=['GET'])
app.add_url_rule(f'{prefix}/', 'create_account',
                 admin_controller.create_account, methods=['POST'])
app.add_url_rule(f'{prefix}/<username>', 'delete_account',
                 admin_controller.delete_account, methods=['DELETE'])


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=5000, debug=True)
