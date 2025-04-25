import os
from dotenv import load_dotenv

def load_config():
    load_dotenv()
    
    config = {
        'database': {
            'FINO_DB_USER': os.getenv('FINO_DB_USER'),
            'FINO_DB_PASSWORD': os.getenv('FINO_DB_PASSWORD'),
            'FINO_DB_HOST': os.getenv('FINO_DB_HOST'),
            'FINO_DB_PORT': int(os.getenv('FINO_DB_PORT')),
            'FINO_DB_NAME': os.getenv('FINO_DB_NAME'),
            'FINO_DB_SSL': os.getenv('FINO_DB_SSL')
        },
        'app': {
            'FINO_APP_PORT': int(os.getenv('FINO_APP_PORT')),
        }
    }
    
    return config