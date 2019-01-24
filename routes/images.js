const express = require('express');
const request = require('request');
const apiConfig = require('../config/api');
const router = express.Router();

router.get('/', (req, res, next) => {
    request.get(apiConfig.url('photos'), {json: true}, (err, resp, body) => {
        body.forEach(photo => {
            delete photo.url;
        });

        res.json(body);
    });
})

module.exports = router;