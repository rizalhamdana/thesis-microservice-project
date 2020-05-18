var express = require('express');
var router = express.Router();
var adapter = require('./adapter');
// var isAuth = require('../middlewares/authentication');

const { isAuth, saveToken } = require("../middlewares/authentication");

var BASE_URL = 'http://citizen-service:8080';

var api = adapter(BASE_URL);
router.get('/citizens', isAuth, (req, res) => {
    var fullPath = '/api/v1/citizens';
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});
router.get('/citizens/:nik', isAuth, (req, res) => {
    var nik = req.params['nik'];
    var fullPath = '/api/v1/citizens/' + nik;
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.post('/citizens', isAuth, (req, res) => {
    var fullPath = '/api/v1/citizens'
    api.post(fullPath, req.body).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.delete('/citizens/:nik', isAuth, (req, res) => {
    var nik = req.params['nik'];
    var fullPath = '/api/v1/citizens/' + nik;
    api.delete(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.post('/citizens/auth', (req, res) => {
    var fullPath = '/api/v1/citizens/auth'
    api.post(fullPath, req.body).then(resp => {
        const token = resp.data;
        console.log(token);
        saveToken(token, JSON.stringify(req.body));
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.put('/citizens/verif/:nik', isAuth, (req, res) => {
    var nik = req.params['nik'];
    var fullPath = '/api/v1/citizens/verify' + nik;
    api.put(fullPath, req.body).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});


router.put('/citizens/:nik', isAuth, (req, res) => {
    var nik = req.params['nik'];
    var fullPath = '/api/v1/citizens/' + nik;
    api.put(fullPath, req.body).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

module.exports = router