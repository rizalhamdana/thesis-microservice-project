var express = require('express');
var router = express.Router();
var adapter = require('./adapter');

const { isAuth } = require("../middlewares/authentication");

var BASE_URL = 'http://family-service:8082';

var api = adapter(BASE_URL);
router.get('/family', isAuth, (req, res) => {
    var fullPath = '/api/v1/family';
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});
router.get('/family/:familyId', isAuth, (req, res) => {
    var familyId = req.params['familyId'];
    var fullPath = '/api/v1/family/' + familyId;
    api.get(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});


router.delete('/family/:familyId', isAuth, (req, res) => {
    var familyId = req.params['familyId'];
    var fullPath = '/api/v1/family/' + familyId;
    api.delete(fullPath).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.put('/family/add/:familyId', isAuth, (req, res) => {
    var familyId = req.params['familyId'];
    var fullPath = '/api/v1/family/add/' + familyId;
    api.put(fullPath, req.body).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.put('/family/location/:familyId', isAuth, (req, res) => {
    var familyId = req.params['familyId'];
    var fullPath = '/api/v1/family/update-location/' + familyId;
    api.put(fullPath, req.body).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

router.put('/family/verify/:familyId', isAuth, (req, res) => {
    var familyId = req.params['familyId'];
    var fullPath = '/api/v1/family/verify/' + familyId;
    api.put(fullPath, req.body).then(resp => {
        res.send(resp.data);
    }).catch(error => {
        res.status(error.response.status);
        res.send(error);
    });
});

module.exports = router