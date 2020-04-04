var jwt = require('jsonwebtoken');
var redis = require('redis')
var jwtAlgorithm = process.env.JWT_ALGORITHM;
var jwtSecretKey = process.env.JWT_PRIVATE_KEY;
var redisPass = process.env.REDIS_PASSWORD || '';

const redisOptions = {
    host: 'redis-master',
    port: 6379,
    logErrors: true
  };
client = redis.createClient(redisOptions);
client.auth(redisPass);

client.on('connect', function(){
    console.log('redis connection established');
});
client.on('error', function(err){
    console.log('Something went wrong...'+ err);
});

function checkTokenCache(token){
    client.get(token, function(error, result){
        if(error){
            return false;
        }
        return true;
    });
    return true;
}

module.exports = (req, res, next) => {
    token = req.header('Token');
    console.log(token)
    if(!token){
        res.status(401).send('Unauthorized');
        return;
    }
    var token = req.headers['Token'];
    var is_auth = checkTokenCache(token)
    if (!is_auth){
        res.status(403).send('Forbidden');
        return;
    }
    next();
}