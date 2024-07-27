const express = require("express");

const app = express()

const PORT = process.env.PORT || 8000

app.get("/", (req, res) => {
    return res.json({
        msg : "hello from backend"
    })
})

app.listen(PORT, ()=>{
    console.log("Server started at port :", PORT)
})