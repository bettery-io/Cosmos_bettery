const event = require("./createEvent.js")

module.exports = app => {
    app.post("/privateevents/create", async (req, res) => {
        event.createEvent(req, res);
    })
}    