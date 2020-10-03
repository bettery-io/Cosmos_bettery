const { exec } = require('child_process');

const createUser = (req, res) => {

    // TODO connect to db

    let email = req.body.email
    let password = req.body.password


    exec(`betterycli keys add ${email}`, (err, stdout, stderr) => {
        if (err) {
            console.log(err);
            res.status(400);
            res.send(err);
            return;
        }
        console.log("1")
        console.log(stdout)
        console.log("2")
        console.log(stderr)
        console.log("3")
        console.log(err)
        let jsonData = JSON.parse(stderr)

        res.status(200);
        res.send({ "status": jsonData })
    })
}

module.exports = {
    createUser
}