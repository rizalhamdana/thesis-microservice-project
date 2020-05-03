var express = require("express");
var router = express.Router();
var adapter = require("./adapter");
var isAuth = require("../middlewares/authentication");

var BASE_URL = "http://birth-service:8081";

var api = adapter(BASE_URL);
router.get("/birth", isAuth, (req, res) => {
  var fullPath = "/api/v1/birth";
  api
    .get(fullPath)
    .then((resp) => {
      res.send(resp.data);
    })
    .catch((error) => {
      res.status(error.response.status);
      res.send(error);
    });
});
router.get("/birth/:birthId", isAuth, (req, res) => {
  var birthId = req.params["birthId"];
  var fullPath = "/api/v1/birth/" + birthId;
  api
    .get(fullPath)
    .then((resp) => {
      res.send(resp.data);
    })
    .catch((error) => {
      res.status(error.response.status);
      res.send(error);
    });
});

router.post("/birth", isAuth, (req, res) => {
  var fullPath = "/api/v1/birth";
  api
    .post(fullPath, req.body)
    .then((resp) => {
      res.send(resp.data);
    })
    .catch((error) => {
      res.status(error.response.status);
      res.send(error);
    });
});

router.delete("/birth/:birthId", isAuth, (req, res) => {
  var birthId = req.params["birthId"];
  var fullPath = "/api/v1/birth/" + birthId;
  api
    .delete(fullPath)
    .then((resp) => {
      res.send(resp.data);
    })
    .catch((error) => {
      res.status(error.response.status);
      res.send(error);
    });
});

router.put("/birth/:birthId", isAuth, (req, res) => {
  var birthId = req.params["birthId"];
  var fullPath = "/api/v1/birth/" + birthId;
  api
    .put(fullPath, req.body)
    .then((resp) => {
      res.send(resp.data);
    })
    .catch((error) => {
      res.status(error.response.status);
      res.send(error);
    });
});

router.put("/birth/verif/:birthId", isAuth, (req, res) => {
  var birthId = req.params["birthId"];
  var fullPath = "/api/v1/birth/verif/" + birthId;
  api
    .put(fullPath, req.body)
    .then((resp) => {
      res.send(resp.data);
    })
    .catch((error) => {
      res.status(error.response.status);
      res.send(error);
    });
});

module.exports = router;
