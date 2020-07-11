const express = require('express')
var bodyParser = require('body-parser')
const app = express()
app.use(bodyParser.json())


const port = process.env.PORT || 8000
const spawn = require("child_process").spawn;

app.get('/', (req, res) => res.send('Hello World!'))

app.post('/process', (req, res) => {
    const pythonProcess = spawn('python',["./open_words/execute.py", req.body.text]);
    pythonProcess.stdout.on('data', (data) => {
        console.log(data.toString())
        res.send(data.toString())
        // Do something with the data returned from python script
    });

})







app.listen(port, () => console.log(`Listening at http://localhost:${port}`))