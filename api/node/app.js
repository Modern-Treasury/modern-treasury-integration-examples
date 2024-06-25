// API usage Dependencies

const express = require('express');
const bodyParser = require('body-parser');
const modernTreasury = require('modern-treasury');
const sessions = require('express-session');

var randtoken = require('rand-token');

require('dotenv').config()

// These are the various configuration values used in this example. They are
// pulled from the ENV for ease of use, but can be defined directly or stored
// elsewhere
const {
  MT_ORG_ID,
  MT_API_KEY,
  MT_PUB_KEY,
} = process.env;

const client = new modernTreasury({
  apiKey: MT_API_KEY, // defaults to process.env["MODERN_TREASURY_API_KEY"]
  organizationId: MT_ORG_ID,
});

const PUB_KEY = MT_PUB_KEY;

// Set up express
const app = express();
app.use(bodyParser());

app.use(sessions({
    secret: randtoken.generate(24),
    saveUninitialized:true,
    resave: false
}));

var session;

// POST route to handle a new account collection form
app.post('/api/create-cp-acf', async function (req, res) {
  try {
    const counterparty = await client.counterparties.create({
      name: req.body['name'],
    });
    const cp_id = counterparty.id;
    const accountCollectionFlow = await client.accountCollectionFlows.create({
      counterparty_id: cp_id,
      payment_types: req.body.rails
    });

    session=req.session;
    session.client_token=accountCollectionFlow.client_token;

    res.redirect('/embed.html');
  }
  
  catch (err) {
    console.log(err)
  }
});

// POST route to handle a new payment form
app.post('/api/create-cp-pf', async function (req, res) {
  try {
    const counterparty = await client.counterparties.create({
      name: req.body['name'],
    });
    const cp_id = counterparty.id;
    const paymentFlow = await client.paymentFlows.create({
      counterparty_id: cp_id,
      amount: req.body['amount'] * 100,
      direction: req.body['direction'],
      currency: req.body['currency'],
      originating_account_id: req.body['originating_account_id']
    });

    session=req.session;
    session.client_token=paymentFlow.client_token;

    res.redirect('/embed.html');
  }

  catch (err) {
    console.log(err)
  }
});

// This endpoint provides configuration to modern-treasury-js
app.get('/config', function (req, res) {
  res.setHeader('Content-Type', 'application/javascript');
  res.send(`window.mtConfig = { publishableKey: '${PUB_KEY}',` + `clientToken: '${session.client_token}',` + `}`);
});

// Mounts express.static to render example forms
const pubDirPath = '../../public';

app.use(express.static(pubDirPath));

// Start the server
app.listen(9001, function () {
  console.log('Listening on port 9001');
});