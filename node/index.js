const express = require("express");
const bodyParser = require('body-parser')
const app = express();
const cors = require('cors');
const http = require('http');

app.use(cors({
    origin: "*"
}))
app.use(bodyParser.json());
app.use(bodyParser.text());
app.use(bodyParser.urlencoded({
    extended: true
}));

require('./services/privateevents')(app);
require('./services/auth')(app);

const httpServer = http.createServer(app);

httpServer.listen(80, async () => {
    console.log("server listen port 80")
})

