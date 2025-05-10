from notion_client import Client
import plaid
from plaid.api import plaid_api

def setup_integrations(config):
    """
    Setup clients for external service integration
    """
    services = {}
    
    notion_client = create_notion_client(config['notion'])
    if notion_client:
        services['notion'] = notion_client
    
    plaid_client = create_plaid_client(config['plaid'])
    if plaid_client:
        services['plaid'] = plaid_client
    
    return services

def create_notion_client(config):
    try:
        notion = Client(auth=config['API_KEY'])
    except Exception as e:
        print(f"Error initializing Notion client: {e}")

    return notion

def create_plaid_client(config):
    try:   
        configuration = plaid.Configuration(
        host=plaid.Environment.Sandbox if config.get('ENV') == 'sandbox' else 
            plaid.Environment.Development if config.get('ENV') == 'development' else
            plaid.Environment.Production
        )
            
        configuration.api_key['clientId'] = config['CLIENT_ID']
        configuration.api_key['secret'] = config['SECRET']
            
        api_client = plaid.ApiClient(configuration)
        plaid_client = plaid_api.PlaidApi(api_client)
    except Exception as e:
        print(f"Error initializing Plaid client: {e}")

    return plaid_client