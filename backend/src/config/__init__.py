import os
from dotenv import load_dotenv

def load_config():
    load_dotenv()
    
    config = {
        'database': {
            'USER': os.getenv('DB_USER'),
            'PASSWORD': os.getenv('DB_PASSWORD'),
            'HOST': os.getenv('DB_HOST'),
            'PORT': int(os.getenv('DB_PORT')),
            'NAME': os.getenv('DB_NAME'),
            'SSL': os.getenv('DB_SSL')
        },
        'app': {
            'PORT': int(os.getenv('APP_PORT')),
        },
        'notion': {
            'API_KEY': os.getenv('NOTION_API_KEY'),
        },
        'plaid': {
            'CLIENT_ID': os.getenv('PLAID_CLIENT_ID'),
            'SECRET': os.getenv('PLAID_SECRET'),
            'ENV': os.getenv('PLAID_ENV', 'sandbox'),
            'PRODUCTS': os.getenv('PLAID_PRODUCTS', 'transactions').split(','),
            'COUNTRY_CODES': os.getenv('PLAID_COUNTRY_CODES', 'US').split(',')
        },
        'redis': {
            'HOST': os.getenv('REDIS_HOST', 'localhost'),
            'PORT': int(os.getenv('REDIS_PORT', 6379)),
            'PASSWORD': os.getenv('REDIS_PASSWORD', ''),
            'DB': int(os.getenv('REDIS_DB', 0)),
            'SSL': os.getenv('REDIS_SSL', 'false').lower() == 'true'
        },
        'supabase': {
            'JWT_SECRET_TOKEN': os.getenv('SUPABASE_JWT_SECRET_TOKEN'),
        }
    }
    
    return config