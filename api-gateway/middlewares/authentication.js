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
    is_auth = client.get(token, function(error, result){
        if(error){
            console.log(error);
            return false;
        }
        else {
            console.log(result);
            return true;
        }
    });
    return is_auth;
}

const saveToken = (token, data) => {
    client.set(token, data, function(err, reply){
        if (err){
            console.log(err);
        }
        return;
    });
    console.log("set expires");
    client.expire(token, 86400)
    return
}


const isAuth = (req, res, next) => {
    token = req.header('Token');
    if(!token){
        res.status(401).send('Unauthorized');
        return;
    }
    var is_auth = checkTokenCache(token)
    if (!is_auth){
        res.status(403).send('Forbidden');
        return;
    }
    next();
}

exports.isAuth = isAuth;
exports.saveToken = saveToken;