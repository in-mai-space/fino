from fastapi import FastAPI
import uvicorn
from config import load_config
from routes import setup_routes
from integrations import setup_integrations
from database import connect_db, Base, get_db

if __name__ == "__main__":
    # set up the configuration
    config = load_config()

    # set up and connect to the database
    db = connect_db(config['database'])

    # initialize a new app
    app = FastAPI(
        title="fino",
        description="syncing bank transactions with Notion for budgeting",
        version="1.0.0",
    )

    # migrate database
    @app.on_event("startup")
    async def startup():
        Base.metadata.create_all(bind=db["engine"])

    # set up all the services
    services = setup_integrations(config)

    # set up app routes
    setup_routes(app, services)

    # make db accessible to all routes
    app.dependency_overrides[get_db] = get_db

    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)