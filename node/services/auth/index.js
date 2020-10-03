const auth = require("./auth.js")

module.exports = app => {
    app.post("/registration", async (req, res) => {
        auth.createUser(req, res);
    })
}    