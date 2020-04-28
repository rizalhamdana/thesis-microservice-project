var express = require('express');
var router = express.Router();
var adapter = require('./adapter');
var isAuth = require('../middlewares/authentication');

var BASE_URL = 'http://married-service:8083';

var api = adapter(BASE_URL);
router.get('/married', isAuth, (req, res) => {
    var fullPath = '/api/v1/married';
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});
router.get('/married/:marriedId', isAuth, (req, res) => {
    var marriedId = req.params['marriedId'];
    var fullPath = '/api/v1/married/' + marriedId;
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.post('/married', isAuth, (req, res) => {
    var fullPath = '/api/v1/married'
    api.post(fullPath, req.body).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.delete('/married/:marriedId', isAuth, (req, res) => {
    var marriedId = req.params['marriedId'];
    var fullPath = '/api/v1/married/' + marriedId;
    api.delete(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.put('/married/verif/:marriedId', isAuth, (req, res) => {
    
    var marriedId = req.params['marriedId'];
    console.log(marriedId);
    var fullPath = '/api/v1/married/verif/' + marriedId;
    api.put(fullPath, req.body).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

module.exports = router