import threading
import requests
import time
import json

# Configuration
API_URL = "http://ilovek8s.com/books"
THREAD_COUNT = 10  # Number of concurrent threads
SLEEP_TIME = 1  # Seconds to wait between each request in a thread

def test_get():
    while True:
        response = requests.get(API_URL)
        print("GET Request: Status Code:", response.status_code)
        time.sleep(SLEEP_TIME)

def test_post():
    new_book = {"id": 3, "title": "New Book", "author": "Author Name"}
    while True:
        response = requests.post(API_URL, json=new_book)
        print("POST Request: Status Code:", response.status_code)
        time.sleep(SLEEP_TIME)

def test_delete():
    while True:
        response = requests.delete(f"{API_URL}/3")  # Assuming ID 3 exists
        print("DELETE Request: Status Code:", response.status_code)
        time.sleep(SLEEP_TIME)

def test_put():
    updated_book = {"id": 3, "title": "Updated Book", "author": "New Author"}
    while True:
        response = requests.put(f"{API_URL}/3", json=updated_book)  # Assuming ID 3 exists
        print("PUT Request: Status Code:", response.status_code)
        time.sleep(SLEEP_TIME)

if __name__ == "__main__":
    # Creating and starting threads
    for _ in range(THREAD_COUNT):
        threading.Thread(target=test_get).start()
        threading.Thread(target=test_post).start()
        threading.Thread(target=test_delete).start()
        threading.Thread(target=test_put).start()

