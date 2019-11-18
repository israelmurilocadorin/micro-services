var restify = require('restify');
var cors = require('cors');

function hotelr(req, res, next){
    var dataRest = [{
        name: 'Hotel Tai',
        local: 'São Paulo',
        money: '100,00',
    },
    {
        name: 'Hotel Tai',
        local: 'São Paulo',
        money: '100,00',
    },
    {
        name: 'Hotel Tai',
        local: 'São Paulo',
        money: '100,00',
    }]

    res.json(dataRest);
    next();

}

var server = restify.createServer({
    name: 'Node 1'
});

/* http://localhost:8989/hotel -> chamar a função hotelr */
server.get('/hotel', hotelr);

server.use(
    function crossOrigin(req,res,next){
        res.header("Access-Control-Allow-Origin", "*");
        res.header("Access-Control-Allow-Headers", "X-Requested-With");
        return next();
    }
);

server.listen(8989, function(){
    console.log('%s sendo executado', server.name);
});
