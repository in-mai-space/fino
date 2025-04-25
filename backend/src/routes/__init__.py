from src.routes.healthcheck import healthcheck_router

def setup_routes(app, services):
    """
    Set up all API routes for the application
    
    Parameters:
    -----------
    app: FastAPI
        The FastAPI application instance to which routes will be attached
    services: Integrations
        The external services that the app will use

    """
    # TODO: setup middlewares for routes

    # setup all routes and inject dependencies
    app.include_router(healthcheck_router)