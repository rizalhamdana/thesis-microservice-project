var express = require('express');
var router = express.Router();
var adapter = require('./adapter');


var BASE_URL = 'http://auth-service:5500';
var api = adapter(BASE_URL);
router.post('/auth', function (req, res) {
    var fullPath = '/api/v1/auth'
    console.log(req.body)
    api.post(fullPath, req.body).then(resp => {
        res.json(resp.data)
    }).catch(error => {
        console.log(error)
        res.status(error.response.status);
        res.send(error);
    });
});

module.exports = router