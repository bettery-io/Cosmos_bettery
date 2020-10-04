const event = require("./createEvent.js");
const getEvent = require("./getEvent.js");
const action = require("./action.js");

module.exports = app => {
    app.post("/privateevent/create", async (req, res) => {
        event.createEvent(req, res);
    })
    app.get("/privateevent/:id", async (req, res) => {
        getEvent.getEventById(req, res);
    })
    app.post("/privateevent/participate", async (req, res) => {
        action.participate(req, res);
    })
    app.post("/privateevent/validate", async (req, res) => {
        action.validate(req, res);
    })
}    