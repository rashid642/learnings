const express = require("express")
const responseTime = require("response-time")
const slowProcess = require("./slowProcess")
const client = require("prom-client")

const { createLogger, transports } = require("winston");
const LokiTransport = require("winston-loki");

const options = {
    transports: [
      new LokiTransport({
        host: "http://127.0.0.1:3100"
      })
    ]
  };
  const logger = createLogger(options);

app = express()
PORT = process.env.PORT | 8000

const collectDefaultMetric = client.collectDefaultMetrics
collectDefaultMetric({register : client.register})

const reqResTime = new client.Histogram({
    name : "http_req_res_time",
    help : "this tells you how much time is taken by a histogram",
    labelNames : ["method", "route", "status_code"],
    buckets : [1, 500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 5500]
})

app.use(responseTime((req, res, time) => {
    reqResTime.labels({
        method : req.method,
        route : req.url,
        status_code : req.statusCode
    }).observe(time)
}))

app.get("/", (req, res) => {
    logger.info("Request came on / route")
    res.json({
        msg : "Hello mate..!!"
    })
})

app.get("/slow", async (req, res) => {
    logger.info("Request came on /slow route")
    try {
        const result = await slowProcess();
        console.log(result)
        res.json({
            msg : "Thanks for waiting..!!"
        })
    } catch (error) {
        logger.error("Request came on / route but some error occured")
        res.json({
            error : error
        })
    }
})

app.get("/metrics", async (req, res) => {
    res.setHeader("Content-Type", client.register.contentType)
    const metrics = await client.register.metrics();
    // console.log("Metrics :", metrics)
    res.send(metrics);
})

app.listen(PORT, ()=>{
    console.log("App running at port :", PORT)
})