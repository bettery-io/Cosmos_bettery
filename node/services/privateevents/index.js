const event = require("./createEvent.js");
const getEvent = require("./getEvent.js");

module.exports = app => {
    app.post("/privateevents/create", async (req, res) => {
        event.createEvent(req, res);
    })
    app.get("/privateevents/:id", async (req, res) => {
        getEvent.getEventById(req, res);
    })
}    