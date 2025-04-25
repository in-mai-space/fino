from fastapi import FastAPI
import uvicorn
from src.config import load_config
from src.routes import setup_routes
from src.integrations import setup_integrations
from src.database import connect_db, Base, get_db

# Load config outside the main block
config = load_config()

# Connect to the database
db = connect_db(config['database'])

# Initialize the app at module level
app = FastAPI(
    title="fino",
    description="syncing bank transactions with Notion for budgeting",
    version="1.0.0",
)

# Setup startup event
@app.on_event("startup")
async def startup():
    Base.metadata.create_all(bind=db["engine"])

# Set up services
services = setup_integrations(config)

# Set up routes
setup_routes(app, services)

# Make db accessible
app.dependency_overrides[get_db] = get_db

# This part only runs when the file is executed directly
if __name__ == "__main__":
    uvicorn.run("src.main:app", host="0.0.0.0", port=8000, reload=True)