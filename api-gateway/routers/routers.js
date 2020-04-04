
var express = require('express');
var router = express.Router()
var citizensRouter = require('./citizens')

router.use((req, res, next) => {
    console.log("Called: ", req.path)
    next()
})

router.use(citizensRouter)

module.exports = router
