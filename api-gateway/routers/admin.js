var express = require('express');
var router = express.Router();
var adapter = require('./adapter');
var isAuth = require('../middlewares/authentication');

var BASE_URL = 'http://admin-service:5000';

var api = adapter(BASE_URL);

router.get('/admin/:username', isAuth, (req, res) => {
    var username = req.params['username'];
    var fullPath = '/api/v1/admin/' + username;
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.post('/admin', isAuth, (req, res) => {
    var fullPath = '/api/v1/admin/'
    api.post(fullPath, req.body).then(resp => {
        res.json(resp.data);
    }).catch(error => {
        console.log(error)
        res.status(error.response.status);
        res.send(error);
    });
});

router.get('/admin', isAuth, (req, res) => {
    var fullPath = '/api/v1/admin/'
    api.get(fullPath).then(resp => {
        res.json(resp.data);
    }).catch(error => {
        console.log(error)
        res.status(error.response.status);
        res.send(error);
    });
});

router.delete('/admin/:username', isAuth, (req, res) => {
    var username = req.params['username'];
    var fullPath = '/api/v1/admin/' + username;
    api.delete(fullPath).then(resp => {
        
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

module.exports = router