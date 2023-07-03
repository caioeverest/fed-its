#!/usr/bin/env python3

from flask import Flask, request, jsonify
import requests
import os
import logging
import hmac
import hashlib

app = Flask(__name__)
secret = os.getenv('HASH_SECRET')
api_key = os.getenv('GOOGLE_API_KEY')


@app.route('/route', methods=['POST'])
def get_route():
    data = request.json
    start = data.get('start')
    end = data.get('end')
    signature = request.headers.get('X-Signature')

    app.logger.info('Received request with start: %s, end: %s', start, end)

    # Calculate the HMAC
    hmac_obj = hmac.new(secret.encode(), json.dumps(
        data).encode(), hashlib.sha256)
    calculated_signature = hmac_obj.hexdigest()

    # Compare the provided signature and the calculated signature
    # if signature != calculated_signature:
    #     return jsonify({'status': 'Signatures do not match.'}), 400

    # Validate the request
    if not start or not end:
        return jsonify({'error': 'Invalid request, please provide both a start and end location.'}), 400

    # Build the URL for the Google Maps Directions API
    url = 'https://maps.googleapis.com/maps/api/directions/json'
    params = {
        'origin': start,
        'destination': end,
        'key': api_key  # Replace this with your Google API Key
    }

    # Send a request to the Google Maps Directions API
    response = requests.get(url, params=params)

    # Check the status of the request
    if response.status_code != 200:
        return jsonify({'error': 'An error occurred when trying to retrieve the route.'}), 500

    # Return the route
    return jsonify(google_maps_to_common_structure(response.json()))


def google_maps_to_common_structure(google_maps_response):
    first_route = google_maps_response['routes'][0]
    first_leg = first_route['legs'][0]
    return {
        "start_address": first_leg['start_address'],
        "end_address": first_leg['end_address'],
        "distance": first_leg['distance'],
        "duration": first_leg['duration'],
        "steps": [
            {
                "instruction": step['html_instructions'],
                "name": step.get('maneuver', ''),
                "distance": step['distance'],
                "duration": step['duration'],
                "start_location": step['start_location'],
                "end_location": step['end_location'],
            } for step in first_leg['steps']
        ]
    }


if __name__ == '__main__':
    app.run(debug=True)
