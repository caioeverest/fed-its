#!/usr/bin/env python3

from flask import Flask, request, jsonify
import requests
import json
import os
import logging
import hmac
import hashlib

app = Flask(__name__)
secret = os.getenv('HASH_SECRET')
api_key = os.getenv('OPEN_ROUTE_SERVICE_API_KEY')


def parse_coordinates(coord_str):
    return list(map(float, coord_str.split(',')))


@app.route('/route', methods=['POST'])
def get_route():
    data = request.json
    start = request.json.get('start')
    end = request.json.get('end')
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
        app.logger.warning(
            'Invalid request, please provide both a start and end location.')
        return jsonify({'error': 'Invalid request, please provide both a start and end location.'}), 400

    app.logger.info('Requesting open route service')
    # Convert addresses to coordinates using OpenRouteService's Geocoding API
    geocode_url = 'https://api.openrouteservice.org/geocode/search'
    headers = {
        'Authorization': api_key,  # Replace this with your OpenRouteService API Key
        'Accept': 'application/json, application/geo+json, application/gpx+xml, img/png; charset=utf-8',
    }
    start_coords = requests.get(geocode_url, headers=headers, params={
                                'text': start}).json()['features'][0]['geometry']['coordinates']
    end_coords = requests.get(geocode_url, headers=headers, params={'text': end}).json()[
        'features'][0]['geometry']['coordinates']

    # Build the URL for the OpenRouteService Directions API
    directions_url = 'https://api.openrouteservice.org/v2/directions/driving-car'
    headers['Content-Type'] = 'application/json; charset=utf-8'

    body = {
        'coordinates': [start_coords, end_coords]
    }

    # Send a request to the OpenRouteService Directions API
    response = requests.post(
        directions_url, headers=headers, data=json.dumps(body))

    # Check the status of the request
    if response.status_code != 200:
        return jsonify({'error': 'An error occurred when trying to retrieve the route.'}), 500

    # Return the route summary
    return jsonify(open_route_to_common_structure(response.json()))


def open_route_to_common_structure(open_route_response):
    first_route = open_route_response['routes'][0]
    first_segment = first_route['segments'][0]
    return {
        "start_address": '',  # OpenRouteService does not provide the start and end addresses
        "end_address": '',
        "distance": {
            # convert to km
            "text": f"{first_route['summary']['distance']/1000} km",
            "value": first_route['summary']['distance']
        },
        "duration": {
            # convert to minutes
            "text": f"{first_route['summary']['duration']/60} mins",
            "value": first_route['summary']['duration']
        },
        "steps": [
            {
                "instruction": step['instruction'],
                "name": step['name'],
                "distance": {
                    "text": f"{step['distance']/1000} km",  # convert to km
                    "value": step['distance']
                },
                "duration": {
                    # convert to minutes
                    "text": f"{step['duration']/60} mins",
                    "value": step['duration']
                },
                "start_location": {},  # OpenRouteService does not provide step coordinates
                "end_location": {},
            } for step in first_segment['steps']
        ]
    }


if __name__ == '__main__':
    app.run(debug=True)
