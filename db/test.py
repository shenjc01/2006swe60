import requests
import sqlite3
import time
from tqdm import tqdm

def get_lat_long(address):
    url = "https://www.onemap.gov.sg/api/common/elastic/search"
    params = {
        "searchVal": address,
        "returnGeom": "Y",
        "getAddrDetails": "Y",
        "pageNum": 1,
    }
    try:
        response = requests.get(url, params=params)

        # Check if the response is valid
        if response.status_code != 200:
            print(f"Error: Received status code {response.status_code} for address: {address}")
            return None, None

        # Try to decode JSON response
        data = response.json()
        if data and 'results' in data and len(data['results']) > 0:
            result = data['results'][0]
            lat = result['LATITUDE']
            lon = result['LONGITUDE']
            return lat, lon
        else:
            print(f"No data found for address: {address}")
            return None, None

    except requests.exceptions.RequestException as e:
        print(f"Request failed for address {address}: {e}")
        return None, None

# Full addresses dataset with name, opening hours, and address
data = [
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "636", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "637", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "638", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "639", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "640", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "641", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "642", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "643", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "644", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "645", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "646", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "647", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "648", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir Road / Jalan Tenaga", "block_number": "649", "collection_frequency": "Wednesday", "address": "Bedok Reservoir Road / Jalan Tenaga, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "761", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "762", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "763", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "764", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "765", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "766", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "767", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "768", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "769", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "770", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "771", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "772", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "773", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok Reservoir View", "block_number": "774", "collection_frequency": "Wednesday", "address": "Bedok Reservoir View, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "31", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "32", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "33", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "34", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "35", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "36", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "37", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "62", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "63", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "64", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "65", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "1", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "2", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "3", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "4", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "19", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "20", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "21", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "5", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "6", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "7", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "8", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "10B", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "10C", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "10D", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "10E", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "10F", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 2", "block_number": "14", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 2, Singapore"},
    {"name": "Bedok South Avenue 3", "block_number": "155", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 3, Singapore"},
    {"name": "Bedok South Avenue 3", "block_number": "156", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 3, Singapore"},
    {"name": "Bedok South Avenue 3", "block_number": "157", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 3, Singapore"},
    {"name": "Bedok South Avenue 3", "block_number": "158", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 3, Singapore"},
    {"name": "Bedok South Avenue 3", "block_number": "159", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 3, Singapore"},
    {"name": "Bedok South Avenue 3", "block_number": "160", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 3, Singapore"},
    {"name": "Bedok South Avenue 3", "block_number": "161", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 3, Singapore"},
    {"name": "Bedok South Avenue 3", "block_number": "162", "collection_frequency": "Monday, Wednesday, Friday", "address": "Bedok South Avenue 3, Singapore"}
]
# SQLite database connection
conn = sqlite3.connect('data.db')
cursor = conn.cursor()

# Loop through each address, geocode it, and insert it into the SQLite database
for entry in tqdm(data):
    if entry["name"] == "Bedok South Avenue 3":
        address = "460"+''.join(x for x in entry['block_number'] if not x.isalpha()).zfill(3)
    elif entry["name"] == "Bedok South Avenue 2":
        address = "460"+''.join(x for x in entry['block_number'] if not x.isalpha()).zfill(3)
    elif entry["name"] == "Bedok Reservoir View":
        address = "470"+''.join(x for x in entry['block_number'] if not x.isalpha()).zfill(3)
    else:
        address = "410"+''.join(x for x in entry['block_number'] if not x.isalpha()).zfill(3)
    latitude, longitude = get_lat_long(address)
    address = f"Blk {entry['block_number']} {entry['name']}, Singapore {address}"
    name = "Blk "+entry['block_number'] + " Recycling Bin"
    if latitude and longitude:
        try:
            # Check if the entry already exists in the database
            cursor.execute('''
                SELECT * FROM Locations WHERE Latitude = ? AND Longitude = ?
            ''', (latitude, longitude))
            existing_entry = cursor.fetchone()

            # Insert or replace the entry based on whether it already exists
            cursor.execute('''
                INSERT OR REPLACE INTO Locations (Name, Address, Latitude, Longitude, "Opening Hours")
                VALUES (?, ?, ?, ?, ?)
            ''', (name, address, latitude, longitude, entry['collection_frequency']))

            if existing_entry:
                print(f"Replaced {entry['name']} Block {entry['block_number']} with new lat: {latitude}, long: {longitude}")
            else:
                print(f"Inserted new entry {entry['name']} Block {entry['block_number']} with lat: {latitude}, long: {longitude}")

        except sqlite3.Error as e:
            print(f"SQLite error for {entry['name']} Block {entry['block_number']}: {e}")
    else:
        print(f"Could not geocode {address}")
    
    # Sleep to avoid overwhelming the API with requests (rate limiting)
    time.sleep(1)

# Commit the changes and close the database connection
conn.commit()
conn.close()
print("All data inserted or replaced successfully!")