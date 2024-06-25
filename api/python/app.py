#! /usr/bin/env python3.6
"""
Python 3.6 or newer required.
"""
import json
import os
from dotenv import load_dotenv
from modern_treasury import ModernTreasury

load_dotenv(verbose=True)
ORG_ID = os.environ.get("MT_ORG_ID")
API_KEY = os.environ.get("MT_API_KEY")
PUB_KEY = os.environ.get("MT_PUB_KEY")

modern_treasury = ModernTreasury(
    # defaults to os.environ.get("MODERN_TREASURY_API_KEY")
    api_key=API_KEY,
    organization_id=ORG_ID,
)

from flask import Flask, render_template, jsonify, request, session, redirect, url_for, Response
from flask_session import Session

PUBLIC_DIR_PATH = os.getenv('PUBLIC_DIR_PATH', '../../public')

app = Flask(__name__, static_folder=PUBLIC_DIR_PATH, static_url_path='')

SESSION_TYPE = "filesystem"

app.config.update(SECRET_KEY=os.urandom(24))

app.config.from_object(__name__)
Session(app)

@app.route('/')
@app.route('/index')
def index():
    return redirect('index.html')

# POST route to handle a new account collection form
@app.route('/api/create-cp-acf', methods=['POST'])
def create_cp_acf():
    try:
        # Create a Counterparty with the name given in the form
        counter_party = modern_treasury.counterparties.create(
            name=request.form['name'],
        )
        cp_id = counter_party.id

        account_collection_flow = modern_treasury.account_collection_flows.create(
            counterparty_id=cp_id,
            payment_types = request.form.getlist('rails[]')
        )
        session['client_token'] = account_collection_flow.client_token
        return redirect(url_for('embed'))
        
    except modern_treasury.APIConnectionError as e:
        print("The server could not be reached")
        print(e.__cause__)  # an underlying Exception, likely raised within httpx.
    except modern_treasury.RateLimitError as e:
        print("A 429 status code was received; we should back off a bit.")
    except modern_treasury.APIStatusError as e:
        print("Another non-200-range status code was received")
        print(e.status_code)
        print(e.response)

# POST route to handle a new payment form
@app.route('/api/create-cp-pf', methods=['POST'])
def create_cp_pf():
    try:
        # Create a Counterparty with the name given in the form
        counter_party = modern_treasury.counterparties.create(
            name=request.form['name'],
        )
        cp_id = counter_party.id

        payent_flow = modern_treasury.payment_flows.create(
            counterparty_id=cp_id,
            amount = int(float(request.form['amount']) * 100),
            direction = request.form['direction'],
            currency = request.form['currency'],
            originating_account_id = request.form['originating_account_id']
        )
        session['client_token'] = payent_flow.client_token
        return redirect(url_for('embed'))
        
    except modern_treasury.APIConnectionError as e:
        print("The server could not be reached")
        print(e.__cause__)  # an underlying Exception, likely raised within httpx.
    except modern_treasury.RateLimitError as e:
        print("A 429 status code was received; we should back off a bit.")
    except modern_treasury.APIStatusError as e:
        print("Another non-200-range status code was received")
        print(e.status_code)
        print(e.response)

@app.route('/embed')
def embed():
    return redirect('embed.html')

@app.route('/uo_embed')
def uo_embed():
    return redirect('uo_embed.html')


# This endpoint provides configuration to modern-treasury-js
@app.route("/config", methods=['GET'])
def config_js():
    if session.get('client_token') is not None:
        return Response("window.mtConfig = { publishableKey: '" + PUB_KEY + "', clientToken: '" + session.get('client_token') + "' }", mimetype='application/javascript')  