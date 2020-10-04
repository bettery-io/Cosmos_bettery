const event = require("./createEvent.js");
const getEvent = require("./getEvent.js");
const part = require("./participate");

module.exports = app => {
    app.post("/privateevent/create", async (req, res) => {
        event.createEvent(req, res);
    })
    app.get("/privateevent/:id", async (req, res) => {
        getEvent.getEventById(req, res);
    })

    app.post("/privateevent/participate", async (req, res) => {
        part.participate(req, res);
    })
}    