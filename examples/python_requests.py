import requests
import random
import string
import concurrent.futures

URL = "http://localhost:8080/writer"

def send_request(msg_key):
    job_id = random.randint(1, 100000)
    resume_id = random.randint(1, 100000)
    user_id = ''.join(random.choices(string.ascii_lowercase, k=10))
    
    params = {
        "msg_key": msg_key,
        "job_id": job_id,
        "resume_id": resume_id,
        "user_id": user_id
    }
    
    response = requests.get(URL, params=params)
    print(f"Request for msg_key='{msg_key}' completed with status code: {response.status_code}")

# Número de peticiones a enviar
num_requests = 10

# Generar cadenas aleatorias para msg_key
random_strings = [''.join(random.choices(string.ascii_lowercase, k=10)) for _ in range(num_requests)]

# Enviar las peticiones concurrentemente
with concurrent.futures.ThreadPoolExecutor() as executor:
    executor.map(send_request, random_strings)
