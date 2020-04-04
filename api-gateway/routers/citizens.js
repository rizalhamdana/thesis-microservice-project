var express = require('express');
var router = express.Router();
var adapter = require('./adapter');
var isAuth = require('../middlewares/authentication');

var BASE_URL = 'http://citizen-service:8080';

var api = adapter(BASE_URL);
router.get('/citizens', isAuth, (req, res) => {
    var fullPath = '/api/v1/citizens';
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    });
});
router.get('/citizens/:nik', isAuth, (req, res) => {
    var nik = req.params['nik'];
    var fullPath = '/api/v1/citizens/' + nik;
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    });
});

router.post('/citizens', isAuth, (req, res) => {
    var fullPath = '/api/v1/citizens/'
    api.post(fullPath, req.body).then(resp => {
        res.send(resp);
    });
});

router.delete('/citizens/:nik', isAuth, (req, res) => {
    var nik = req.params['nik'];
    var fullPath = '/api/v1/citizens/' + nik;
    api.delete(fullPath).then(resp => {
        res.send(resp.data);
    });
});

router.put('/citizens/:nik', isAuth, (req, res) => {
    var nik = req.params['nik'];
    var fullPath = '/api/v1/citizens/' + nik;
    api.put(fullPath, req.body).then(resp => {
        res.send(resp.data);
    });
});

module.exports = router