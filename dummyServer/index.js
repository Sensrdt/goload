const express = require('express')
const app = express()

app.get('/', (req, res) => {
	  res.send(`Hello World! from ${process.argv[2]}`)
})


app.listen(process.argv[2], () => {
	  console.log(`Example app listening at http://localhost:${process.argv[2]}`)
})
