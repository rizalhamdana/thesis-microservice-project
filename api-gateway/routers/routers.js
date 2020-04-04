
var express = require('express');
var router = express.Router()
var citizensRouter = require('./citizens')
var marriedRouter = require('./married')
var birthRouter = require('./birth')
// var familyRouter = require('./family')
var adminRouter = require('./admin')
var authRouter = require('./auth')


router.use((req, res, next) => {
    console.log("Called: ", req.path)
    next()
})

router.use(citizensRouter)
router.use(marriedRouter)
router.use(birthRouter)
// router.use(familyRouter)
router.use(adminRouter)
router.use(authRouter)

module.exports = router
