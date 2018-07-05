var express = require('express');
var router = express.Router();

/* GET users listing. */
router.get('/', function (req, res, next) {
  res.send('respond with a resource');
});

/* POST Add user. */
router.post('/Add', function (req, res, next) {
  console.log(req.body, req.body.images)
  res.send('hello')
})

module.exports = router;